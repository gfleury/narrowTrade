package trader

import (
	"log"
	"math"
	"sort"
	"time"

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
	historicalPrices []float64
	priceRecoveries  int
}

func (i *InstrumentNaiveData) CalculateCrosses(expectedPriceRecovery float64) {
	log.Println("Calculating crosses for", i.instrument.GetSymbol(), expectedPriceRecovery, len(i.historicalPrices))
	if i.historicalPrices == nil {
		return
	}

	for index, price := range i.historicalPrices {
		expectedPrice := price * (1 + expectedPriceRecovery)

		if index+1 < len(i.historicalPrices) && i.historicalPrices[index+1] >= expectedPrice {
			i.priceRecoveries++
			log.Println(i.instrument.GetSymbol(), "Recovered from previous day", price, i.historicalPrices[index+1], ">=", expectedPrice)
		}
	}
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

	t.data = t.createStocksNaive(param.Symbols, param.PercentProfit)

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

		if buyPrice <= 0 {
			log.Println("Calculated BuyPrice below/equal 0:", buyPrice)
			failedTrades++
			cashPerSymbol = t.availableCash.GetSplit(qntSymbols - (idx + failedTrades))
			continue
		}

		// indicator unavailable, go with percentage
		profitPrice = i.CalculatePriceWithThickSize(buyPrice, -param.PercentProfit)
		if param.PercentLoss > 0 {
			stopLossPrice = i.CalculatePriceWithThickSize(buyPrice, param.PercentLoss)
		} else {
			stopLossPrice = 0
		}
		distanceToMarket = i.CalculatePriceWithThickSize(buyPrice-stopLossPrice, 70)

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

		order := i.GetOrder().
			WithType(orderType).
			WithAmount(amount).
			WithPrice(buyPrice).
			WithDuration(durationType).
			WithTakeProfit(profitPrice)

		if stopLossPrice > 0 {
			order = order.WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05)
		}

		or, err := t.Buy(order)
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

func (t *StockNaive) createStocksNaive(uics []int, percentProfit float64) []InstrumentNaiveData {
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
		n.historicalPrices, err = t.Analyser().HistoricalPrices(i)
		if err != nil {
			log.Println("Failed to get historical prices, going with percentage", i.GetSymbolSimple(), err)
		}

		n.CalculateCrosses(percentProfit / 100)

		naiveStockData = append(naiveStockData, n)
	}

	sort.Slice(naiveStockData, func(i, j int) bool {
		return naiveStockData[i].priceRecoveries > naiveStockData[j].priceRecoveries
	})

	return naiveStockData
}
