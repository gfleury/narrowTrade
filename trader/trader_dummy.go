package trader

import (
	"github.com/gfleury/narrowTrade/analysis"
	"github.com/gfleury/narrowTrade/models"
)

type DummyTrader struct {
	Trader
	InstrumentAnalyser analysis.InstrumentAnalyser
	ModeledAPI         *models.ModeledAPI
}

func (t *DummyTrader) placeOrder(order *models.Order) (*models.OrderResponse, error) {
	return &models.OrderResponse{}, nil
}

func (t *DummyTrader) getOpenOrders() float64 {
	return 0.0
}

func (t *DummyTrader) Buy(o *models.Order) (*models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(models.Buy))
}

func (t *DummyTrader) Sell(o *models.Order) (*models.OrderResponse, error) {
	return t.placeOrder(o.WithBuySell(models.Sell))
}

func (t *DummyTrader) Api() *models.ModeledAPI {
	return t.ModeledAPI
}

func (t *DummyTrader) Analyser() analysis.InstrumentAnalyser {
	return t.InstrumentAnalyser
}

func (t *DummyTrader) GetOrders() ([]string, []float64) {
	return []string{}, []float64{}
}

func (t *DummyTrader) Trade(param TradeParameter) error {
	return nil
}
