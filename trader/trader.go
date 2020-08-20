package trader

import (
	"fmt"
	"os"

	"github.com/gfleury/intensiveTrade/saxo_models"
)

type OrderOption int

const (
	DurationType OrderOption = iota
	AccountKey
	TakeProfit
	StopLoss
	OrderPrice
	StopLimitPrice
	TrailingStop
)

type OrderOptions struct {
	Type  OrderOption
	Value interface{}
}

type TrailingStopValues struct {
	TrailingStopDistanceToMarket float64
	TrailingStopStep             float64
	OrderPrice                   float64
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
	openOrders     *saxo_models.OrderList
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
	case OrderPrice:
		order.OrderPrice = op.Value.(float64)
	case TakeProfit:
		takeProfitOrder := *order
		takeProfitOrder.OrderType = saxo_models.Limit
		takeProfitOrder.BuySell = saxo_models.Sell
		takeProfitOrder.OrderDuration.DurationType = saxo_models.GoodTillCancel
		takeProfitOrder.OrderPrice = op.Value.(float64)
		order.Orders = append(order.Orders, takeProfitOrder)
	case StopLoss:
		stopLossOrder := *order
		stopLossOrder.OrderType = saxo_models.StopIfTraded
		stopLossOrder.BuySell = saxo_models.Sell
		stopLossOrder.OrderDuration.DurationType = saxo_models.GoodTillCancel
		stopLossOrder.OrderPrice = op.Value.(float64)
		order.Orders = append(order.Orders, stopLossOrder)
	case StopLimitPrice:
		order.StopLimitPrice = op.Value.(float64)
	case TrailingStop:
		values := op.Value.(TrailingStopValues)
		trailingStop := *order
		trailingStop.OrderType = saxo_models.TrailingStopIfTraded
		trailingStop.BuySell = saxo_models.Sell
		trailingStop.OrderDuration.DurationType = saxo_models.GoodTillCancel
		trailingStop.TrailingStopDistanceToMarket = values.TrailingStopDistanceToMarket
		trailingStop.TrailingStopStep = values.TrailingStopStep
		trailingStop.OrderPrice = values.OrderPrice
		order.Orders = append(order.Orders, trailingStop)
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

	bal, err := t.API.GetBalanceMe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	costOpenOrders := t.OpenOrders()

	if (r.EstimatedCashRequired*r.InstrumentToAccountConversionRate)+costOpenOrders >=
		float64(bal.InitialMargin.MarginAvailable) {
		order.GoodToGo = false
	}

	if !order.GoodToGo {
		return nil, fmt.Errorf("Order is not good to go, pre-check result: %s, estimated cost: %f, initial margin available: %d, %s: %s", r.PreCheckResult, r.EstimatedCashRequired, bal.InitialMargin.MarginAvailable, r.ErrorInfo.ErrorCode, r.ErrorInfo.Message)
	}

	or, err := t.API.Order(order)

	if err == nil {
		t.tradedOrdersID = append(t.tradedOrdersID, or)
	}

	return or, err
}

func (t *BasicSaxoTrader) OpenOrders() float64 {
	ol, err := t.API.GetOrdersMe()
	if err == nil {
		t.openOrders = ol
	}

	total := float64(0)

	if t.openOrders != nil {
		for _, oo := range t.openOrders.Data {
			if oo.BuySell == saxo_models.Buy {
				total += oo.Price * float64(oo.Amount)
			}
		}
	}

	return total
}
