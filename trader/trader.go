package trader

import (
	"fmt"

	"github.com/gfleury/intensiveTrade/saxo_models"
)

type OrderOptions struct {
	Name  string
	Value string
}

type Trader interface {
	Buy(i saxo_models.Instrument, t saxo_models.OrderType, amount int, op ...OrderOptions) error
	Sell(i saxo_models.Instrument, t saxo_models.OrderType, amount int, op ...OrderOptions) error
}

type BasicSaxoTrader struct {
	Trader
	AccountKey     string
	API            *saxo_models.ModeledAPI
	tradedOrdersID []*saxo_models.OrderResponse
}

func (t *BasicSaxoTrader) Buy(i saxo_models.Instrument, ot saxo_models.OrderType, amount int, opts ...OrderOptions) (*saxo_models.OrderResponse, error) {
	return t.placeOrder(i, saxo_models.Buy, ot, amount, opts...)
}

func (t *BasicSaxoTrader) Sell(i saxo_models.Instrument, ot saxo_models.OrderType, amount int, opts ...OrderOptions) (*saxo_models.OrderResponse, error) {
	return t.placeOrder(i, saxo_models.Sell, ot, amount, opts...)
}

func (t *BasicSaxoTrader) placeOrder(i saxo_models.Instrument, bs saxo_models.BuySell, ot saxo_models.OrderType, amount int, op ...OrderOptions) (*saxo_models.OrderResponse, error) {
	order := i.GetOrder(bs, ot, amount)

	order.AccountKey = t.AccountKey
	order.OrderDuration.DurationType = saxo_models.DayOrder

	r, err := t.API.PreOrder(order)
	if err != nil {
		return nil, err
	}

	if !order.GoodToGo {
		return nil, fmt.Errorf("Order is not good to go, pre-check result: %s, estimated cost: %d", r.PreCheckResult, r.EstimatedCashRequired)
	}

	or, err := t.API.Order(order)

	t.tradedOrdersID = append(t.tradedOrdersID, or)

	return or, err
}
