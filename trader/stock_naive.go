package trader

import (
	"log"
	"sort"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/utils"

	iex "github.com/goinvest/iexcloud/v2"
)

type NaiveStockData struct {
	instrument       models.InstrumentDetails
	iexQuote         iex.Quote
	buyRecomendation int
	betterBuy        bool
}

func (t *BasicSaxoTrader) GetCashPerSymbol(n int) (float64, error) {
	balance, err := t.ModeledAPI.GetBalanceMe()
	if err != nil {
		return 0, err
	}

	openOrdersTotal := t.getOpenOrders()

	return (balance.InitialMargin.MarginAvailable - openOrdersTotal) / float64(n), nil
}

func (t *BasicSaxoTrader) BuyStocksNaive(symbols []string, percentLoss, percentProfit float64) error {

	naiveStockData := t.createStocksNaive(symbols)

	cashPerSymbol, err := t.GetCashPerSymbol(len(naiveStockData))
	if err != nil {
		return err
	}
	log.Printf("Using %f per symbol, totalzing %f\n", cashPerSymbol, cashPerSymbol*float64(len(naiveStockData)))

	for _, n := range naiveStockData {
		iexQuote := n.iexQuote
		i := n.instrument

		if i == nil {
			log.Println("Instrument is nil, something is wrong:", n)
			continue
		}

		log.Println("Trying to get price Saxo Bank price", i.GetAssetType(), i.GetSymbol())
		buyPrice, err := t.ModeledAPI.GetInstrumentPrice(i)
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

		profitPrice := i.CalculatePriceWithThickSize(buyPrice, -percentProfit)
		stopLossPrice := i.CalculatePriceWithThickSize(buyPrice, percentLoss)
		distanceToMarket := i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 0)

		amount := int(utils.RoundDown(float64(cashPerSymbol)/buyPrice, 0))

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

		log.Println(or)
	}
	return nil
}

func (t *BasicSaxoTrader) createStocksNaive(symbols []string) []NaiveStockData {
	naiveStockData := []NaiveStockData{}
	for _, symbol := range symbols {
		n := NaiveStockData{}
		log.Println("Getting Saxo Bank symbol for", symbol)
		i, err := t.ModeledAPI.GetInstrument(symbol)
		if err != nil {
			continue
		}
		id, err := t.ModeledAPI.GetInstrumentDetails(i)
		if err != nil {
			continue
		}
		n.instrument = id

		log.Println("Trying to get price IEXCloud price", i.GetAssetType(), i.GetSymbolSimple())
		n.iexQuote, err = t.IEXClient.Quote(t.ModeledAPI.Ctx, i.GetSymbolSimple())
		if err != nil {
			n.iexQuote = iex.Quote{}
		}

		log.Println("Trying to get recommendation from IEXCloud analysis", i.GetAssetType(), i.GetSymbolSimple())
		recommendations, err := t.IEXClient.RecommendationTrends(t.ModeledAPI.Ctx, i.GetSymbolSimple())
		if err != nil {
			continue
		}
		recommendation := utils.IEXRecomendationReduce(recommendations)
		n.buyRecomendation = recommendation.BuyRatings
		n.betterBuy = recommendation.BuyRatings > recommendation.SellRatings

		naiveStockData = append(naiveStockData, n)
	}

	sort.Slice(naiveStockData, func(i, j int) bool {
		return naiveStockData[i].buyRecomendation > naiveStockData[j].buyRecomendation
	})

	return naiveStockData
}
