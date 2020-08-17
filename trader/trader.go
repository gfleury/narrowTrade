package trader

import (
	"fmt"

	"github.com/gfleury/intensiveTrade/saxo_models"
)

type OrderOption int

const (
	DurationType OrderOption = iota
	AccountKey
	TakeProfit
	StopLoss
)

type OrderOptions struct {
	Type  OrderOption
	Value interface{}
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

func NewOrderOption(opt OrderOption, v interface{}) OrderOptions {
	return OrderOptions{
		Type:  opt,
		Value: v,
	}
}

func (op *OrderOptions) ApplyOption(order *saxo_models.Order) error {
	switch op.Type {
	case DurationType:
		order.OrderDuration.DurationType = saxo_models.DurationType(op.Value.(string))
	case AccountKey:
		order.AccountKey = op.Value.(string)
	case TakeProfit:
		takeProfitOrder := *order
		takeProfitOrder.OrderType = saxo_models.Limit
		takeProfitOrder.BuySell = saxo_models.Sell
		takeProfitOrder.OrderDuration.DurationType = saxo_models.GoodTillCancel
		takeProfitOrder.OrderPrice = op.Value.(float64)
		order.Orders = append(order.Orders, takeProfitOrder)
	case StopLoss:
		stopLossOrder := *order
		stopLossOrder.OrderType = saxo_models.Stop
		stopLossOrder.BuySell = saxo_models.Sell
		stopLossOrder.OrderDuration.DurationType = saxo_models.GoodTillCancel
		stopLossOrder.OrderPrice = op.Value.(float64)
		order.Orders = append(order.Orders, stopLossOrder)
	}
	return nil
}

func (t *BasicSaxoTrader) Buy(i saxo_models.Instrument, ot saxo_models.OrderType, amount int, opts ...OrderOptions) (*saxo_models.OrderResponse, error) {
	return t.placeOrder(i, saxo_models.Buy, ot, amount, opts...)
}

func (t *BasicSaxoTrader) Sell(i saxo_models.Instrument, ot saxo_models.OrderType, amount int, opts ...OrderOptions) (*saxo_models.OrderResponse, error) {
	return t.placeOrder(i, saxo_models.Sell, ot, amount, opts...)
}

func (t *BasicSaxoTrader) placeOrder(i saxo_models.Instrument, bs saxo_models.BuySell, ot saxo_models.OrderType, amount int, opts ...OrderOptions) (*saxo_models.OrderResponse, error) {
	order := i.GetOrder(bs, ot, amount)

	order.OrderDuration.DurationType = saxo_models.DayOrder
	order.AccountKey = t.AccountKey

	for _, opt := range opts {
		err := opt.ApplyOption(order)
		if err != nil {
			return nil, err
		}
	}

	r, err := t.API.PreOrder(order)
	if err != nil {
		return nil, err
	}

	if !order.GoodToGo {
		return nil, fmt.Errorf("Order is not good to go, pre-check result: %s, estimated cost: %d", r.PreCheckResult, r.EstimatedCashRequired)
	}

	or, err := t.API.Order(order)

	if err == nil {
		t.tradedOrdersID = append(t.tradedOrdersID, or)
	}

	return or, err
}
