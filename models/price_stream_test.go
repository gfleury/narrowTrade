// +build integration
// +build !unit

package models

import (
	check "gopkg.in/check.v1"
)

func (s *ModeledApiSuite) TestcreateSubscriptioWebSocket(c *check.C) {
	mc, wsC, err := s.ma.CreatePriceStream(&SaxoInstrument{Identifier: 22, AssetType: FxSpot})
	c.Assert(err, check.IsNil)
	c.Assert(wsC, check.NotNil)
	defer wsC.Close()
	c.Assert(mc, check.NotNil)
	msg := <-mc
	c.Assert(msg, check.NotNil)
	c.Assert(int(msg.PayloadSize), check.Equals, len(msg.Payload))
	c.Assert(int(msg.RIDSize), check.Equals, len(msg.RID))
}
