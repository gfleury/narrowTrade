package trader

import (
	"log"
	"math"
	"sort"

	"github.com/gfleury/narrowTrade/analysis"
	"github.com/gfleury/narrowTrade/models"
)

type StockNaive struct {
	*BasicSaxoTrader
	data []StockNaiveData
}

type StockNaiveData struct {
	instrument       models.InstrumentDetails
	quote            *analysis.Quote
	buyRecomendation int
	betterBuy        bool
	indicator        []float64
}

func (t *StockNaive) GetCashPerSymbol(qntInstruments int, availableCash float64) (float64, error) {
	// if cashToUse equal zero, than use the AvailableMargin from the account
	if availableCash == 0 {
		balance, err := t.Api().GetBalanceMe()
		if err != nil {
			return 0, err
		}
		availableCash = balance.InitialMargin.MarginAvailable
	}

	openOrdersTotal := t.getOpenOrders()

	available := availableCash - openOrdersTotal
	if available > 0 {
		return available / float64(qntInstruments), nil
	}
	return 0, nil
}

func (t *StockNaive) GetNewCashPerSymbol(oldCashPerSymbol, totalInvest float64, failedTrades int) float64 {
	newCashPerSymbol, err := t.GetCashPerSymbol(len(t.data)-failedTrades, totalInvest)
	if err == nil {
		log.Printf("Rebalancing cash per symbol from: %f to: %f", oldCashPerSymbol, newCashPerSymbol)
		return newCashPerSymbol
	}
	return oldCashPerSymbol
}

func (t *StockNaive) Trade(param TradeParameter) error {

	t.data = t.createStocksNaive(param.Symbols)

	cashPerSymbol, err := t.GetCashPerSymbol(len(t.data), param.TotalInvest)
	if err != nil {
		return err
	}
	log.Printf("Using %f per symbol, totalzing %f\n", cashPerSymbol, cashPerSymbol*float64(len(t.data)))

	failedTrades := 0
	for idx, n := range t.data {
		var profitPrice, stopLossPrice, distanceToMarket float64

		if n.buyRecomendation != 0 && !n.betterBuy {
			log.Println("Skipping symbol as it does not seem good to buy", n.instrument)
		}

		analysisQuote := n.quote
		i := n.instrument

		if i == nil {
			log.Println("Instrument is nil, something is wrong:", n)
			failedTrades++
			cashPerSymbol = t.GetNewCashPerSymbol(cashPerSymbol, param.TotalInvest, (idx + 1 + failedTrades))
			continue
		}

		log.Println("Fetching price from Saxo Bank for", i.GetAssetType(), i.GetSymbol())
		buyPrice, err := t.Api().GetInstrumentPrice(i)
		if err != nil {
			return err
		}

		orderType := models.OrderType(models.Market)

		durationType := models.DurationType(models.DayOrder)

		log.Printf("Got price %f - %f\n", buyPrice, analysisQuote.RealtimePrice)
		if (buyPrice > analysisQuote.RealtimePrice && buyPrice*0.8 < analysisQuote.RealtimePrice) ||
			buyPrice == 0 {
			buyPrice = i.CalculatePriceWithThickSize(analysisQuote.RealtimePrice, 0)
			orderType = models.Limit
			durationType = models.GoodTillCancel
		}

		if buyPrice <= 0 {
			log.Println("Calculated BuyPrice below/equal 0:", buyPrice)
			failedTrades++
			cashPerSymbol = t.GetNewCashPerSymbol(cashPerSymbol, param.TotalInvest, (idx + 1 + failedTrades))
			continue
		}

		if n.indicator == nil {
			// indicator unavailable, go with percentage
			profitPrice = i.CalculatePriceWithThickSize(buyPrice, -param.PercentProfit)
			stopLossPrice = i.CalculatePriceWithThickSize(buyPrice, param.PercentLoss)
			distanceToMarket = i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 70)
		} else {
			// If percentProfit higher than bbands_higher go with percentProfit
			if n.indicator[2] < buyPrice*((param.PercentProfit/100)+1) {
				profitPrice = i.CalculatePriceWithThickSize(buyPrice, -param.PercentProfit)
			} else {
				profitPrice = i.CalculatePriceWithThickSize(n.indicator[2], -param.PercentProfit)
			}
			// If percentLoss lower than bbands_lower go with bbands_lower
			if n.indicator[0] > buyPrice*(1-(param.PercentLoss/100)) {
				stopLossPrice = i.CalculatePriceWithThickSize(n.indicator[0], 0)
			} else {
				stopLossPrice = i.CalculatePriceWithThickSize(buyPrice, param.PercentLoss)
			}
			distanceToMarket = i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 0)
		}

		amount := int(math.Ceil(cashPerSymbol / buyPrice))
		if amount < 1 || buyPrice > cashPerSymbol {
			log.Printf("Not enough available money to buy %s %s", i.GetAssetType(), i.GetSymbol())
			failedTrades++
			cashPerSymbol = t.GetNewCashPerSymbol(cashPerSymbol, param.TotalInvest, (idx + 1 + failedTrades))
			continue
		}

		log.Printf("Trying to buy %s %s for %d x %f with stopLoss %f (%f) and takeProfit %f\n",
			i.GetAssetType(), i.GetSymbol(), amount, buyPrice, stopLossPrice,
			distanceToMarket, profitPrice)

		or, err := t.Buy(
			i.GetOrder().
				WithType(orderType).
				WithAmount(amount).
				WithPrice(buyPrice).
				WithDuration(durationType).
				WithTakeProfit(profitPrice).
				WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05))
		if err != nil {
			orderError := models.GetOrderError(err)
			if orderError != nil && models.BusinessRuleViolation(orderError) {
				failedTrades++
				cashPerSymbol = t.GetNewCashPerSymbol(cashPerSymbol, param.TotalInvest, (idx + 1 + failedTrades))
				continue
			}
			return err
		}

		// Save Ordered price for next calculation
		or.TotalPrice = float64(amount) * buyPrice
		log.Println(or)

		// rebalance cashPerSymbol based on remaining cash
		cashPerSymbol = t.GetNewCashPerSymbol(cashPerSymbol, param.TotalInvest, (idx + 1 + failedTrades))
		log.Printf("Rebalancing, using %f per symbol, totalzing %f\n", cashPerSymbol, cashPerSymbol*float64(len(t.data)-(idx+1)))
	}
	return nil
}

