package external_lists

import (
	"testing"

	check "gopkg.in/check.v1"
)

type ExternalListsSuite struct {
}

func (s *ExternalListsSuite) SetUpSuite(c *check.C) {
}

func (s *ExternalListsSuite) TearDownSuite(c *check.C) {
}

func (s *ExternalListsSuite) SetUpTest(c *check.C) {
}

var _ = check.Suite(&ExternalListsSuite{})

func Test(t *testing.T) {
	check.TestingT(t)
}

func (s *ExternalListsSuite) TestWikipedia_sp100_getsymbols(c *check.C) {
	ws := NewWikipediaSP100()

	symbols, err := ws.GetSymbols()

	c.Assert(err, check.IsNil)
	c.Assert(len(symbols), check.Equals, 101)
}

func (s *ExternalListsSuite) TestWikipedia_sp500_getsymbols(c *check.C) {
	ws := NewWikipediaSP500()

	symbols, err := ws.GetSymbols()

	c.Assert(err, check.IsNil)
	c.Assert(len(symbols), check.Equals, 505)
}
