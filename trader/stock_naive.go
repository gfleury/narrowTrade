package trader

import (
	"log"
	"math"
	"sort"
	"time"

	"github.com/gfleury/narrowTrade/analysis"
	"github.com/gfleury/narrowTrade/config"
	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/utils"
)

type StockNaive struct {
	*BasicSaxoTrader
	data               []InstrumentNaiveData
	availableCash      *AvailableCash
	boughtSymbolsCache *utils.Cache
}

type InstrumentNaiveData struct {
	instrument       models.InstrumentDetails
	quote            *analysis.Quote
	buyRecomendation int
	betterBuy        bool
	indicator        []float64
}

func (t *StockNaive) UpdateAvailableCash(availableCash float64) error {

	// Lazy initialization of availableCash
	if t.availableCash == nil {
		t.availableCash = &AvailableCash{}
	}

	// if cashToUse equal zero, than use the AvailableMargin from the account
	if availableCash == 0 {
		balance, err := t.Api().GetBalanceMe()
		if err != nil {
			return err
		}
		t.availableCash.SetAvailableCash(balance.InitialMargin.MarginAvailable)
	} else {
		t.availableCash.SetAvailableCash(availableCash)
	}

	openOrdersTotal := t.getOpenOrders()

	t.availableCash.Spent(openOrdersTotal)

	return nil
}

func (t *StockNaive) Trade(param TradeParameter) error {
	// Initialize symbol caching
	if t.boughtSymbolsCache == nil {
		t.boughtSymbolsCache = &utils.Cache{}
	}

	t.data = t.createStocksNaive(param.Symbols)

	err := t.UpdateAvailableCash(param.TotalInvest)
	if err != nil {
		return err
	}

	qntSymbols := param.MaxBuySymbols

	cashPerSymbol := t.availableCash.GetSplit(qntSymbols)

	log.Printf("Using %f per symbol, totalzing %f\n", cashPerSymbol, t.availableCash.GetTotal())

	failedTrades := 0
	successfulTrades := 0
	for idx, n := range t.data {
		var profitPrice, stopLossPrice, distanceToMarket float64

		if t.boughtSymbolsCache.Get(n.instrument.GetSymbol()) != nil {
			// We bought this symbol before and the timeout still haven't expired
			continue
		}

		if n.buyRecomendation != 0 && !n.betterBuy {
			log.Println("Skipping symbol as it does not seem good to buy", n.instrument)
		}

		analysisQuote := n.quote
		i := n.instrument

		if i == nil {
			log.Println("Instrument is nil, something is wrong:", n)
			failedTrades++
			cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
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
			cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
			continue
		}

		if n.indicator == nil || len(n.indicator) < 3 {
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
		if amount < 1 || buyPrice > cashPerSymbol || buyPrice*float64(amount) < config.MINIMUM_TRADE_VALUE {
			log.Printf("Not enough available money to buy OR trade value too low %s %s for %f",
				i.GetAssetType(), i.GetSymbol(), buyPrice*float64(amount))
			failedTrades++
			cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
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
				WithTakeProfit(profitPrice))
		// WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05))
		if err != nil {
			orderError := models.GetOrderError(err)
			if orderError != nil && models.BusinessRuleViolation(orderError) {
				failedTrades++
				cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
				continue
			}
			return err
		}

		// Add bought symbol with 2 hours cache
		t.boughtSymbolsCache.Put(n.instrument.GetSymbol(), true, time.Hour*2)

		// Save Ordered price for next calculation
		or.TotalPrice = float64(amount) * buyPrice
		log.Println(or)

		// rebalance cashPerSymbol based on remaining cash
		t.availableCash.Spent(float64(amount) * buyPrice)
		cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))

		successfulTrades++

		if successfulTrades > param.MaxBuySymbols {
			log.Printf("Bought %d symbols stopping for now", successfulTrades)
			break
		}
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

func (t *StockNaive) createStocksNaive(uics []int) []InstrumentNaiveData {
	naiveStockData := []InstrumentNaiveData{}
	for _, uic := range uics {
		n := InstrumentNaiveData{}
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
			log.Println("Failed to get Analyser, flatting symbols to same level", i.GetSymbolSimple(), err)
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
