//go:build integration && !unit
// +build integration,!unit

package trader

import (
	"github.com/gfleury/narrowTrade/analysis"
	"github.com/gfleury/narrowTrade/models"
	check "gopkg.in/check.v1"
)

func (s *Suite) TestBuyStocksNaive_2_5(c *check.C) {
	t := StockNaive{
		&BasicSaxoTrader{
			AccountKey:         s.acc.GetAccountKeyMe(),
			ModeledAPI:         s.ma,
			InstrumentAnalyser: &analysis.DummyAnalyser{},
		},
		nil,
		nil,
	}

	// Buy Apple and Intel
	err := t.Trade(TradeParameter{"", []int{211, 247}, 2, 2, 5, 0})
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestBuyStocksNaive_1_2(c *check.C) {
	t := &StockNaive{
		&BasicSaxoTrader{
			AccountKey:         s.acc.GetAccountKeyMe(),
			ModeledAPI:         s.ma,
			InstrumentAnalyser: &analysis.DummyAnalyser{},
		},
		nil,
		nil,
	}

	// Buy Apple and Intel
	err := t.Trade(TradeParameter{"", []int{211, 247}, 2, 1, 2, 0})
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestCalculateCrosses_NoHistoricalPrices(c *check.C) {
	data := &InstrumentNaiveData{
		instrument:       models.InstrumentDetails{Symbol: "TEST"},
		historicalPrices: nil,
		priceRecoveries:  0,
	}

	data.CalculateCrosses(0.05)
	c.Assert(data.priceRecoveries, check.Equals, 0)
}

func (s *Suite) TestCalculateCrosses_NoRecoveries(c *check.C) {
	data := &InstrumentNaiveData{
		instrument:       models.InstrumentDetails{Symbol: "TEST"},
		historicalPrices: []float64{100, 95, 90, 85},
		priceRecoveries:  0,
	}

	data.CalculateCrosses(0.05)
	c.Assert(data.priceRecoveries, check.Equals, 0)
}

func (s *Suite) TestCalculateCrosses_WithRecoveries(c *check.C) {
	data := &InstrumentNaiveData{
		instrument:       models.InstrumentDetails{Symbol: "TEST"},
		historicalPrices: []float64{100, 105, 110, 115},
		priceRecoveries:  0,
	}

	data.CalculateCrosses(0.05)
	c.Assert(data.priceRecoveries, check.Equals, 3)
}

func (s *Suite) TestCalculateCrosses_MixedRecoveries(c *check.C) {
	data := &InstrumentNaiveData{
		instrument:       models.InstrumentDetails{Symbol: "TEST"},
		historicalPrices: []float64{100, 105, 90, 95, 110},
		priceRecoveries:  0,
	}

	data.CalculateCrosses(0.05)
	c.Assert(data.priceRecoveries, check.Equals, 2)
}
