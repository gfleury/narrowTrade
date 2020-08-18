package saxo_models

import (
	"github.com/antihax/optional"
	"github.com/gfleury/intensiveTrade/saxo_openapi"
	"github.com/mitchellh/mapstructure"
)

type Order struct {
	AccountKey    string    `json:"AccountKey"`
	Amount        int       `json:"Amount"`
	BuySell       BuySell   `json:"BuySell"`
	OrderType     OrderType `json:"OrderType"`
	ManualOrder   bool      `json:"ManualOrder"`
	Uic           int32     `json:"Uic"`
	AssetType     AssetType `json:"AssetType"`
	OrderDuration struct {
		DurationType DurationType `json:"DurationType"`
	} `json:"OrderDuration"`
	Orders         []Order `json:"Orders"`
	OrderPrice     float64 `json:"OrderPrice,omitempty"`
	GoodToGo       bool    `json:"-"`
	EstimatedPrice float64 `json:"-"`
}

type ErrorInfo struct {
	ErrorCode string `json:"ErrorCode"`
	Message   string `json:"Message"`
}

type PreOrderResponse struct {
	EstimatedCashRequired             float64   `json:"EstimatedCashRequired"`
	EstimatedCashRequiredCurrency     string    `json:"EstimatedCashRequiredCurrency"`
	InstrumentToAccountConversionRate float64   `json:"InstrumentToAccountConversionRate"`
	PreCheckResult                    string    `json:"PreCheckResult"`
	ErrorInfo                         ErrorInfo `json:"ErrorInfo"`
}

type OrderResponse struct {
	OrderID   string          `json:"OrderId"`
	Orders    []OrderResponse `json:"Orders"`
	ErrorInfo *struct {
		ErrorCode string `json:"ErrorCode,omitempty"`
		Message   string `json:"Message,omitempty"`
	} `json:"ErrorInfo,omitempty"`
}

func (ma *ModeledAPI) PreOrder(o *Order) (*PreOrderResponse, error) {
	b := &PreOrderResponse{}
	data, _, err := ma.Client.TradingApi.PreCheckOrder(ma.Ctx, &saxo_openapi.TradingApiPreCheckOrderOpts{Body: optional.NewInterface(&o)})
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, b)
	if b.PreCheckResult == "Ok" {
		o.GoodToGo = true
		o.EstimatedPrice = b.EstimatedCashRequired
	}
	return b, err
}

func (ma *ModeledAPI) Order(o *Order) (*OrderResponse, error) {
	b := &OrderResponse{}
	data, _, err := ma.Client.TradingApi.PlaceOrder(ma.Ctx, &saxo_openapi.TradingApiPlaceOrderOpts{Body: optional.NewInterface(&o)})
	if err != nil {
		mapErr := GetMapError(err)
		if mapErr == nil {
			return nil, err
		}
		errMarsh := mapstructure.Decode(mapErr, b)
		if errMarsh != nil {
			return nil, err
		}
		return b, err
	}

	err = mapstructure.Decode(data, b)

	return b, err
}
