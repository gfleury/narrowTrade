package saxo_models

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

// TestGetPriceTick tests the instrument tick based on the order price
// https://www.developer.saxo/openapi/learn/order-placement#TickSizeandTickSizeSchemes
func (s *Suite) TestGetPriceTick(c *check.C) {
	e := []Elements{
		{0.4999, 0.0001},
		{0.9995, 0.0005},
		{4.999, 0.001},
		{9.995, 0.005},
		{49.99, 0.01},
		{99.95, 0.05},
		{499.9, 0.1},
		{999.5, 0.5},
		{4999.0, 1.0},
		{9995.0, 5.0},
	}

	tick := GetPriceTick(446.4, e)
	c.Assert(tick, check.Equals, 0.1)

	tick = GetPriceTick(3.8, e)
	c.Assert(tick, check.Equals, 0.001)

	tick = GetPriceTick(499.9, e)
	c.Assert(tick, check.Equals, 0.1)

	tick = GetPriceTick(500, e)
	c.Assert(tick, check.Equals, 0.5)

	tick = GetPriceTick(9999, e)
	c.Assert(tick, check.Equals, 5.0)

	e = []Elements{
		{0.9999, 0.0001},
	}

	tick = GetPriceTick(100, e)
	c.Assert(tick, check.Equals, 0.0001)
}

// TestCalculatePriceWithThickSize tests the Price rounding rule based on the instrument tick
func (s *Suite) TestCalculatePriceWithThickSize(c *check.C) {
	id := &SaxoInstrumentDetails{
		Format: Format{
			Decimals:      2,
			OrderDecimals: 2,
		},
		TickSizeScheme: TickSizeScheme{
			DefaultTickSize: 0.005,
			Elements: []Elements{
				{0.4999, 0.0001},
				{0.9995, 0.0005},
				{4.999, 0.001},
				{9.995, 0.005},
				{49.99, 0.01},
				{99.95, 0.05},
				{499.9, 0.1},
				{999.5, 0.5},
				{4999.0, 1.0},
				{9995.0, 5.0},
			},
		},
	}

	price := id.CalculatePriceWithThickSize(446.4, 2)
	c.Assert(price, check.Equals, 437.5)
}
