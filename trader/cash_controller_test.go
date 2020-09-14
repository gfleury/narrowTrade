// +build !integration
// +build unit

package trader

import (
	"testing"

	"github.com/gfleury/narrowTrade/utils"
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

func (s *Suite) TestAvailableCash(c *check.C) {
	ac := &AvailableCash{}

	ac.SetAvailableCash(10000.00)

	c.Assert(ac.GetTotal(), check.Equals, 10000.00)

	c.Assert(ac.GetSplit(10), check.Equals, 1000.00)

	ac.Spent(1938.00)

	for i := 10; i > 0; i-- {
		c.Assert(utils.Round(ac.GetSplit(i), 1), check.Equals, 806.2)
		ac.Spent(806.2)
	}

}
