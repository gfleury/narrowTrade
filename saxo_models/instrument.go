package saxo_models

import (
	"fmt"

	"github.com/antihax/optional"
	"github.com/gfleury/intensiveTrade/saxo_openapi"
	"github.com/mitchellh/mapstructure"
)

type Instrument interface {
	GetSymbol() string
	GetID() int
	GetPrice() float64
	GetOrder(bs BuySell, ot OrderType, amount int) *Order
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
	Identifier     int       `json:"Identifier"`
	IssuerCountry  string    `json:"IssuerCountry"`
	PrimaryListing int       `json:"PrimaryListing"`
	SummaryType    string    `json:"SummaryType"`
	Symbol         string    `json:"Symbol"`
	TradableAs     []string  `json:"TradableAs"`
}

func (ma *ModeledAPI) GetInstrument(symbol string) (Instrument, error) {
	i := &SaxoInstrument{}
	data, _, err := ma.Client.ReferenceDataApi.GetSummaries(ma.Ctx, &saxo_openapi.ReferenceDataApiGetSummariesOpts{Keywords: optional.NewString(symbol)})
	if err != nil {
		return nil, err
	}

	is := data.(map[string]interface{})["Data"].([]interface{})
	if len(is) < 1 {
		return nil, fmt.Errorf("No instrument found with such symbol")
	}

	err = mapstructure.Decode(is[0], i)
	return i, err
}

func (i *SaxoInstrument) GetSymbol() string {
	return i.Symbol
}

func (i *SaxoInstrument) GetID() int {
	return i.Identifier
}

func (i *SaxoInstrument) GetPrice() float64 {
	return 0.1
}

func (i *SaxoInstrument) GetOrder(bs BuySell, ot OrderType, amount int) *Order {
	return &Order{
		Uic:         i.Identifier,
		AssetType:   i.AssetType,
		Amount:      amount,
		BuySell:     bs,
		OrderType:   ot,
		ManualOrder: true,
	}
}
