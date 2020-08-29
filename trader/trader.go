package trader

import (
	"fmt"
	"log"

	"github.com/gfleury/narrowTrade/models"
	iex "github.com/goinvest/iexcloud/v2"
)

type Trader interface {
	Api() *models.ModeledAPI
	Buy(order *models.Order) (*models.OrderResponse, error)
	Sell(order *models.Order) (*models.OrderResponse, error)
	getOpenOrders() float64
	IEXApi() *iex.Client
}

type ComplexTrader interface {
	Trader
	Trade(param TradeParameter) error
	GetOrders() ([]string, []float64)
}

type TradeParameter struct {
	Watchlist                               string
	Symbols                                 []int
	PercentLoss, PercentProfit, TotalInvest float64
}

type BasicSaxoTrader struct {
	Trader
	AccountKey     string
	ModeledAPI     *models.ModeledAPI
	IEXAPI         *iex.Client
	tradedOrdersID []*models.OrderResponse
	openOrders     *models.OrderList
}

func (t *BasicSaxoTrader) placeOrder(order *models.Order) (*models.OrderResponse, error) {
	// Define AccontKey from trader in all Orders
	order.AccountKey = t.AccountKey
	for idx := range order.Orders {
		order.Orders[idx].AccountKey = t.AccountKey
	}

	r, err := t.ModeledAPI.PreOrder(order)
	if err != nil {
		return nil, err
	}

	bal, err := t.ModeledAPI.GetBalanceMe()
	if err != nil {
		return nil, err
	}

	costOpenOrders := t.getOpenOrders()

	if (r.EstimatedCashRequired*r.InstrumentToAccountConversionRate)+costOpenOrders >=
		float64(bal.InitialMargin.MarginAvailable) {
		order.GoodToGo = false
	}

	if !order.GoodToGo {
		return nil, fmt.Errorf("Order is not good to go, pre-check result: %s, estimated cost: %f, initial margin available: %f, %s: %s", r.PreCheckResult, r.EstimatedCashRequired, bal.InitialMargin.MarginAvailable, r.ErrorInfo.ErrorCode, r.ErrorInfo.Message)
	}

	or, err := t.ModeledAPI.Order(order)

	if err == nil {
		t.tradedOrdersID = append(t.tradedOrdersID, or)
	}

	return or, err
}

func (t *BasicSaxoTrader) getOpenOrders() float64 {
	ol, err := t.ModeledAPI.GetOrdersMe()
	if err == nil {
		t.openOrders = ol
	} else {
		log.Println("Unable to GetOrdersMe:", err)
	}

	total := float64(0)

	if t.openOrders != nil {
		for _, oo := range t.openOrders.Data {
			if oo.BuySell == models.Buy {
				price, err := t.ModeledAPI.GetInstrumentPrice(
					&models.SaxoInstrument{Identifier: oo.Uic, AssetType: "Stock"})
				if err != nil {
					continue
				}
				total += price * float64(oo.Amount)
			}
		}
	}

	return total
}

func (t *BasicSaxoTrader) Buy(o *models.Order) (*models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(models.Buy))
}

func (t *BasicSaxoTrader) Sell(o *models.Order) (*models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(models.Sell))
}

func (t *BasicSaxoTrader) Api() *models.ModeledAPI {
	return t.ModeledAPI
}

func (t *BasicSaxoTrader) IEXApi() *iex.Client {
	return t.IEXAPI
}
