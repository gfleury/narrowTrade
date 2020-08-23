package trader

import (
	"github.com/gfleury/narrowTrade/tests"
	check "gopkg.in/check.v1"
)

func (s *Suite) TestBuyStocksNaive(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
		IEXClient:  tests.GetIEXSandboxClient(),
	}

	err := t.BuyStocksNaive([]string{"AAPL"}, 2, 5)
	c.Assert(err, check.IsNil)
}
