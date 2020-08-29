package trader

import (
	"log"
	"math"
	"sort"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/utils"

	iex "github.com/goinvest/iexcloud/v2"
)

type StockNaive struct {
	*BasicSaxoTrader
}

type StockNaiveData struct {
	instrument       models.InstrumentDetails
	iexQuote         iex.Quote
	buyRecomendation int
	betterBuy        bool
	bbands           []float64
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

	return (availableCash - openOrdersTotal) / float64(qntInstruments), nil
}

func (t *StockNaive) Trade(param TradeParameter) error {

	naiveStockData := t.createStocksNaive(param.Symbols)

	cashPerSymbol, err := t.GetCashPerSymbol(len(naiveStockData), param.TotalInvest)
	if err != nil {
		return err
	}
	log.Printf("Using %f per symbol, totalzing %f\n", cashPerSymbol, cashPerSymbol*float64(len(naiveStockData)))

	for _, n := range naiveStockData {
		var profitPrice, stopLossPrice, distanceToMarket float64

		iexQuote := n.iexQuote
		i := n.instrument

		if i == nil {
			log.Println("Instrument is nil, something is wrong:", n)
			continue
		}

		log.Println("Fetching price from Saxo Bank for", i.GetAssetType(), i.GetSymbol())
		buyPrice, err := t.Api().GetInstrumentPrice(i)
		if err != nil {
			return err
		}

		orderType := models.OrderType(models.Market)

		durationType := models.DurationType(models.DayOrder)

		log.Printf("Got price %f - %f\n", buyPrice, iexQuote.IEXRealtimePrice)
		if (buyPrice > iexQuote.IEXRealtimePrice && buyPrice*0.8 < iexQuote.IEXRealtimePrice) ||
			buyPrice == 0 {
			buyPrice = i.CalculatePriceWithThickSize(iexQuote.IEXRealtimePrice, 0)
			orderType = models.Limit
			durationType = models.GoodTillCancel
		}

		if buyPrice <= 0 {
			log.Println("Calculated BuyPrice below/equal 0:", buyPrice)
			continue
		}

		if n.bbands == nil {
			// bollinger bands unavailable, go with percentage
			profitPrice = i.CalculatePriceWithThickSize(buyPrice, -param.PercentProfit)
			stopLossPrice = i.CalculatePriceWithThickSize(buyPrice, param.PercentLoss)
			distanceToMarket = i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 0)
		} else {
			// If percentProfit higher than bbands_higher go with percentProfit
			if n.bbands[2] < buyPrice*((param.PercentProfit/100)+1) {
				profitPrice = i.CalculatePriceWithThickSize(buyPrice, -param.PercentProfit)
			} else {
				profitPrice = i.CalculatePriceWithThickSize(n.bbands[2], -param.PercentProfit)
			}
			// If percentLoss lower than bbands_lower go with bbands_lower
			if n.bbands[0] > buyPrice*(1-(param.PercentLoss/100)) {
				stopLossPrice = i.CalculatePriceWithThickSize(n.bbands[0], 0)
			} else {
				stopLossPrice = i.CalculatePriceWithThickSize(buyPrice, param.PercentLoss)
			}
			distanceToMarket = i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 0)
		}

		amount := int(math.Round(float64(cashPerSymbol) / buyPrice))
		if amount < 1 {
			log.Printf("Not enough available money to buy %s %s", i.GetAssetType(), i.GetSymbol())
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
			return err
		}

		// Save Ordered price for next calculation
		or.TotalPrice = float64(amount) * buyPrice
		log.Println(or)
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

		log.Println("Fetching price IEXCloud price", i.GetAssetType(), i.GetSymbolSimple())
		n.iexQuote, err = t.IEXApi().Quote(t.Api().Ctx, i.GetSymbolSimple())
		if err != nil {
			n.iexQuote = iex.Quote{}
		}

		log.Println("Fetching recommendation from IEXCloud analysis", i.GetAssetType(), i.GetSymbolSimple())
		recommendations, err := t.IEXApi().RecommendationTrends(t.Api().Ctx, i.GetSymbolSimple())
		if err != nil {
			log.Println("Failed to get IEXCloud analysis, ignoring symbol", i.GetSymbolSimple())
			continue
		}
		recommendation := utils.IEXRecomendationReduce(recommendations)
		n.buyRecomendation = recommendation.BuyRatings
		n.betterBuy = recommendation.BuyRatings > recommendation.SellRatings

		log.Println("Fetching bollinger bands from IEXCloud", i.GetAssetType(), i.GetSymbolSimple())
		bbands, err := t.IEXApi().Indicator(t.Api().Ctx, i.GetSymbolSimple(), "bbands", "ytd")
		if err != nil {
			log.Println("Failed to get bollinger bands, going with percentage", i.GetSymbolSimple(), err)
		} else {
			n.bbands = utils.IEXBBandsReduction(bbands, 4)
		}
		naiveStockData = append(naiveStockData, n)
	}

	sort.Slice(naiveStockData, func(i, j int) bool {
		return naiveStockData[i].buyRecomendation > naiveStockData[j].buyRecomendation
	})

	return naiveStockData
}
