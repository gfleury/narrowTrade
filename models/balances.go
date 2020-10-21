package models

import (
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Balance struct {
	CalculationReliability string  `json:"CalculationReliability"`
	CashBalance            float64 `json:"CashBalance"`
	ChangesScheduled       bool    `json:"ChangesScheduled"`
	ClosedPositionsCount   float64 `json:"ClosedPositionsCount"`
	CollateralAvailable    float64 `json:"CollateralAvailable"`
	CollateralCreditValue  struct {
		Line           float64 `json:"Line"`
		UtilizationPct float64 `json:"UtilizationPct"`
	} `json:"CollateralCreditValue"`
	CostToClosePositions float64 `json:"CostToClosePositions"`
	Currency             string  `json:"Currency"`
	CurrencyDecimals     float64 `json:"CurrencyDecimals"`
	InitialMargin        struct {
		CollateralAvailable   float64 `json:"CollateralAvailable"`
		CollateralCreditValue struct {
			Line           float64 `json:"Line"`
			UtilizationPct float64 `json:"UtilizationPct"`
		} `json:"CollateralCreditValue"`
		MarginAvailable              float64 `json:"MarginAvailable"`
		MarginCollateralNotAvailable float64 `json:"MarginCollateralNotAvailable"`
		MarginUsedByCurrentPositions float64 `json:"MarginUsedByCurrentPositions"`
		MarginUtilizationPct         float64 `json:"MarginUtilizationPct"`
		NetEquityForMargin           float64 `json:"NetEquityForMargin"`
	} `json:"InitialMargin"`
	IsPortfolioMarginModelSimple     bool    `json:"IsPortfolioMarginModelSimple"`
	MarginAvailableForTrading        float64 `json:"MarginAvailableForTrading"`
	MarginCollateralNotAvailable     float64 `json:"MarginCollateralNotAvailable"`
	MarginExposureCoveragePct        float64 `json:"MarginExposureCoveragePct"`
	MarginNetExposure                float64 `json:"MarginNetExposure"`
	MarginUsedByCurrentPositions     float64 `json:"MarginUsedByCurrentPositions"`
	MarginUtilizationPct             float64 `json:"MarginUtilizationPct"`
	NetEquityForMargin               float64 `json:"NetEquityForMargin"`
	NetPositionsCount                float64 `json:"NetPositionsCount"`
	NonMarginPositionsValue          float64 `json:"NonMarginPositionsValue"`
	OpenPositionsCount               float64 `json:"OpenPositionsCount"`
	OptionPremiumsMarketValue        float64 `json:"OptionPremiumsMarketValue"`
	OrdersCount                      float64 `json:"OrdersCount"`
	OtherCollateral                  float64 `json:"OtherCollateral"`
	SettlementValue                  float64 `json:"SettlementValue"`
	TotalValue                       float64 `json:"TotalValue"`
	TransactionsNotBooked            float64 `json:"TransactionsNotBooked"`
	UnrealizedMarginClosedProfitLoss float64 `json:"UnrealizedMarginClosedProfitLoss"`
	UnrealizedMarginOpenProfitLoss   float64 `json:"UnrealizedMarginOpenProfitLoss"`
	UnrealizedMarginProfitLoss       float64 `json:"UnrealizedMarginProfitLoss"`
	UnrealizedPositionsValue         float64 `json:"UnrealizedPositionsValue"`
}

func (ma *ModeledAPI) GetBalanceMe() (*Balance, error) {
	b := &Balance{}
	ma.Throttle()
	ctx, cancel := context.WithTimeout(ma.Ctx, 15*time.Second)
	defer cancel()

	data, _, err := ma.Client.PortfolioApi.GetBalance(ctx)
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, b)
	return b, err
}
