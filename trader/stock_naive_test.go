// +build integration
// +build !unit

package trader

import (
	"github.com/gfleury/narrowTrade/analysis"
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
