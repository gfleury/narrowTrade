package trader

import (
	"log"

	"github.com/gfleury/narrowTrade/models"

	iex "github.com/goinvest/iexcloud/v2"
)

type NaiveStockData struct {
	instrument models.InstrumentDetails
	iexQuote   iex.Quote
}

func (t *BasicSaxoTrader) BuyStocksNaive(symbols []string, percentLoss, percentProfit float64) error {

	naiveStockData := t.createStocksNaive(symbols)

	for _, n := range naiveStockData {
		iexQuote := n.iexQuote
		i := n.instrument

		log.Println("Trying to get price Saxo Bank price", i.GetAssetType(), i.GetSymbol())
		buyPrice, err := t.ModeledAPI.GetInstrumentPrice(i)
		if err != nil {
			return err
		}

		orderType := models.OrderType(models.Market)

		durationType := models.DurationType(models.DayOrder)

		log.Printf("Got price %f - %f\n", buyPrice, iexQuote.IEXRealtimePrice)
		if buyPrice > iexQuote.IEXRealtimePrice && buyPrice*0.8 < iexQuote.IEXRealtimePrice {
			buyPrice = i.CalculatePriceWithThickSize(iexQuote.IEXRealtimePrice, 0)
			orderType = models.Limit
			durationType = models.GoodTillCancel
		}

		profitPrice := i.CalculatePriceWithThickSize(buyPrice, -percentProfit)
		stopLossPrice := i.CalculatePriceWithThickSize(buyPrice, percentLoss)
		distanceToMarket := i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 0)

		log.Printf("Trying to buy %s %s for %f with stopLoss %f (%f) and takeProfit %f\n",
			i.GetAssetType(), i.GetSymbol(), buyPrice, stopLossPrice, distanceToMarket, profitPrice)

		or, err := t.Buy(
			i.GetOrder().
				WithType(orderType).
				WithAmount(10).
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
	naiveStockData := make([]NaiveStockData, len(symbols))
	for idx, symbol := range symbols {
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
		naiveStockData[idx] = n
	}

	return naiveStockData
}
