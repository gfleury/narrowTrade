// +build !integration
// +build unit

package portfolio

import (
	"strings"
	"testing"

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
			},
			Investment{
				WatchlistID:     "",
				ValuePercentage: 10.0,
				Trader:          ForexNaive,
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
			},
			Investment{
				WatchlistID:     "",
				ValuePercentage: 10.0,
				Trader:          ForexNaive,
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
- valuepercentage: 10
  trader: ForexNaive
`)
	p := &Portfolio{}

	err := p.Load(data)
	c.Assert(err, check.IsNil)

	err = p.Validate()
	c.Assert(err, check.IsNil)

}