func (t *StockNaive) GetOrders() ([]string, []float64) {
	return flatOrderResponse(t.tradedOrdersID)
}

func flatOrderResponse(orders []*models.OrderResponse) ([]string, []float64) {
	ordersID := []string{}
	prices := []float64{}
	for _, order := range orders {
		ordersID = append(ordersID, order.OrderID)
		prices = append(prices, order.TotalPrice)
		if order.Orders != nil {
			o, p := flatOrderResponse(order.Orders)
			ordersID = append(ordersID, o...)
			prices = append(prices, p...)
		}
	}
	return ordersID, prices
}

func (t *StockNaive) createStocksNaive(uics []int) []StockNaiveData {
	naiveStockData := []StockNaiveData{}
	for _, uic := range uics {
		n := StockNaiveData{}
		log.Println("Getting Saxo Bank symbol for", uic)
		uicInstrument := &models.SaxoInstrument{
			Identifier: int32(uic),
		}
		i, err := t.Api().GetInstrumentDetails(uicInstrument)
		if err != nil {
			continue
		}
		n.instrument = i

		log.Println("Fetching price at Analyser", i.GetAssetType(), i.GetSymbolSimple())
		n.quote, err = t.Analyser().Quote(i)
		if err != nil {
			n.quote = &analysis.Quote{}
		}

		log.Println("Fetching recommendation from Analyser", i.GetAssetType(), i.GetSymbolSimple())
		recommendation, err := t.Analyser().OneAnalysis(i)
		if err != nil {
			log.Println("Failed to get Analyser, flatting symbols to same level", i.GetSymbolSimple())
			n.buyRecomendation = 0
			n.betterBuy = false
		} else {
			n.buyRecomendation = recommendation.BuyRatings
			n.betterBuy = recommendation.BuyRatings > recommendation.SellRatings
		}

		log.Println("Fetching bollinger bands from Analyser", i.GetAssetType(), i.GetSymbolSimple())
		n.indicator, err = t.Analyser().Indicator(i, analysis.BBANDS)
		if err != nil {
			log.Println("Failed to get bollinger bands, going with percentage", i.GetSymbolSimple(), err)
		}

		naiveStockData = append(naiveStockData, n)
	}

	sort.Slice(naiveStockData, func(i, j int) bool {
		return naiveStockData[i].buyRecomendation > naiveStockData[j].buyRecomendation
	})

	return naiveStockData
}
