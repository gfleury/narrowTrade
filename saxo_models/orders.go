package saxo_models

import (
	"github.com/antihax/optional"
	"github.com/gfleury/intensiveTrade/saxo_openapi"
	"github.com/mitchellh/mapstructure"
)

type Duration struct {
	DurationType DurationType `json:"DurationType"`
}

// Order is used on POST Orders Endpoints (EX. Placing an Order)
type Order struct {
	AccountKey                   string    `json:"AccountKey"`
	Amount                       int       `json:"Amount"`
	BuySell                      BuySell   `json:"BuySell"`
	OrderType                    OrderType `json:"OrderType"`
	ManualOrder                  bool      `json:"ManualOrder"`
	Uic                          int32     `json:"Uic"`
	AssetType                    AssetType `json:"AssetType"`
	OrderDuration                Duration  `json:"OrderDuration"`
	Orders                       []Order   `json:"Orders"`
	OrderPrice                   float64   `json:"OrderPrice,omitempty"`
	StopLimitPrice               float64   `json:"StopLimitPrice,omitempty"`
	TrailingstopDistanceToMarket float64   `json:"TrailingstopDistanceToMarket,omitempty"`
	TrailingStopStep             float64   `json:"TrailingStopStep,omitempty"`

	// Internal fields
	GoodToGo       bool    `json:"-"`
	EstimatedPrice float64 `json:"-"`
}

type OrderList struct {
	Count int           `json:"__count"`
	Data  []ActiveOrder `json:"Data"`
}

type Exchange struct {
	Description string `json:"Description"`
	ExchangeID  string `json:"ExchangeId"`
	IsOpen      bool   `json:"IsOpen"`
}

type RelatedOpenOrders struct {
	Amount        int      `json:"Amount"`
	Duration      Duration `json:"Duration"`
	OpenOrderType string   `json:"OpenOrderType"`
	OrderID       string   `json:"OrderId"`
	OrderPrice    float64  `json:"OrderPrice"`
	Status        string   `json:"Status"`
}

// ActiveOrder Is used on the GET Orders Endpoints (EX. List/Getting an Order)
type ActiveOrder struct {
	Amount                   int                 `json:"Amount"`
	BuySell                  BuySell             `json:"BuySell"`
	OrderType                OrderType           `json:"OrderType"`
	AccountID                string              `json:"AccountId,omitempty"`
	CalculationReliability   string              `json:"CalculationReliability,omitempty"`
	ClientID                 string              `json:"ClientId,omitempty"`
	ClientKey                string              `json:"ClientKey,omitempty"`
	ClientName               string              `json:"ClientName,omitempty"`
	CorrelationKey           string              `json:"CorrelationKey,omitempty"`
	CurrentPrice             int                 `json:"CurrentPrice,omitempty"`
	CurrentPriceDelayMinutes int                 `json:"CurrentPriceDelayMinutes,omitempty"`
	CurrentPriceType         string              `json:"CurrentPriceType,omitempty"`
	DistanceToMarket         float64             `json:"DistanceToMarket,omitempty"`
	Duration                 *Duration           `json:"Duration,omitempty"`
	Exchange                 *Exchange           `json:"Exchange,omitempty"`
	IsForceOpen              bool                `json:"IsForceOpen,omitempty"`
	Isin                     string              `json:"Isin,omitempty"`
	IsMarketOpen             bool                `json:"IsMarketOpen,omitempty"`
	MarketPrice              int                 `json:"MarketPrice,omitempty"`
	MarketValue              int                 `json:"MarketValue,omitempty"`
	NonTradableReason        string              `json:"NonTradableReason,omitempty"`
	OpenOrderType            string              `json:"OpenOrderType,omitempty"`
	OrderAmountType          string              `json:"OrderAmountType,omitempty"`
	OrderID                  string              `json:"OrderId,omitempty"`
	OrderRelation            string              `json:"OrderRelation,omitempty"`
	OrderTime                string              `json:"OrderTime,omitempty"`
	Price                    float64             `json:"Price,omitempty"`
	RelatedOpenOrders        []RelatedOpenOrders `json:"RelatedOpenOrders,omitempty"`
	RelatedPositionID        string              `json:"RelatedPositionId,omitempty"`
	Status                   string              `json:"Status,omitempty"`
	TradingStatus            string              `json:"TradingStatus,omitempty"`
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
	ma.Throttle()
	data, _, err := ma.Client.TradingApi.PreCheckOrder(ma.Ctx, &saxo_openapi.TradingApiPreCheckOrderOpts{Body: optional.NewInterface(&o)})
	defer ma.UpdateLastCall()
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
	ma.Throttle()
	data, _, err := ma.Client.TradingApi.PlaceOrder(ma.Ctx, &saxo_openapi.TradingApiPlaceOrderOpts{Body: optional.NewInterface(&o)})
	defer ma.UpdateLastCall()
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

func (ma *ModeledAPI) GetOrdersMe() (*OrderList, error) {
	b := &OrderList{}

	ma.Throttle()
	data, _, err := ma.Client.PortfolioApi.GetOpenOrders(ma.Ctx, &saxo_openapi.PortfolioApiGetOpenOrdersOpts{})
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, b)

	return b, err
}
