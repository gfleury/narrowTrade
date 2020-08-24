package trader

import (
	"github.com/gfleury/narrowTrade/tests"
	check "gopkg.in/check.v1"
)

func (s *Suite) TestBuyStocksNaive_2_5(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
		IEXClient:  tests.GetIEXSandboxClient(),
	}

	err := t.BuyStocksNaive([]string{"AAPL"}, 2, 5)
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestBuyStocksNaive_1_2(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
		IEXClient:  tests.GetIEXSandboxClient(),
	}

	err := t.BuyStocksNaive([]string{"AAPL"}, 1, 2)
	c.Assert(err, check.IsNil)
}
