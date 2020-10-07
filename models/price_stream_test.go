// +build integration
// +build !unit

package models

import (
	"time"

	check "gopkg.in/check.v1"
)

func (s *ModeledApiSuite) TestcreateSubscriptioWebSocket(c *check.C) {
	mc, wsC, err := s.ma.CreatePriceStream(&SaxoInstrument{Identifier: 22, AssetType: FxSpot})
	c.Assert(err, check.IsNil)
	c.Assert(wsC, check.NotNil)
	defer wsC.Close()
	c.Assert(mc, check.NotNil)
	quote := <-mc
	c.Assert(quote, check.NotNil)
	// t, err := time.Parse(time.RFC3339, quote.LastUpdated)
	// c.Assert(err, check.IsNil)
	if time.Since(quote.LastUpdated) > 10*time.Minute {
		c.Errorf("Lastupdate more than 10m behind")
	}
	if quote.Quote.Ask < 0 {
		c.Errorf("quote.Ask lower or equal to zero")
	}
	if quote.Quote.Bid < 0 {
		c.Errorf("quote.Bid lower or equal to zero")
	}
	if quote.Quote.Mid < 0 {
		c.Errorf("quote.Mid lower or equal to zero")
	}
}
