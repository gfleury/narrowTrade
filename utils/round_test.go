// +build !integration

package utils

import (
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

func (s *Suite) TestRound(c *check.C) {
	c.Assert(Round(1.23456, 1), check.Equals, 1.2)
	c.Assert(Round(1.23456, 2), check.Equals, 1.23)
	c.Assert(Round(1.23456, 3), check.Equals, 1.235)
	c.Assert(Round(1.23456, 4), check.Equals, 1.2346)
	c.Assert(Round(499.23456, 4), check.Equals, 499.2346)
	c.Assert(Round(499.23456, 2), check.Equals, 499.23)
	c.Assert(Round(499.23456, 3), check.Equals, 499.235)
	c.Assert(Round(499.23456, 4), check.Equals, 499.2346)

}

func (s *Suite) TestRoundNearestTick(c *check.C) {
	c.Assert(RoundNearestTick(446.4, 2, 0.1), check.Equals, 437.5)
	c.Assert(RoundNearestTick(446.4, 0, 0.1), check.Equals, 446.4)
	c.Assert(RoundNearestTick(446.467, 0, 0.01), check.Equals, 446.47)
	c.Assert(RoundNearestTick(446.467, 0, 0.001), check.Equals, 446.467)
	c.Assert(RoundNearestTick(446.467, 0, 0.005), check.Equals, 446.465)
	c.Assert(RoundNearestTick(446.468, 0, 0.005), check.Equals, 446.470)
	c.Assert(RoundNearestTick(446.468, 0, 0.0005), check.Equals, 446.468)
	c.Assert(RoundNearestTick(446.4682, 0, 0.0005), check.Equals, 446.468)

	c.Assert(RoundNearestTick(446.4682, -2, 0.1), check.Equals, 455.4)
	c.Assert(RoundNearestTick(446.4682, -4, 0.1), check.Equals, 464.3)
}

func (s *Suite) TestGetTickDecimals(c *check.C) {
	c.Assert(GetTickDecimals(0.1), check.Equals, 1)
	c.Assert(GetTickDecimals(0.01), check.Equals, 2)
	c.Assert(GetTickDecimals(0.001), check.Equals, 3)
	c.Assert(GetTickDecimals(0.001), check.Equals, 3)
	c.Assert(GetTickDecimals(0.005), check.Equals, 3)
	c.Assert(GetTickDecimals(0.0099999), check.Equals, 3)
}
