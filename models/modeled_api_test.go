// +build integration

package models

import (
	"context"

	"github.com/antihax/optional"
	"github.com/gfleury/narrowTrade/saxo_oauth2"
	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/gfleury/narrowTrade/tests"
	"golang.org/x/oauth2"
	check "gopkg.in/check.v1"

	"testing"
)

type ModeledApiSuite struct {
	acc         *Accounts
	ma          *ModeledAPI
	tokenSource oauth2.TokenSource
}

func (s *ModeledApiSuite) SetUpSuite(c *check.C) {
	ctx := context.Background()

	oauth2cfg := tests.GetTestOauth()

	token, err := saxo_oauth2.GetToken(ctx, oauth2cfg)
	c.Assert(err, check.IsNil)

	s.tokenSource = oauth2cfg.TokenSource(ctx, token)

	client := saxo_openapi.NewAPIClient(saxo_openapi.NewConfiguration())
	auth := context.WithValue(ctx, saxo_openapi.ContextOAuth2, s.tokenSource)
	auth = context.WithValue(auth, saxo_openapi.ContextMockedDataID, "001")

	s.ma = &ModeledAPI{
		Ctx:    auth,
		Client: client,
	}

	s.acc, err = s.ma.GetAccounts()
	c.Assert(err, check.IsNil)
}

func (s *ModeledApiSuite) TearDownSuite(c *check.C) {
	if s.tokenSource != nil {
		s.ma.Throttle()
		_, _, err := s.ma.Client.PortfolioApi.ResetAccount(s.ma.Ctx,
			s.acc.GetAccountKeyMe(), &saxo_openapi.PortfolioApiResetAccountOpts{
				Body: optional.NewInterface(map[string]string{"NewBalance": "10000"}),
			})
		c.Assert(err, check.IsNil)
		s.ma.UpdateLastCall()

		// Always write token back if everything went ok
		token, err := s.tokenSource.Token()
		c.Assert(err, check.IsNil)

		err = saxo_oauth2.PersistToken(token)
		c.Assert(err, check.IsNil)
	}
}

var _ = check.Suite(&ModeledApiSuite{})

func TestModeledAPI(t *testing.T) {
	check.TestingT(t)
}

func (s *ModeledApiSuite) TestGetInstrumentDetails(c *check.C) {

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	id, err := s.ma.GetInstrumentDetails(i)
	c.Assert(err, check.IsNil)

	c.Assert(id.GetSymbol(), check.Equals, i.GetSymbol())
}

func (s *ModeledApiSuite) TestGetInstrumentPrice(c *check.C) {

	i, err := s.ma.GetInstrument("IBIO:xase")
	c.Assert(err, check.IsNil)

	p, err := s.ma.GetInstrumentPrice(i)
	c.Assert(err, check.IsNil)

	c.Assert(p, check.Equals, 0.0)
}

func (s *ModeledApiSuite) TestGetInstrumentDetails_IBIO(c *check.C) {

	i, err := s.ma.GetInstrument("IBIO:xase")
	c.Assert(err, check.IsNil)

	id, err := s.ma.GetInstrumentDetails(i)
	c.Assert(err, check.IsNil)

	c.Assert(id.GetSymbol(), check.Equals, i.GetSymbol())

	priceTick := id.CalculatePriceWithThickSize(2.12345, 0)
	c.Assert(priceTick, check.Equals, 2.12)
}
