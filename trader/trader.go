package trader

import (
	"fmt"
	"os"

	"github.com/gfleury/narrowTrade/models"
)

type Trader interface {
	Buy(order models.Order) (*models.OrderResponse, error)
	Sell(order models.Order) (*models.OrderResponse, error)
}

type BasicSaxoTrader struct {
	Trader
	AccountKey     string
	API            *models.ModeledAPI
	tradedOrdersID []*models.OrderResponse
	openOrders     *models.OrderList
}

func (t *BasicSaxoTrader) Buy(o *models.Order) (*models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(models.Buy))
}

func (t *BasicSaxoTrader) Sell(o *models.Order) (*models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(models.Sell))
}

func (t *BasicSaxoTrader) placeOrder(order *models.Order) (*models.OrderResponse, error) {
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
			if oo.BuySell == models.Buy {
				total += oo.Price * float64(oo.Amount)
			}
		}
	}

	return total
}
