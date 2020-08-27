package models

import (
	"github.com/antihax/optional"
	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/mitchellh/mapstructure"
)

type WatchlistRequest struct {
	Arguments   Arguments `json:"Arguments"`
	RefreshRate int       `json:"RefreshRate"`
	Format      string    `json:"Format"`
	ContextID   string    `json:"ContextId"`
	ReferenceID string    `json:"ReferenceId"`
}

type Arguments struct {
	WatchlistID string `json:"WatchlistId"`
	AssetTypes  string `json:"AssetTypes"`
	Index       int    `json:"Index"`
	RowCount    int    `json:"RowCount"`
}

type Watchlist struct {
	ContextID         string   `json:"ContextId"`
	Format            string   `json:"Format"`
	InactivityTimeout int      `json:"InactivityTimeout"`
	ReferenceID       string   `json:"ReferenceId"`
	RefreshRate       int      `json:"RefreshRate"`
	Snapshot          Snapshot `json:"Snapshot"`
	State             string   `json:"State"`
}

type DisplayAndFormat struct {
	Currency    string `json:"Currency"`
	Decimals    int    `json:"Decimals"`
	Description string `json:"Description"`
	Format      string `json:"Format"`
	Name        string `json:"Name"`
	Symbol      string `json:"Symbol"`
}

type Rows struct {
	Amount                            float64          `json:"Amount,omitempty"`
	AnalystTargetPrice                float64          `json:"AnalystTargetPrice,omitempty"`
	Ask                               float64          `json:"Ask,omitempty"`
	AskSize                           float64          `json:"AskSize,omitempty"`
	AssetType                         string           `json:"AssetType"`
	AverageExpectedSalesGrowthRatePct float64          `json:"AverageExpectedSalesGrowthRatePct,omitempty"`
	Bid                               float64          `json:"Bid,omitempty"`
	BidSize                           float64          `json:"BidSize,omitempty"`
	Country                           string           `json:"Country"`
	DelayedByMinutes                  int              `json:"DelayedByMinutes,omitempty"`
	DisplayAndFormat                  DisplayAndFormat `json:"DisplayAndFormat"`
	Eps                               float64          `json:"Eps,omitempty"`
	ExchangeID                        string           `json:"ExchangeId,omitempty"`
	FiveYearsReturnPct                float64          `json:"FiveYearsReturnPct,omitempty"`
	High                              float64          `json:"High,omitempty"`
	Index                             int              `json:"Index"`
	Isin                              string           `json:"Isin,omitempty"`
	IsMarketOpen                      bool             `json:"IsMarketOpen,omitempty"`
	IsSrdEligible                     bool             `json:"IsSrdEligible"`
	IssuerCountry                     string           `json:"IssuerCountry,omitempty"`
	IssuerName                        string           `json:"IssuerName,omitempty"`
	IsTradable                        bool             `json:"IsTradable,omitempty"`
	LastClose                         float64          `json:"LastClose,omitempty"`
	LastTraded                        float64          `json:"LastTraded,omitempty"`
	LastTradedSize                    float64          `json:"LastTradedSize,omitempty"`
	LastUpdated                       string           `json:"LastUpdated,omitempty"`
	Low                               float64          `json:"Low,omitempty"`
	MarketCap                         float64          `json:"MarketCap,omitempty"`
	MarketCapLocalCurrency            float64          `json:"MarketCapLocalCurrency,omitempty"`
	MarketState                       string           `json:"MarketState,omitempty"`
	Mid                               float64          `json:"Mid,omitempty"`
	MinimumTradeSize                  float64          `json:"MinimumTradeSize,omitempty"`
	NetChange                         float64          `json:"NetChange,omitempty"`
	OneDayReturnPct                   float64          `json:"OneDayReturnPct,omitempty"`
	OneMonthReturnPct                 float64          `json:"OneMonthReturnPct,omitempty"`
	OneWeekReturnPct                  float64          `json:"OneWeekReturnPct,omitempty"`
	OneYearReturnPct                  float64          `json:"OneYearReturnPct,omitempty"`
	Open                              float64          `json:"Open,omitempty"`
	PercentChange                     float64          `json:"PercentChange,omitempty"`
	Price                             float64          `json:"Price,omitempty"`
	PriceEarnings                     float64          `json:"PriceEarnings,omitempty"`
	PriceSource                       string           `json:"PriceSource,omitempty"`
	PriceSourceType                   string           `json:"PriceSourceType,omitempty"`
	PriceTargetUpside                 float64          `json:"PriceTargetUpside,omitempty"`
	PriceTypeAsk                      string           `json:"PriceTypeAsk,omitempty"`
	PriceTypeBid                      string           `json:"PriceTypeBid,omitempty"`
	SixMonthsReturnPct                float64          `json:"SixMonthsReturnPct,omitempty"`
	Spread                            float64          `json:"Spread,omitempty"`
	SupportedOrderTypes               []string         `json:"SupportedOrderTypes,omitempty"`
	ThreeMonthsReturnPct              float64          `json:"ThreeMonthsReturnPct,omitempty"`
	ThreeYearsReturnPct               float64          `json:"ThreeYearsReturnPct,omitempty"`
	TradableOn                        []string         `json:"TradableOn,omitempty"`
	TwoMonthsReturnPct                float64          `json:"TwoMonthsReturnPct,omitempty"`
	TwoYearsReturnPct                 float64          `json:"TwoYearsReturnPct,omitempty"`
	Uic                               int              `json:"Uic"`
	Volume                            float64          `json:"Volume,omitempty"`
	WatchlistInstrumentID             string           `json:"WatchlistInstrumentId"`
	YtdReturnPct                      float64          `json:"YtdReturnPct,omitempty"`
	DividendYieldPct                  float64          `json:"DividendYieldPct,omitempty"`
}

type Snapshot struct {
	RowCount int    `json:"RowCount"`
	Rows     []Rows `json:"Rows"`
}

func (ma *ModeledAPI) GetWatchlist(wr *WatchlistRequest) (*Watchlist, error) {
	w := &Watchlist{}
	ma.Throttle()
	data, _, err := ma.Client.TradingApi.GetWatchlist(ma.Ctx,
		&saxo_openapi.TradingApiGetWatchlistOpts{Body: optional.NewInterface(wr)})
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, w)
	return w, err
}
