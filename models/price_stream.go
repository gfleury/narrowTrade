package models

import (
	"fmt"
	"log"
	"time"

	"github.com/antihax/optional"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"

	"github.com/gfleury/narrowTrade/saxo_openapi"
)

type PriceStreamRequest struct {
	Arguments   PriceRequestArguments `json:"Arguments"`
	ContextID   string                `json:"ContextId"`
	ReferenceID string                `json:"ReferenceId"`
}

type PriceRequestArguments struct {
	Uic       int32  `json:"Uic"`
	AssetType string `json:"AssetType"`
}

type PriceStreamResponse struct {
	ContextID         string        `json:"ContextId"`
	Format            string        `json:"Format"`
	InactivityTimeout int           `json:"InactivityTimeout"`
	ReferenceID       string        `json:"ReferenceId"`
	RefreshRate       int           `json:"RefreshRate"`
	Snapshot          PriceSnapshot `json:"Snapshot"`
	State             string        `json:"State"`
}

type PriceSnapshot struct {
	AssetType   string `json:"AssetType"`
	LastUpdated string `json:"LastUpdated"`
	PriceSource string `json:"PriceSource"`
	Quote       Quote  `json:"Quote"`
	Uic         int    `json:"Uic"`
}

func (ma *ModeledAPI) GetPriceStream(i Instrument) (chan Quote, error) {
	pr := &PriceStreamResponse{}

	contextID := fmt.Sprintf("explorer_%d", time.Now().Unix())

	ma.Throttle()
	data, _, err := ma.Client.TradingApi.AddSubscriptionAsyncTrading6(ma.Ctx,
		&saxo_openapi.TradingApiAddSubscriptionAsyncTrading6Opts{
			Body: optional.NewInterface(PriceStreamRequest{
				Arguments: PriceRequestArguments{
					AssetType: string(i.GetAssetType()),
					Uic:       i.GetID(),
				},
				ReferenceID: "C_857",
				ContextID:   contextID,
			}),
		})
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, pr)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func createSubscriptioWebSocket(connectionToken, authorization, contextID string) error {
	c, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf(
			"wss://streaming.saxotrader.com/sim/openapi/streaming/connection/connect?transport=webSockets&connectionToken=%s&authorization=%s&context=%s&tid=2",
			connectionToken, authorization, contextID), nil)
	if err != nil {
		return err
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	return nil
}
