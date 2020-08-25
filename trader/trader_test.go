package trader

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/saxo_oauth2"
	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/gfleury/narrowTrade/tests"

	"github.com/antihax/optional"
	"golang.org/x/oauth2"
	check "gopkg.in/check.v1"
)

type Suite struct {
	acc         *models.Accounts
	ma          *models.ModeledAPI
	tokenSource oauth2.TokenSource
	ex          *models.Exchange
	ip          *models.InfoPrice
	id          models.InstrumentDetails
	i           models.Instrument
}

func (s *Suite) SetUpSuite(c *check.C) {
	ctx := context.Background()

	oauth2cfg := tests.GetTestOauth()

	token, err := saxo_oauth2.GetToken(ctx, oauth2cfg)
	c.Assert(err, check.IsNil)

	s.tokenSource = oauth2cfg.TokenSource(ctx, token)

	client := saxo_openapi.NewAPIClient(saxo_openapi.NewConfiguration())
	auth := context.WithValue(ctx, saxo_openapi.ContextOAuth2, s.tokenSource)
	auth = context.WithValue(auth, saxo_openapi.ContextMockedDataID, "001")

	s.ma = &models.ModeledAPI{
		Ctx:    auth,
		Client: client,
	}

	s.acc, err = s.ma.GetAccounts()
	c.Assert(err, check.IsNil)

	s.i, err = s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	s.ex, err = s.ma.GetExchange(s.i.GetExchangeID())
	c.Assert(err, check.IsNil)

	s.ip, err = s.ma.GetInfoPrice(s.i)
	c.Assert(err, check.IsNil)

	s.id, err = s.ma.GetInstrumentDetails(s.i)
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
	time.Sleep(2 * time.Second)
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
		ModeledAPI: s.ma,
	}

	i, err := s.ma.GetInstrument("AAPL:xnas")
	c.Assert(err, check.IsNil)

	or, err := t.Buy(
		i.GetOrder().
			WithAmount(10).
			WithType(models.Market))

	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
	}
}

func (s *Suite) TestTradeSimple_Buy_Limit_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Limit).
			WithPrice(s.ip.Quote.Ask))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Limit)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
		c.Assert(ol.Data[0].Price, check.Equals, s.ip.Quote.Ask)
	}
}

func (s *Suite) TestTradeSimple_Buy_Market_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Market).
			WithStopLoss(s.ip.Quote.Bid - 5))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrder := models.RelatedOpenOrders{
			Amount: 10,
			Duration: models.Duration{
				DurationType: "GoodTillCancel",
			},
			OpenOrderType: "StopIfTraded",
			OrderPrice:    s.ip.Quote.Bid - 5,
			OrderID:       or.Orders[0].OrderID,
			Status:        "NotWorking",
		}

		c.Assert(ol.Data[0].RelatedOpenOrders[0], check.DeepEquals, relatedOrder)
	}
}

func (s *Suite) TestTradeSimple_Buy_Market_TakeProfit_StopLoss_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Market).
			WithStopLoss(s.ip.Quote.Bid - 3).
			WithTakeProfit(s.ip.Quote.Ask + 3))
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []models.RelatedOpenOrders{{
			Amount: 10,
			Duration: models.Duration{
				DurationType: models.GoodTillCancel,
			},
			OpenOrderType: models.StopIfTraded,
			OrderPrice:    s.ip.Quote.Bid - 3,
			OrderID:       or.Orders[1].OrderID,
			Status:        "NotWorking",
		},
			{
				Amount: 10,
				Duration: models.Duration{
					DurationType: models.GoodTillCancel,
				},
				OpenOrderType: models.Limit,
				OrderPrice:    s.ip.Quote.Ask + 3,
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
		ModeledAPI: s.ma,
	}

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Limit).
			WithDuration(models.GoodTillCancel).
			WithPrice(s.ip.Quote.Ask).
			WithStopLoss(s.ip.Quote.Bid - 3).
			WithTakeProfit(s.ip.Quote.Ask + 3),
	)
	if err != nil {
		fmt.Println(models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Limit)
		c.Assert(ol.Data[0].Price, check.Equals, s.ip.Quote.Ask)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []models.RelatedOpenOrders{{
			Amount: 10,
			Duration: models.Duration{
				DurationType: models.GoodTillCancel,
			},
			OpenOrderType: models.StopIfTraded,
			OrderPrice:    s.ip.Quote.Bid - 3,
			OrderID:       or.Orders[1].OrderID,
			Status:        "NotWorking",
		},
			{
				Amount: 10,
				Duration: models.Duration{
					DurationType: models.GoodTillCancel,
				},
				OpenOrderType: models.Limit,
				OrderPrice:    s.ip.Quote.Ask + 3,
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
		ModeledAPI: s.ma,
	}

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.StopIfTraded).
			WithPrice(s.ip.Quote.Ask + 1))
	if err != nil {
		fmt.Println(models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	ex, err := s.ma.GetExchange(s.i.GetExchangeID())
	c.Assert(err, check.IsNil)

	if !ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.StopIfTraded)
		c.Assert(ol.Data[0].Price, check.Equals, s.ip.Quote.Ask+1)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
	}
}

