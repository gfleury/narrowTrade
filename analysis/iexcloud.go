package analysis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/utils"

	iex "github.com/goinvest/iexcloud/v2"
)

type cacheEntry struct {
	value  interface{}
	expire time.Time
}

type IEXAnalyser struct {
	Ctx    context.Context
	Client *iex.Client
	cache  map[string]*cacheEntry
}

func (a *IEXAnalyser) Init() {
	a.cache = map[string]*cacheEntry{}
}

func (a *IEXAnalyser) Analysis() []InstrumentAnalysis {
	return []InstrumentAnalysis{}
}

func (a *IEXAnalyser) Indicator(i models.Instrument, indicatorName IndicatorName) ([]float64, error) {
	// Check if we have the data cached
	cacheKey := cacheKey(i.GetSymbolSimple(), indicatorName.String())
	if cachedValue, ok := a.cache[cacheKey]; ok {
		if !time.Now().After(cachedValue.expire) {
			return cachedValue.value.([]float64), nil
		}
	}
	indicator, err := a.Client.Indicator(a.Ctx, i.GetSymbolSimple(), indicatorName.String(), "3m")
	if err != nil {
		return nil, err
	}

	switch indicatorName {
	case EMA:
		a.cache[cacheKey] = &cacheEntry{indicator.Indicator[0], time.Now().Add(6 * time.Hour)}
	case BBANDS:
		a.cache[cacheKey] = &cacheEntry{utils.IEXBBandsReduction(indicator, 5), time.Now().Add(6 * time.Hour)}
	}
	return a.cache[cacheKey].value.([]float64), nil
}

func (a *IEXAnalyser) Quote(i models.Instrument) (*Quote, error) {
	// Check if we have the data cached
	cacheKey := cacheKey(i.GetSymbolSimple(), "quote")
	if cachedValue, ok := a.cache[cacheKey]; ok {
		if !time.Now().After(cachedValue.expire) {
			return cachedValue.value.(*Quote), nil
		}
	}
	quote, err := a.Client.Quote(a.Ctx, i.GetSymbolSimple())
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(quote)
	if err != nil {
		return nil, err
	}
	aQuote := &Quote{}
	err = json.Unmarshal(data, aQuote)
	if err != nil {
		return nil, err
	}

	a.cache[cacheKey] = &cacheEntry{aQuote, time.Now().Add(60 * time.Second)}

	return a.cache[cacheKey].value.(*Quote), nil
}

func (a *IEXAnalyser) OneAnalysis(i models.Instrument) (*InstrumentAnalysis, error) {
	// Check if we have the data cached
	cacheKey := cacheKey(i.GetSymbolSimple(), "recommendation")
	if cachedValue, ok := a.cache[cacheKey]; ok {
		if !time.Now().After(cachedValue.expire) {
			return cachedValue.value.(*InstrumentAnalysis), nil
		}
	}
	recommendations, err := a.Client.RecommendationTrends(a.Ctx, i.GetSymbolSimple())
	if err != nil {
		return nil, err
	}

	recommendation := utils.IEXRecomendationReduce(recommendations)

	data, err := json.Marshal(recommendation)
	if err != nil {
		return nil, err
	}
	instrumentAnalysis := &InstrumentAnalysis{}
	err = json.Unmarshal(data, instrumentAnalysis)
	if err != nil {
		return nil, err
	}

	a.cache[cacheKey] = &cacheEntry{instrumentAnalysis, time.Now().Add(24 * time.Hour)}

	return a.cache[cacheKey].value.(*InstrumentAnalysis), nil
}

func (a *IEXAnalyser) Analyse(is []models.Instrument) error {
	for _, i := range is {
		_, err := a.OneAnalysis(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func cacheKey(symbol, cacheName string) string {
	return fmt.Sprintf("%s:%s", symbol, cacheName)
}
