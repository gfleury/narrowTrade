package models

import (
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
)

type InfoPrice struct {
	AssetType   string `json:"AssetType"`
	LastUpdated string `json:"LastUpdated"`
	PriceSource string `json:"PriceSource"`
	Quote       Quote  `json:"Quote"`
	Uic         int    `json:"Uic"`
}

func (ma *ModeledAPI) GetInfoPrice(i Instrument) (*InfoPrice, error) {
	b := &InfoPrice{}
	ma.Throttle()
	ctx, cancel := context.WithTimeout(ma.Ctx, 15*time.Second)
	defer cancel()

	data, _, err := ma.Client.TradingApi.GetInfoPriceAsync(ctx,
		[]string{string(i.GetAssetType())}, i.GetID(), nil)
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, b)
	return b, err
}
