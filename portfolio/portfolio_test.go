// +build !integration
// +build unit

package portfolio

import (
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/gfleury/narrowTrade/trader"
	check "gopkg.in/check.v1"
)

type Suite struct {
}

func (s *Suite) SetUpSuite(c *check.C) {
}

func (s *Suite) TearDownSuite(c *check.C) {
}

func (s *Suite) SetUpTest(c *check.C) {
}

var _ = check.Suite(&Suite{})

func Test(t *testing.T) {
	check.TestingT(t)
}

func (s *Suite) TestPortfolioSave(c *check.C) {
	p := &Portfolio{
		Distribution: []Investment{
			Investment{
				WatchlistID:     "2491746",
				ValuePercentage: 90.0,
				Trader:          StockNaive,
			},
			Investment{
				WatchlistID:     "",
				ValuePercentage: 10.0,
				Trader:          ForexNaive,
			},
		},
	}

	c.Assert(p, check.NotNil)

	err := p.Save()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestPortfolioValidate(c *check.C) {
	p := &Portfolio{
		Distribution: []Investment{
			Investment{
				WatchlistID:     "2491746",
				ValuePercentage: 90.0,
				Trader:          StockNaive,
				Parameters: trader.TradeParameter{
					PercentLoss:   1,
					PercentProfit: 3,
				},
			},
			Investment{
				Symbols:         []string{"USDAUS"},
				ValuePercentage: 10.0,
				Trader:          ForexNaive,
				Parameters: trader.TradeParameter{
					PercentLoss:   1,
					PercentProfit: 3,
				},
			},
		},
	}
	err := p.Validate()
	c.Assert(err, check.IsNil)

	p = &Portfolio{
		Distribution: []Investment{
			Investment{
				WatchlistID:     "2491746",
				ValuePercentage: 90.1,
				Trader:          StockNaive,
				Parameters: trader.TradeParameter{
					PercentLoss:   1,
					PercentProfit: 3,
				},
			},
			Investment{
				Symbols:         []string{"USDAUS"},
				ValuePercentage: 10.0,
				Trader:          ForexNaive,
				Parameters: trader.TradeParameter{
					PercentLoss:   1,
					PercentProfit: 3,
				},
			},
		},
	}
	err = p.Validate()
	c.Assert(err.Error(), check.Matches, "Portifolio investments can't be more.*")

	p = &Portfolio{
		Distribution: []Investment{
			Investment{
				WatchlistID: "2491746",
				Trader:      StockNaive,
			},
		},
	}
	err = p.Validate()
	c.Assert(err.Error(), check.Matches, "Investment.*")

}

func (s *Suite) TestPortfolioLoad(c *check.C) {
	data := strings.NewReader(`
distribution:
- watchlistid: "2491746"
  valuepercentage: 90
  trader: StockNaive
  parameters:
    percentprofit: 5
    percentloss: 2
- valuepercentage: 10
  trader: ForexNaive
  symbols:
  - USDCAD
  parameters:
    percentprofit: 5
    percentloss: 2
`)
	p := &Portfolio{}

	err := p.Load(data)
	c.Assert(err, check.IsNil)

	err = p.Validate()
	c.Assert(err, check.IsNil)

}

func (s *Suite) TestPortfolioRebalance(c *check.C) {
	data := strings.NewReader(`
distribution:
- watchlistid: "2491746"
  valuepercentage: 90
  trader: StockNaive
  parameters:
    percentprofit: 5
    percentloss: 2
- valuepercentage: 10
  trader: ForexNaive
  symbols:
  - USDCAD
  parameters:
    percentprofit: 5
    percentloss: 2
`)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	p := &Portfolio{
		Traders: map[TraderName]trader.ComplexTrader{
			StockNaive: &trader.DummyTrader{
				ModeledAPI: &models.ModeledAPI{
					Ctx:    ctx,
					Client: saxo_openapi.NewAPIClient(saxo_openapi.NewConfiguration()),
				},
			},
		},
	}

	err := p.Load(data)
	c.Assert(err, check.IsNil)

	err = p.Rebalance()
	c.Assert(err, check.NotNil)

}

func (s *Suite) TestPortfolioSP100Rebalance(c *check.C) {
	data := strings.NewReader(`
distribution:
- externallist: "SP100"
  valuepercentage: 90
  trader: StockNaive
  parameters:
    percentprofit: 5
    percentloss: 2
- valuepercentage: 10
  trader: ForexNaive
  symbols:
  - USDCAD
  parameters:
    percentprofit: 5
    percentloss: 2
`)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	p := &Portfolio{
		Traders: map[TraderName]trader.ComplexTrader{
			StockNaive: &trader.DummyTrader{
				ModeledAPI: &models.ModeledAPI{
					Ctx:    ctx,
					Client: saxo_openapi.NewAPIClient(saxo_openapi.NewConfiguration()),
				},
			},
		},
	}

	err := p.Load(data)
	c.Assert(err, check.IsNil)

	err = p.Rebalance()
	c.Assert(err, check.NotNil)

}
