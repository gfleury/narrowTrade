package trader

import (
	"fmt"
	"os"

	"github.com/gfleury/intensiveTrade/saxo_models"
)

type Trader interface {
	Buy(order saxo_models.Order) (*saxo_models.OrderResponse, error)
	Sell(order saxo_models.Order) (*saxo_models.OrderResponse, error)
}

type BasicSaxoTrader struct {
	Trader
	AccountKey     string
	API            *saxo_models.ModeledAPI
	tradedOrdersID []*saxo_models.OrderResponse
	openOrders     *saxo_models.OrderList
}

func (t *BasicSaxoTrader) Buy(o *saxo_models.Order) (*saxo_models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(saxo_models.Buy))
}

func (t *BasicSaxoTrader) Sell(o *saxo_models.Order) (*saxo_models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(saxo_models.Sell))
}

func (t *BasicSaxoTrader) placeOrder(order *saxo_models.Order) (*saxo_models.OrderResponse, error) {
	// Define AccontKey from trader in all Orders
	order.AccountKey = t.AccountKey
	for idx := range order.Orders {
		order.Orders[idx].AccountKey = t.AccountKey
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
