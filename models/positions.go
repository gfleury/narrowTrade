package models

import (
	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/mitchellh/mapstructure"
)

type PositionList struct {
	Count int            `json:"__count"`
	Data  []PositionData `json:"Data"`
}

type PositionBase struct {
	AccountID                  string              `json:"AccountId"`
	AccountKey                 string              `json:"AccountKey"`
	Amount                     int                 `json:"Amount"`
	AssetType                  string              `json:"AssetType"`
	CanBeClosed                bool                `json:"CanBeClosed"`
	ClientID                   string              `json:"ClientId"`
	CloseConversionRateSettled bool                `json:"CloseConversionRateSettled"`
	CorrelationKey             string              `json:"CorrelationKey"`
	ExecutionTimeOpen          string              `json:"ExecutionTimeOpen"`
	IsForceOpen                bool                `json:"IsForceOpen"`
	IsMarketOpen               bool                `json:"IsMarketOpen"`
	OpenPrice                  float64             `json:"OpenPrice"`
	RelatedOpenOrders          []RelatedOpenOrders `json:"RelatedOpenOrders"`
	SourceOrderID              string              `json:"SourceOrderId"`
	Status                     string              `json:"Status"`
	Uic                        int                 `json:"Uic"`
	ValueDate                  string              `json:"ValueDate"`
}

type PositionView struct {
	CalculationReliability          string  `json:"CalculationReliability"`
	CurrentPrice                    int     `json:"CurrentPrice"`
	CurrentPriceDelayMinutes        int     `json:"CurrentPriceDelayMinutes"`
	CurrentPriceType                string  `json:"CurrentPriceType"`
	Exposure                        int     `json:"Exposure"`
	ExposureCurrency                string  `json:"ExposureCurrency"`
	ExposureInBaseCurrency          int     `json:"ExposureInBaseCurrency"`
	InstrumentPriceDayPercentChange int     `json:"InstrumentPriceDayPercentChange"`
	MarketValue                     int     `json:"MarketValue"`
	ProfitLossOnTrade               float64 `json:"ProfitLossOnTrade"`
	ProfitLossOnTradeInBaseCurrency float64 `json:"ProfitLossOnTradeInBaseCurrency"`
	TradeCostsTotal                 int     `json:"TradeCostsTotal"`
	TradeCostsTotalInBaseCurrency   int     `json:"TradeCostsTotalInBaseCurrency"`
}

type PositionData struct {
	NetPositionID string       `json:"NetPositionId"`
	PositionBase  PositionBase `json:"PositionBase"`
	PositionID    string       `json:"PositionId"`
	PositionView  PositionView `json:"PositionView"`
}

func (ma *ModeledAPI) GetPositionsMe() (*PositionList, error) {
	b := &PositionList{}

	ma.Throttle()
	data, _, err := ma.Client.PortfolioApi.GetPositions(ma.Ctx, &saxo_openapi.PortfolioApiGetPositionsOpts{})
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, b)

	return b, err
}
