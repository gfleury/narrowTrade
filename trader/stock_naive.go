package trader

import (
	"fmt"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/utils"

	iex "github.com/goinvest/iexcloud/v2"
)

func (t *BasicSaxoTrader) BuyStocksNaive(buyingInstruments []models.Instrument, iexQuotes []iex.Quote) {

	for idx, i := range buyingInstruments {
		if i == nil {
			continue
		}

		iexQuote := iexQuotes[idx]
		fmt.Println("Trying to get price Saxo Bank price", i.GetAssetType(), i.GetSymbol())
		buyPrice, err := t.API.GetInstrumentPrice(i)
		if err != nil {
			fmt.Println(err)
			fmt.Println(models.GetStringError(err))
			continue
		}

		orderType := models.OrderType(models.Market)

		durationType := models.DurationType(models.DayOrder)

		fmt.Printf("Got price %f - %f\n", buyPrice, iexQuote.IEXRealtimePrice)
		if buyPrice > iexQuote.IEXRealtimePrice && buyPrice*0.8 < iexQuote.IEXRealtimePrice {
			buyPrice = iexQuote.IEXRealtimePrice
			orderType = models.Limit
			durationType = models.GoodTillCancel
		}

		fmt.Println("Trying to buy", i.GetAssetType(), i.GetSymbol(), "for", buyPrice)
		or, err := t.Buy(
			i.GetOrder().
				WithType(orderType).
				WithAmount(1000).
				WithDuration(durationType).
				WithStopLoss(utils.Round(buyPrice*0.97, 2)).
				WithTakeProfit(utils.Round(buyPrice*1.03, 2)))
		if err != nil {
			fmt.Println(err)
			fmt.Println(models.GetStringError(err))
			continue
		}

		fmt.Println(or)
	}
}
