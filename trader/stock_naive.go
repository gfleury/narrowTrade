package trader

import (
	"fmt"
	"time"

	"github.com/gfleury/intensiveTrade/saxo_models"
	"github.com/gfleury/intensiveTrade/utils"

	iex "github.com/goinvest/iexcloud/v2"
)

func (t *BasicSaxoTrader) BuyStocksNaive(buyingInstruments []saxo_models.Instrument, iexQuotes []iex.Quote) {

	for idx, i := range buyingInstruments {
		time.Sleep(time.Second)
		iexQuote := iexQuotes[idx]
		fmt.Println("Trying to get price Saxo Bank price", i.GetAssetType(), i.GetSymbol())
		buyPrice, err := t.API.GetInstrumentPrice(i)
		if err != nil {
			fmt.Println(err)
			fmt.Println(saxo_models.GetStringError(err))
			continue
		}

		opts := []OrderOptions{}

		orderType := saxo_models.OrderType(saxo_models.Market)

		durationType := saxo_models.DayOrder

		fmt.Printf("Got price %f - %f\n", buyPrice, iexQuote.IEXRealtimePrice)
		if buyPrice > iexQuote.IEXRealtimePrice {
			buyPrice = iexQuote.IEXRealtimePrice
			orderType = saxo_models.Limit
			opts = append(opts, NewOrderOption(OrderPrice, utils.Round(buyPrice)))
			durationType = saxo_models.GoodTillCancel
		}

		opts = append(opts, NewOrderOption(DurationType, durationType))
		opts = append(opts, NewOrderOption(TakeProfit, utils.Round(buyPrice*1.03)))
		opts = append(opts, NewOrderOption(StopLoss, utils.Round(buyPrice*0.97)))

		fmt.Println("Trying to buy", i.GetAssetType(), i.GetSymbol(), "for", buyPrice)
		or, err := t.Buy(i, orderType, 1000, opts...)

		if err != nil {
			fmt.Println(err)
			fmt.Println(saxo_models.GetStringError(err))
			continue
		}

		fmt.Println(or)
	}
}
