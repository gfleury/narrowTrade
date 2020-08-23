package models

import (
	"context"
	"time"

	"github.com/gfleury/narrowTrade/saxo_openapi"
)

type ModeledAPI struct {
	Ctx         context.Context
	Client      *saxo_openapi.APIClient
	lastAPICall time.Time
}

func (ma *ModeledAPI) Throttle() {
	timePassed := time.Now().Sub(ma.lastAPICall)
	if timePassed < time.Second {
		time.Sleep(time.Second - timePassed)
	}
}

func (ma *ModeledAPI) UpdateLastCall() {
	ma.lastAPICall = time.Now()
}
