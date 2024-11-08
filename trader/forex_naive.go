package trader

// trade/v1/prices/subscriptions get stream of prices

import (
	"log"
	"math"

	"github.com/gfleury/narrowTrade/config"
	"github.com/gfleury/narrowTrade/models"
)

type ForexNaive struct {
	*StockNaive
}

func (t *ForexNaive) Trade(param TradeParameter) error {

	t.data = t.createStocksNaive(param.Symbols, param.PercentProfit)

	err := t.UpdateAvailableCash(param.TotalInvest)
	if err != nil {
		return err
	}

	qntSymbols := len(t.data)

	cashPerSymbol := t.availableCash.GetSplit(qntSymbols)

	log.Printf("Using %f per symbol, totalzing %f\n", cashPerSymbol, t.availableCash.GetTotal())

	failedTrades := 0
	for idx, n := range t.data {
		i := n.instrument

		if i == nil {
			log.Println("Instrument is nil, something is wrong:", n)
			failedTrades++
			cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
			continue
		}

		orderType := models.OrderType(models.Market)

		durationType := models.DurationType(models.DayOrder)

		buyPrice := 0.0

		if buyPrice <= 0 {
			log.Println("Calculated BuyPrice below/equal 0:", buyPrice)
			failedTrades++
			cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
			continue
		}

		amount := int(math.Ceil(cashPerSymbol / buyPrice))
		if amount < 1 || buyPrice > cashPerSymbol || buyPrice*float64(amount) < config.MINIMUM_TRADE_VALUE {
			log.Printf("Not enough available money to buy OR trade value too low %s %s for %f", i.GetAssetType(), i.GetSymbol(), buyPrice*float64(amount))
			failedTrades++
			cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
			continue
		}

		log.Printf("Trying to buy %s %s for %d x %f",
			i.GetAssetType(), i.GetSymbol(), amount, buyPrice)

		or, err := t.Buy(
			i.GetOrder().
				WithType(orderType).
				WithAmount(amount).
				WithPrice(buyPrice).
				WithDuration(durationType))
		if err != nil {
			orderError := models.GetOrderError(err)
			if orderError != nil && models.BusinessRuleViolation(orderError) {
				failedTrades++
				cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
				continue
			}
			return err
		}

		// Save Ordered price for next calculation
		or.TotalPrice = float64(amount) * buyPrice
		log.Println(or)

		// rebalance cashPerSymbol based on remaining cash
		t.availableCash.Spent(float64(amount) * buyPrice)
		cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
	}
	return nil
}
