package models

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/gfleury/narrowTrade/utils"

	"github.com/antihax/optional"
	"github.com/mitchellh/mapstructure"
)

type Instrument interface {
	GetSymbol() string
	GetSymbolSimple() string
	GetExchangeID() string
	GetID() int32
	GetAssetType() AssetType
	GetOrder() *Order
}

type InstrumentDetails interface {
	Instrument
	CalculatePriceWithThickSize(price float64, percentage float64) float64
	GetDecimals() int
	GetMinimumTradeSize() int
	GetMinimumOrderValue() int
}

type SaxoInstruments struct {
	Data []Instrument `json:"Data"`
}

type SaxoInstrument struct {
	AssetType      AssetType `json:"AssetType"`
	CurrencyCode   string    `json:"CurrencyCode"`
	Description    string    `json:"Description"`
	ExchangeID     string    `json:"ExchangeId"`
	GroupID        int       `json:"GroupId"`
	Identifier     int32     `json:"Identifier"`
	IssuerCountry  string    `json:"IssuerCountry"`
	PrimaryListing int       `json:"PrimaryListing"`
	SummaryType    string    `json:"SummaryType"`
	Symbol         string    `json:"Symbol"`
	TradableAs     []string  `json:"TradableAs"`
}

type SaxoInstrumentDetails struct {
	AmountDecimals        int                  `json:"AmountDecimals"`
	AssetType             AssetType            `json:"AssetType"`
	CurrencyCode          string               `json:"CurrencyCode"`
	DefaultAmount         int                  `json:"DefaultAmount"`
	DefaultSlippage       int                  `json:"DefaultSlippage"`
	DefaultSlippageType   string               `json:"DefaultSlippageType"`
	Description           string               `json:"Description"`
	Exchange              Exchange             `json:"Exchange"`
	Format                Format               `json:"Format"`
	GroupID               int                  `json:"GroupId"`
	IncrementSize         int                  `json:"IncrementSize"`
	IsBarrierEqualsStrike bool                 `json:"IsBarrierEqualsStrike"`
	IsComplex             bool                 `json:"IsComplex"`
	IsTradable            bool                 `json:"IsTradable"`
	LotSizeType           string               `json:"LotSizeType"`
	MinimumOrderValue     int                  `json:"MinimumOrderValue"`
	MinimumTradeSize      int                  `json:"MinimumTradeSize"`
	OrderDistances        OrderDistances       `json:"OrderDistances"`
	PriceCurrency         string               `json:"PriceCurrency"`
	PriceToContractFactor int                  `json:"PriceToContractFactor"`
	PrimaryListing        int                  `json:"PrimaryListing"`
	RelatedInstruments    []RelatedInstruments `json:"RelatedInstruments"`
	StandardAmounts       []int                `json:"StandardAmounts"`
	SupportedOrderTypes   []string             `json:"SupportedOrderTypes"`
	Symbol                string               `json:"Symbol"`
	TickSizeScheme        TickSizeScheme       `json:"TickSizeScheme"`
	TradableAs            []string             `json:"TradableAs"`
	TradableOn            []string             `json:"TradableOn"`
	TradingSignals        string               `json:"TradingSignals"`
	TurboDirection        string               `json:"TurboDirection"`
	Uic                   int32                `json:"Uic"`
}

type Format struct {
	Decimals      int `json:"Decimals"`
	OrderDecimals int `json:"OrderDecimals"`
}

type OrderDistances struct {
	EntryDefaultDistance          float64 `json:"EntryDefaultDistance"`
	EntryDefaultDistanceType      string  `json:"EntryDefaultDistanceType"`
	StopLimitDefaultDistance      int     `json:"StopLimitDefaultDistance"`
	StopLimitDefaultDistanceType  string  `json:"StopLimitDefaultDistanceType"`
	StopLossDefaultDistance       float64 `json:"StopLossDefaultDistance"`
	StopLossDefaultDistanceType   string  `json:"StopLossDefaultDistanceType"`
	StopLossDefaultEnabled        bool    `json:"StopLossDefaultEnabled"`
	StopLossDefaultOrderType      string  `json:"StopLossDefaultOrderType"`
	TakeProfitDefaultDistance     float64 `json:"TakeProfitDefaultDistance"`
	TakeProfitDefaultDistanceType string  `json:"TakeProfitDefaultDistanceType"`
	TakeProfitDefaultEnabled      bool    `json:"TakeProfitDefaultEnabled"`
}

type RelatedInstruments struct {
	AssetType string `json:"AssetType"`
	Uic       int    `json:"Uic"`
}

type Elements struct {
	HighPrice float64 `json:"HighPrice"`
	TickSize  float64 `json:"TickSize"`
}

type TickSizeScheme struct {
	DefaultTickSize float64    `json:"DefaultTickSize"`
	Elements        []Elements `json:"Elements"`
}

