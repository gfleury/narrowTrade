// +build integration
// +build !unit

package trader

import (
	"github.com/gfleury/narrowTrade/analysis"
	check "gopkg.in/check.v1"
)

func (s *Suite) TestBuyForexNaive_5_6(c *check.C) {
	t := ForexNaive{
		StockNaive: &StockNaive{
			BasicSaxoTrader: &BasicSaxoTrader{
				AccountKey:         s.acc.GetAccountKeyMe(),
				ModeledAPI:         s.ma,
				InstrumentAnalyser: &analysis.DummyAnalyser{},
			},
		},
	}

	// Buy USDEUR 5/6 ticks
	err := t.Trade(TradeParameter{"", []int{22}, 1, 5, 6, 0})
	c.Assert(err, check.IsNil)
}
