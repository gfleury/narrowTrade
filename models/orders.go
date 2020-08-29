package models

import (
	"github.com/antihax/optional"
	"github.com/gfleury/narrowTrade/saxo_openapi"
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
	TrailingStopDistanceToMarket float64   `json:"TrailingStopDistanceToMarket,omitempty"`
	TrailingStopStep             float64   `json:"TrailingStopStep,omitempty"`

	// Internal fields
	GoodToGo       bool    `json:"-"`
	EstimatedPrice float64 `json:"-"`
}

type OrderList struct {
	Count int           `json:"__count"`
	Data  []ActiveOrder `json:"Data"`
}

type ExchangeOrder struct {
	Description string `json:"Description"`
	ExchangeID  string `json:"ExchangeId"`
	IsOpen      bool   `json:"IsOpen"`
}

type RelatedOpenOrders struct {
	Amount                       int      `json:"Amount"`
	Duration                     Duration `json:"Duration"`
	OpenOrderType                string   `json:"OpenOrderType"`
	OrderID                      string   `json:"OrderId"`
	OrderPrice                   float64  `json:"OrderPrice"`
	Status                       string   `json:"Status"`
	TrailingStopDistanceToMarket float64  `json:"TrailingStopDistanceToMarket"`
	TrailingStopStep             float64  `json:"TrailingStopStep"`
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
	Exchange                 *ExchangeOrder      `json:"Exchange,omitempty"`
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
	OrderID    string           `json:"OrderId"`
	TotalPrice float64          `json:"-"`
	Orders     []*OrderResponse `json:"Orders"`
	ErrorInfo  *struct {
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
	data, _, err := ma.Client.TradingApi.PlaceOrder(ma.Ctx,
		&saxo_openapi.TradingApiPlaceOrderOpts{Body: optional.NewInterface(&o)})
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

func (o *Order) WithType(orderType OrderType) *Order {
	o.OrderType = orderType
	return o
}

func (o *Order) WithBuySell(buySell BuySell) *Order {
	o.BuySell = buySell
	return o
}

func (o *Order) WithAmount(amount int) *Order {
	o.Amount = amount
	return o
}

func (o *Order) WithPrice(price float64) *Order {
	o.OrderPrice = price
	return o
}

func (o *Order) WithStopLimitPrice(stopLimitPrice float64) *Order {
	o.StopLimitPrice = stopLimitPrice
	return o
}

func (o *Order) WithDuration(duration DurationType) *Order {
	o.OrderDuration.DurationType = duration
	return o
}

func (o *Order) WithTakeProfit(sellingPrice float64) *Order {
	// Copy original order to keep uic/amout
	takeProfitOrder := *o
	takeProfitOrder.OrderType = Limit
	takeProfitOrder.BuySell = Sell
	takeProfitOrder.OrderDuration.DurationType = GoodTillCancel
	takeProfitOrder.OrderPrice = sellingPrice
	o.Orders = append(o.Orders, takeProfitOrder)
	return o
}

func (o *Order) WithStopLoss(sellingPrice float64) *Order {
	// Copy original order to keep uic/amout
	stopLossOrder := *o
	stopLossOrder.OrderType = StopIfTraded
	stopLossOrder.BuySell = Sell
	stopLossOrder.OrderDuration.DurationType = GoodTillCancel
	stopLossOrder.OrderPrice = sellingPrice
	o.Orders = append(o.Orders, stopLossOrder)
	return o
}

func (o *Order) WithStopLossTrailingStop(sellingPrice, distanceToMarket, step float64) *Order {
	// Copy original order to keep uic/amout
	trailingStop := *o
	trailingStop.OrderType = TrailingStopIfTraded
	trailingStop.BuySell = Sell
	trailingStop.OrderDuration.DurationType = GoodTillCancel
	trailingStop.TrailingStopDistanceToMarket = distanceToMarket
	trailingStop.TrailingStopStep = step
	trailingStop.OrderPrice = sellingPrice
	o.Orders = append(o.Orders, trailingStop)
	return o
}

func (o *Order) WithStopLossStopLimit(sellingPrice, stopLimitPrice float64) *Order {
	// Copy original order to keep uic/amout
	stopLimit := *o
	stopLimit.OrderType = StopLimit
	stopLimit.BuySell = Sell
	stopLimit.OrderDuration.DurationType = GoodTillCancel
	stopLimit.StopLimitPrice = stopLimitPrice
	stopLimit.OrderPrice = sellingPrice
	o.Orders = append(o.Orders, stopLimit)
	return o
}
