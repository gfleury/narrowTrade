package trader

import (
	"context"
	"fmt"
	"testing"

	"github.com/gfleury/intensiveTrade/saxo_models"
	"github.com/gfleury/intensiveTrade/saxo_oauth2"
	"github.com/gfleury/intensiveTrade/saxo_openapi"
	"github.com/gfleury/intensiveTrade/tests"

	"github.com/antihax/optional"
	"golang.org/x/oauth2"
	check "gopkg.in/check.v1"
)

type Suite struct {
	acc         *saxo_models.Accounts
	ma          *saxo_models.ModeledAPI
	tokenSource oauth2.TokenSource
	ex          *saxo_models.Exchange
}

func (s *Suite) SetUpSuite(c *check.C) {
	ctx := context.Background()

	oauth2cfg := tests.GetTestOauth()

	token, err := saxo_oauth2.GetToken(ctx, oauth2cfg)
	c.Assert(err, check.IsNil)

	s.tokenSource = oauth2cfg.TokenSource(ctx, token)

	client := saxo_openapi.NewAPIClient(saxo_openapi.NewConfiguration())
	auth := context.WithValue(oauth2.NoContext, saxo_openapi.ContextOAuth2, s.tokenSource)
	auth = context.WithValue(auth, saxo_openapi.ContextMockedDataID, "001")

	s.ma = &saxo_models.ModeledAPI{
		Ctx:    auth,
		Client: client,
	}

	s.acc, err = s.ma.GetAccounts()
	c.Assert(err, check.IsNil)

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	s.ex, err = s.ma.GetExchange(i.GetExchangeID())
	c.Assert(err, check.IsNil)

}

func (s *Suite) TearDownSuite(c *check.C) {
	if s.tokenSource != nil {
		s.SetUpTest(c)

		// Always write token back if everything went ok
		token, err := s.tokenSource.Token()
		c.Assert(err, check.IsNil)

		err = saxo_oauth2.PersistToken(token)
		c.Assert(err, check.IsNil)
	}
}

func (s *Suite) SetUpTest(c *check.C) {
	s.ma.Throttle()
	_, _, err := s.ma.Client.PortfolioApi.ResetAccount(s.ma.Ctx,
		s.acc.GetAccountKeyMe(), &saxo_openapi.PortfolioApiResetAccountOpts{
			Body: optional.NewInterface(map[string]string{"NewBalance": "10000"}),
		})
	s.ma.UpdateLastCall()
	c.Assert(err, check.IsNil)
}

var _ = check.Suite(&Suite{})

func Test(t *testing.T) {
	check.TestingT(t)
}

func (s *Suite) TestTradeSimple_Buy_Market_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(saxo_models.Market))

	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
	}
}

func (s *Suite) TestTradeSimple_Buy_Limit_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(saxo_models.Limit).
			WithPrice(ip.Quote.Ask))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Limit)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
		c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Ask)
	}
}

func (s *Suite) TestTradeSimple_Buy_Market_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(saxo_models.Market).
			WithStopLoss(ip.Quote.Bid - 5))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrder := saxo_models.RelatedOpenOrders{
			Amount: 10,
			Duration: saxo_models.Duration{
				DurationType: "GoodTillCancel",
			},
			OpenOrderType: "StopIfTraded",
			OrderPrice:    ip.Quote.Bid - 5,
			OrderID:       or.Orders[0].OrderID,
			Status:        "NotWorking",
		}

		c.Assert(ol.Data[0].RelatedOpenOrders[0], check.DeepEquals, relatedOrder)
	}
}

func (s *Suite) TestTradeSimple_Buy_Market_TakeProfit_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(saxo_models.Market).
			WithStopLoss(ip.Quote.Bid - 3).
			WithTakeProfit(ip.Quote.Ask + 3))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []saxo_models.RelatedOpenOrders{{
			Amount: 10,
			Duration: saxo_models.Duration{
				DurationType: saxo_models.GoodTillCancel,
			},
			OpenOrderType: saxo_models.StopIfTraded,
			OrderPrice:    ip.Quote.Bid - 3,
			OrderID:       or.Orders[1].OrderID,
			Status:        "NotWorking",
		},
			{
				Amount: 10,
				Duration: saxo_models.Duration{
					DurationType: saxo_models.GoodTillCancel,
				},
				OpenOrderType: saxo_models.Limit,
				OrderPrice:    ip.Quote.Ask + 3,
				OrderID:       or.Orders[0].OrderID,
				Status:        "NotWorking",
			},
		}

		c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
	}
}

func (s *Suite) TestTradeSimple_Buy_Limit_TakeProfit_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(saxo_models.Limit).
			WithDuration(saxo_models.GoodTillCancel).
			WithPrice(ip.Quote.Ask).
			WithStopLoss(ip.Quote.Bid - 3).
			WithTakeProfit(ip.Quote.Ask + 3),
	)
	if err != nil {
		fmt.Println(saxo_models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Limit)
		c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Ask)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []saxo_models.RelatedOpenOrders{{
			Amount: 10,
			Duration: saxo_models.Duration{
				DurationType: saxo_models.GoodTillCancel,
			},
			OpenOrderType: saxo_models.StopIfTraded,
			OrderPrice:    ip.Quote.Bid - 3,
			OrderID:       or.Orders[1].OrderID,
			Status:        "NotWorking",
		},
			{
				Amount: 10,
				Duration: saxo_models.Duration{
					DurationType: saxo_models.GoodTillCancel,
				},
				OpenOrderType: saxo_models.Limit,
				OrderPrice:    ip.Quote.Ask + 3,
				OrderID:       or.Orders[0].OrderID,
				Status:        "NotWorking",
			},
		}

		c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
	}
}

func (s *Suite) TestTradeSimple_Buy_StopIfTraded_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i.GetOrder().
		WithAmount(10).
		WithType(saxo_models.StopIfTraded).
		WithPrice(ip.Quote.Ask))
	if err != nil {
		fmt.Println(saxo_models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	ex, err := s.ma.GetExchange(i.GetExchangeID())
	c.Assert(err, check.IsNil)

	if !ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.StopIfTraded)
		c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Ask)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
	}
}

func (s *Suite) TestTradeSimple_Buy_StopLimit_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(saxo_models.StopLimit).
			WithPrice(ip.Quote.Ask).
			WithStopLimitPrice(ip.Quote.Ask + 3))
	if err != nil {
		fmt.Println(saxo_models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.StopLimit)
		c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Ask)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
	}
}

func (s *Suite) TestTradeSimple_Buy_Market_TrailingStop_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	id, err := s.ma.GetInstrumentDetails(i)
	c.Assert(err, check.IsNil)

	stopLossPrice := id.CalculatePriceWithThickSize(ip.Quote.Ask, 1)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(saxo_models.Market).
			// WithPrice(ip.Quote.Ask).
			WithStopLossTrailingStop(stopLossPrice, ip.Quote.Ask-stopLossPrice, 0.05))
	if err != nil {
		fmt.Println(saxo_models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []saxo_models.RelatedOpenOrders{{
			Amount: 10,
			Duration: saxo_models.Duration{
				DurationType: saxo_models.GoodTillCancel,
			},
			OpenOrderType:                saxo_models.TrailingStopIfTraded,
			OrderPrice:                   ip.Quote.Ask - 2.37,
			TrailingStopDistanceToMarket: 2.37,
			TrailingStopStep:             0.5,
			OrderID:                      or.Orders[0].OrderID,
			Status:                       "NotWorking",
		},
		}

		c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
	}
}
