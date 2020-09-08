package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/antihax/optional"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"

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

type QuoteStream struct {
	LastUpdated time.Time `json:"LastUpdated"`
	Quote       struct {
		Ask float64 `json:"Ask"`
		Bid float64 `json:"Bid"`
		Mid float64 `json:"Mid"`
	} `json:"Quote"`
}

func (ma *ModeledAPI) CreatePriceStream(i Instrument) (chan *QuoteStream, *websocket.Conn, error) {
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
		return nil, nil, err
	}

	err = mapstructure.Decode(data, pr)
	if err != nil {
		return nil, nil, err
	}

	authorizationToken := ""
	if tok, ok := ma.Ctx.Value(saxo_openapi.ContextOAuth2).(oauth2.TokenSource); ok {
		// We were able to grab an oauth2 token from the context
		var latestToken *oauth2.Token
		if latestToken, err = tok.Token(); err != nil {
			return nil, nil, err
		}
		authorizationToken = latestToken.AccessToken
	}

	wsChan, wsConn, err := createSubscriptioWebSocket(authorizationToken, pr.ContextID)
	if err != nil {
		return nil, nil, err
	}

	quoteChan := make(chan *QuoteStream)
	go func() {
		for {
			defer wsConn.Close()
			msg := <-wsChan
			if int(msg.PayloadSize) != len(msg.Payload) {
				log.Printf("Got corrupted msg %v", msg)
				continue
			}
			if string(msg.RID) != pr.ReferenceID {
				log.Printf("Got non expected ReferenceID msg %v", msg)
				continue
			}
			q := &QuoteStream{}

			err = json.Unmarshal(msg.Payload, q)
			if err != nil {
				log.Println("Unable to Unmarshal Paylog to Quote:", err)
			}
			log.Println(string(msg.Payload))
			quoteChan <- q
		}
	}()

	return quoteChan, wsConn, err
}
