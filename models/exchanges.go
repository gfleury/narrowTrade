package models

import (
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
)

type ListExchanges struct {
	Count int        `json:"__count"`
	Data  []Exchange `json:"Data"`
}

type ExchangeSessions struct {
	EndTime   string `json:"EndTime"`
	StartTime string `json:"StartTime"`
	State     string `json:"State"`
}

type Exchange struct {
	AllDay               bool               `json:"AllDay"`
	CountryCode          string             `json:"CountryCode"`
	Currency             string             `json:"Currency"`
	ExchangeID           string             `json:"ExchangeId"`
	ExchangeSessions     []ExchangeSessions `json:"ExchangeSessions"`
	Mic                  string             `json:"Mic"`
	Name                 string             `json:"Name"`
	PriceSourceName      string             `json:"PriceSourceName,omitempty"`
	TimeZoneID           string             `json:"TimeZoneId"`
	TimeZone             int                `json:"TimeZone"`
	TimeZoneAbbreviation string             `json:"TimeZoneAbbreviation"`
	TimeZoneOffset       string             `json:"TimeZoneOffset"`

	// Internal use
	IsOpen bool `json:",omitempty"`
}

func (ma *ModeledAPI) GetExchange(id string) (*Exchange, error) {
	b := &Exchange{}
	ma.Throttle()
	ctx, cancel := context.WithTimeout(ma.Ctx, 15*time.Second)
	defer cancel()

	data, _, err := ma.Client.ReferenceDataApi.GetExchange(ctx, id)
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, b)
	now := time.Now()
	if err == nil {
		for _, session := range b.ExchangeSessions {
			start, err := time.Parse(time.RFC3339, session.StartTime)
			if err != nil {
				continue
			}

			end, err := time.Parse(time.RFC3339, session.EndTime)
			if err != nil {
				continue
			}
			if (now.After(start) && now.Before(end)) && session.State != "Closed" {
				b.IsOpen = true
			}
		}

	}
	return b, err
}
