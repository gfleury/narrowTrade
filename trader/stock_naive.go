package trader

import (
	"log"

	"github.com/gfleury/narrowTrade/models"

	iex "github.com/goinvest/iexcloud/v2"
)

func (t *BasicSaxoTrader) BuyStocksNaive(buyingInstruments []models.InstrumentDetails,
	iexQuotes []iex.Quote, percentLoss, percentProfit float64) {

	for idx, i := range buyingInstruments {
		if i == nil {
			continue
		}

		iexQuote := iexQuotes[idx]
		log.Println("Trying to get price Saxo Bank price", i.GetAssetType(), i.GetSymbol())
		buyPrice, err := t.API.GetInstrumentPrice(i)
		if err != nil {
			log.Println(err)
			log.Println(models.GetStringError(err))
			continue
		}

		orderType := models.OrderType(models.Market)

		durationType := models.DurationType(models.DayOrder)

		log.Printf("Got price %f - %f\n", buyPrice, iexQuote.IEXRealtimePrice)
		if buyPrice > iexQuote.IEXRealtimePrice && buyPrice*0.8 < iexQuote.IEXRealtimePrice {
			buyPrice = i.CalculatePriceWithThickSize(iexQuote.IEXRealtimePrice, 0)
			orderType = models.Limit
			durationType = models.GoodTillCancel
		}

		profitPrice := i.CalculatePriceWithThickSize(buyPrice, 100+percentProfit)
		stopLossPrice := i.CalculatePriceWithThickSize(buyPrice, percentLoss)
		distanceToMarket := i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 0)

		log.Println("Trying to buy", i.GetAssetType(), i.GetSymbol(), "for", buyPrice)
		or, err := t.Buy(
			i.GetOrder().
				WithType(orderType).
				WithAmount(1000).
				WithDuration(durationType).
				WithTakeProfit(profitPrice).
				WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05))
		if err != nil {
			log.Println(err)
			log.Println(models.GetStringError(err))
			continue
		}

		log.Println(or)
	}
}
