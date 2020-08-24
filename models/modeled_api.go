package models

import (
	"context"
	"sync"
	"time"

	"github.com/gfleury/narrowTrade/saxo_openapi"
)

type ModeledAPI struct {
	Ctx         context.Context
	Client      *saxo_openapi.APIClient
	lastAPICall time.Time
	m           sync.Mutex
}

func (ma *ModeledAPI) Throttle() {
	ma.m.Lock()
	timePassed := time.Since(ma.lastAPICall)
	ma.m.Unlock()
	if timePassed < time.Second {
		time.Sleep(time.Second - timePassed)
	}
}

func (ma *ModeledAPI) UpdateLastCall() {
	ma.m.Lock()
	ma.lastAPICall = time.Now()
	ma.m.Unlock()
}