func (ma *ModeledAPI) GetInstrument(symbol string) (Instrument, error) {
	var assetType optional.Interface

	if key := ma.cache.Get(symbol); key != nil {
		return key.(Instrument), nil
	}

	i := &SaxoInstrument{}
	if len(symbol) == 6 {
		assetType = optional.NewInterface(FxSpot)
	} else {
		assetType = optional.NewInterface(Stock)
	}

	ma.Throttle()
	data, _, err := ma.Client.ReferenceDataApi.GetSummaries(ma.Ctx,
		&saxo_openapi.ReferenceDataApiGetSummariesOpts{
			Keywords:   optional.NewString(symbol),
			AssetTypes: assetType,
		})
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	is := data.(map[string]interface{})["Data"].([]interface{})
	if len(is) < 1 {
		return nil, fmt.Errorf("No instrument found with such symbol")
	}

	err = mapstructure.Decode(is[0], i)

	if err == nil {
		ma.cache.Put(symbol, i, 24*time.Hour)
	}

	return i, err
}

func (ma *ModeledAPI) GetInstrumentDetails(i Instrument) (InstrumentDetails, error) {
	d := &SaxoInstrumentDetails{}
	ma.Throttle()
	data, _, err := ma.Client.ReferenceDataApi.GetDetailsForOneInstrument(ma.Ctx,
		Stock,
		i.GetID(),
		nil)
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, d)
	return d, err
}

func (ma *ModeledAPI) GetInstrumentPrice(i Instrument) (float64, error) {
	iprice := &InstrumentPrice{}
	ma.Throttle()
	data, _, err := ma.Client.TradingApi.GetInfoPriceAsync(ma.Ctx, []string{string(i.GetAssetType())}, i.GetID(), nil)
	defer ma.UpdateLastCall()
	if err != nil {
		return 0, err
	}

	err = mapstructure.Decode(data, iprice)
	if err != nil {
		return 0, err
	}

	return iprice.Quote.Ask, nil
}

func (i *SaxoInstrument) GetAssetType() AssetType {
	return i.AssetType
}

func (i *SaxoInstrument) GetSymbol() string {
	return i.Symbol
}

func (i *SaxoInstrument) GetSymbolSimple() string {
	return strings.Split(i.GetSymbol(), ":")[0]
}

func (i *SaxoInstrument) GetID() int32 {
	return i.Identifier
}

func (i *SaxoInstrument) GetOrder() *Order {
	return &Order{
		Uic:         i.Identifier,
		AssetType:   i.AssetType,
		ManualOrder: true,
		OrderDuration: Duration{
			DurationType: DayOrder,
		},
	}
}

func (i *SaxoInstrument) GetExchangeID() string {
	return i.ExchangeID
}

// SaxoInstrumentDetails

func (i *SaxoInstrumentDetails) GetAssetType() AssetType {
	return i.AssetType
}

func (i *SaxoInstrumentDetails) GetSymbol() string {
	return i.Symbol
}

func (i *SaxoInstrumentDetails) GetSymbolSimple() string {
	return strings.Split(i.GetSymbol(), ":")[0]
}

func (i *SaxoInstrumentDetails) GetID() int32 {
	return i.Uic
}

func (i *SaxoInstrumentDetails) GetOrder() *Order {
	return &Order{
		Uic:         i.Uic,
		AssetType:   i.AssetType,
		ManualOrder: true,
		OrderDuration: Duration{
			DurationType: DayOrder,
		},
	}
}

func (i *SaxoInstrumentDetails) GetExchangeID() string {
	return i.Exchange.ExchangeID
}

func (id *SaxoInstrumentDetails) CalculatePriceWithThickSize(price float64, percentage float64) float64 {
	//orderPrice = Round((currentPrice- currentPrice*n/100)/tickSize)*tickSize
	tickSize := 0.1
	if id.TickSizeScheme.DefaultTickSize > 0 {
		tickSize = GetPriceTick(price, id.TickSizeScheme.Elements)
		if tickSize == 0 {
			tickSize = id.TickSizeScheme.DefaultTickSize
		}
	}

	return utils.Round(utils.RoundNearestTick(price, percentage, tickSize), id.GetDecimals())
}

func (id *SaxoInstrumentDetails) GetDecimals() int {
	return id.Format.OrderDecimals
}

func (id *SaxoInstrumentDetails) GetMinimumTradeSize() int {
	return id.MinimumTradeSize
}

func (id *SaxoInstrumentDetails) GetMinimumOrderValue() int {
	return id.MinimumOrderValue
}

func GetPriceTick(price float64, elements []Elements) float64 {
	if len(elements) < 2 {
		return 0
	}

	sort.Slice(elements,
		func(i, j int) bool {
			return elements[i].HighPrice < elements[j].HighPrice
		})

	// Check if its higher than the last element (highest)
	if price > elements[len(elements)-1].HighPrice {
		return elements[len(elements)-1].TickSize
	}

	// Check if its smaller than the first element (smallest)
	if price < elements[0].HighPrice {
		return elements[0].TickSize
	}

	for _, e := range elements {
		if price <= e.HighPrice {
			return e.TickSize
		}
	}
	return 0
}
