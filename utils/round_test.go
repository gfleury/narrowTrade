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
