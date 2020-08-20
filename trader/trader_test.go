package trader

import (
	"context"

	"github.com/antihax/optional"
	"github.com/gfleury/intensiveTrade/saxo_models"
	"github.com/gfleury/intensiveTrade/saxo_oauth2"
	"github.com/gfleury/intensiveTrade/saxo_openapi"
	"github.com/gfleury/intensiveTrade/tests"
	"golang.org/x/oauth2"
	check "gopkg.in/check.v1"

	"testing"
)

type Suite struct {
	acc         *saxo_models.Accounts
	ma          *saxo_models.ModeledAPI
	tokenSource oauth2.TokenSource
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
}

func (s *Suite) TearDownSuite(c *check.C) {
	s.SetUpTest(c)

	// Always write token back if everything went ok
	token, err := s.tokenSource.Token()
	c.Assert(err, check.IsNil)

	err = saxo_oauth2.PersistToken(token)
	c.Assert(err, check.IsNil)
}

func (s *Suite) SetUpTest(c *check.C) {
	_, _, err := s.ma.Client.PortfolioApi.ResetAccount(s.ma.Ctx,
		s.acc.GetAccountKeyMe(), &saxo_openapi.PortfolioApiResetAccountOpts{
			Body: optional.NewInterface(map[string]string{"NewBalance": "10000"}),
		})
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

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.Market, 10)
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
	c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Market)
	c.Assert(ol.Data[0].Amount, check.Equals, 10)
}

func (s *Suite) TestTradeSimple_Buy_Limit_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.Limit, 10, NewOrderOption(OrderPrice, ip.Quote.Mid))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
	c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Limit)
	c.Assert(ol.Data[0].Amount, check.Equals, 10)
	c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Mid)
}

func (s *Suite) TestTradeSimple_Buy_Market_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.Market, 10, NewOrderOption(StopLoss, ip.Quote.Bid))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
	c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Market)
	c.Assert(ol.Data[0].Amount, check.Equals, 10)

	relatedOrder := saxo_models.RelatedOpenOrders{
		Amount: 10,
		Duration: saxo_models.Duration{
			DurationType: "GoodTillCancel",
		},
		OpenOrderType: "StopIfTraded",
		OrderPrice:    ip.Quote.Bid,
		OrderID:       or.Orders[0].OrderID,
		Status:        "NotWorking",
	}

	c.Assert(ol.Data[0].RelatedOpenOrders[0], check.DeepEquals, relatedOrder)
}

func (s *Suite) TestTradeSimple_Buy_Market_TakeProfit_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.Market, 10,
		NewOrderOption(StopLoss, ip.Quote.Bid-3),
		NewOrderOption(TakeProfit, ip.Quote.Ask+3))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

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

func (s *Suite) TestTradeSimple_Buy_Limit_TakeProfit_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.Limit, 10,
		NewOrderOption(StopLoss, ip.Quote.Bid-3),
		NewOrderOption(TakeProfit, ip.Quote.Ask+3),
		NewOrderOption(OrderPrice, ip.Quote.Mid),
	)
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
	c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Limit)
	c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Mid)
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

func (s *Suite) TestTradeSimple_Buy_StopIfTraded_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.StopIfTraded, 10,
		NewOrderOption(OrderPrice, ip.Quote.Mid),
	)
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
	c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.StopIfTraded)
	c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Mid)
	c.Assert(ol.Data[0].Amount, check.Equals, 10)
}

func (s *Suite) TestTradeSimple_Buy_StopLimit_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.StopLimit, 10,
		NewOrderOption(OrderPrice, ip.Quote.Mid),
		NewOrderOption(StopLimitPrice, ip.Quote.Mid+3),
	)
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
	c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.StopLimit)
	c.Assert(ol.Data[0].Price, check.Equals, ip.Quote.Mid)
	c.Assert(ol.Data[0].Amount, check.Equals, 10)
}

func (s *Suite) TestTradeSimple_Buy_Market_TrailingStop_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		API:        s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL")
	c.Assert(err, check.IsNil)

	ip, err := s.ma.GetInfoPrice(i)
	c.Assert(err, check.IsNil)

	or, err := t.Buy(i, saxo_models.Market, 10,
		NewOrderOption(TrailingStop, TrailingStopValues{
			OrderPrice:                   ip.Quote.Mid - 2.37,
			TrailingStopDistanceToMarket: 2.37,
			TrailingStopStep:             0.5,
		}),
	)
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	c.Assert(ol.Data[0].BuySell, check.Equals, saxo_models.Buy)
	c.Assert(ol.Data[0].OpenOrderType, check.Equals, saxo_models.Market)
	c.Assert(ol.Data[0].Amount, check.Equals, 10)

	relatedOrders := []saxo_models.RelatedOpenOrders{{
		Amount: 10,
		Duration: saxo_models.Duration{
			DurationType: saxo_models.GoodTillCancel,
		},
		OpenOrderType:                saxo_models.TrailingStopIfTraded,
		OrderPrice:                   ip.Quote.Mid - 2.37,
		TrailingStopDistanceToMarket: 2.37,
		TrailingStopStep:             0.5,
		OrderID:                      or.Orders[0].OrderID,
		Status:                       "NotWorking",
	},
	}

	c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
}