func (s *Suite) TestTradeSimple_Buy_StopLimit_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.StopLimit).
			WithPrice(s.ip.Quote.Ask).
			WithStopLimitPrice(s.ip.Quote.Ask + 2).
			WithStopLoss(s.ip.Quote.Ask - 1))
	if err != nil {
		fmt.Println(models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.StopLimit)
		c.Assert(ol.Data[0].Price, check.Equals, s.ip.Quote.Ask)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)
	}
}

func (s *Suite) TestTradeSimple_Buy_Market_TrailingStop_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	stopLossPrice := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask, 2)
	distanceToMarket := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask-stopLossPrice, 0)

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Market).
			WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05))
	if err != nil {
		fmt.Println(models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []models.RelatedOpenOrders{{
			Amount: 10,
			Duration: models.Duration{
				DurationType: models.GoodTillCancel,
			},
			OpenOrderType:                models.TrailingStopIfTraded,
			OrderPrice:                   stopLossPrice,
			TrailingStopDistanceToMarket: distanceToMarket,
			TrailingStopStep:             0.05,
			OrderID:                      or.Orders[0].OrderID,
			Status:                       "NotWorking",
		},
		}

		c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
	}
}

func (s *Suite) TestTradeSimple_Buy_Limit_TrailingStop_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	stopLossPrice := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask, 2)
	distanceToMarket := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask-stopLossPrice, 0)

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Limit).
			WithPrice(s.ip.Quote.Ask).
			WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05))
	if err != nil {
		fmt.Println(models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Limit)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []models.RelatedOpenOrders{{
			Amount: 10,
			Duration: models.Duration{
				DurationType: models.GoodTillCancel,
			},
			OpenOrderType:                models.TrailingStopIfTraded,
			OrderPrice:                   stopLossPrice,
			TrailingStopDistanceToMarket: distanceToMarket,
			TrailingStopStep:             0.05,
			OrderID:                      or.Orders[0].OrderID,
			Status:                       "NotWorking",
		},
		}

		c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
	}
}

func (s *Suite) TestTradeSimple_Buy_Limit_TakeProfit_TrailingStop_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	stopLossPrice := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask, 2)
	distanceToMarket := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask-stopLossPrice, 0)

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Limit).
			WithPrice(s.ip.Quote.Ask).
			WithTakeProfit(s.ip.Quote.Ask+3).
			WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05))
	if err != nil {
		fmt.Println(models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Limit)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []models.RelatedOpenOrders{{
			Amount: 10,
			Duration: models.Duration{
				DurationType: models.GoodTillCancel,
			},
			OpenOrderType:                models.TrailingStopIfTraded,
			OrderPrice:                   stopLossPrice,
			TrailingStopDistanceToMarket: distanceToMarket,
			TrailingStopStep:             0.05,
			OrderID:                      or.Orders[1].OrderID,
			Status:                       "NotWorking",
		},
			{
				Amount: 10,
				Duration: models.Duration{
					DurationType: models.GoodTillCancel,
				},
				OpenOrderType: models.Limit,
				OrderPrice:    s.ip.Quote.Ask + 3,
				OrderID:       or.Orders[0].OrderID,
				Status:        "NotWorking",
			},
		}

		c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
	}
}

func (s *Suite) TestTradeSimple_Buy_Market_TakeProfit_TrailingStop_10_APPL(c *check.C) {
	t := BasicSaxoTrader{
		AccountKey: s.acc.GetAccountKeyMe(),
		ModeledAPI: s.ma,
	}

	stopLossPrice := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask, 2)
	distanceToMarket := s.id.CalculatePriceWithThickSize(s.ip.Quote.Ask-stopLossPrice, 0)

	or, err := t.Buy(
		s.i.GetOrder().
			WithAmount(10).
			WithType(models.Market).
			WithTakeProfit(s.ip.Quote.Ask+3).
			WithStopLossTrailingStop(stopLossPrice, distanceToMarket, 0.05))
	if err != nil {
		fmt.Println(models.GetStringError(err))
	}
	c.Assert(err, check.IsNil)

	c.Assert(or.ErrorInfo, check.IsNil)

	ol, err := s.ma.GetOrdersMe()
	c.Assert(err, check.IsNil)

	if !s.ex.IsOpen {
		c.Assert(ol.Data[0].BuySell, check.Equals, models.Buy)
		c.Assert(ol.Data[0].OpenOrderType, check.Equals, models.Market)
		c.Assert(ol.Data[0].Amount, check.Equals, 10)

		relatedOrders := []models.RelatedOpenOrders{{
			Amount: 10,
			Duration: models.Duration{
				DurationType: models.GoodTillCancel,
			},
			OpenOrderType:                models.TrailingStopIfTraded,
			OrderPrice:                   stopLossPrice,
			TrailingStopDistanceToMarket: distanceToMarket,
			TrailingStopStep:             0.05,
			OrderID:                      or.Orders[1].OrderID,
			Status:                       "NotWorking",
		},
			{
				Amount: 10,
				Duration: models.Duration{
					DurationType: models.GoodTillCancel,
				},
				OpenOrderType: models.Limit,
				OrderPrice:    s.ip.Quote.Ask + 3,
				OrderID:       or.Orders[0].OrderID,
				Status:        "NotWorking",
			},
		}

		c.Assert(ol.Data[0].RelatedOpenOrders, check.DeepEquals, relatedOrders)
	}
}
