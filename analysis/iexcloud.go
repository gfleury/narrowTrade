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

type IEXAnalyser struct {
	Ctx    context.Context
	Client *iex.Client
	cache  map[string]*utils.CacheEntry
}

func (a *IEXAnalyser) Init() {
	a.cache = map[string]*utils.CacheEntry{}
}

func (a *IEXAnalyser) Analysis() []InstrumentAnalysis {
	return []InstrumentAnalysis{}
}

func (a *IEXAnalyser) Indicator(i models.Instrument, indicatorName IndicatorName) ([]float64, error) {
	ctx, cancel := context.WithTimeout(a.Ctx, 15*time.Second)
	defer cancel()

	// Check if we have the data cached
	cacheKey := cacheKey(i.GetSymbolSimple(), indicatorName.String())
	if cachedValue, ok := a.cache[cacheKey]; ok {
		if !time.Now().After(cachedValue.Expire) {
			return cachedValue.Value.([]float64), nil
		}
	}
	indicator, err := a.Client.Indicator(ctx, i.GetSymbolSimple(), indicatorName.String(), "3m")
	if err != nil {
		return nil, err
	}

	switch indicatorName {
	case EMA:
		a.cache[cacheKey] = &utils.CacheEntry{indicator.Indicator[0], time.Now().Add(6 * time.Hour)}
	case BBANDS:
		a.cache[cacheKey] = &utils.CacheEntry{utils.IEXBBandsReduction(indicator, 5), time.Now().Add(6 * time.Hour)}
	}
	return a.cache[cacheKey].Value.([]float64), nil
}

func (a *IEXAnalyser) Quote(i models.Instrument) (*Quote, error) {
	ctx, cancel := context.WithTimeout(a.Ctx, 15*time.Second)
	defer cancel()

	// Check if we have the data cached
	cacheKey := cacheKey(i.GetSymbolSimple(), "quote")
	if cachedValue, ok := a.cache[cacheKey]; ok {
		if !time.Now().After(cachedValue.Expire) {
			return cachedValue.Value.(*Quote), nil
		}
	}
	quote, err := a.Client.Quote(ctx, i.GetSymbolSimple())
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

	a.cache[cacheKey] = &utils.CacheEntry{aQuote, time.Now().Add(60 * time.Second)}

	return a.cache[cacheKey].Value.(*Quote), nil
}

func (a *IEXAnalyser) OneAnalysis(i models.Instrument) (*InstrumentAnalysis, error) {
	ctx, cancel := context.WithTimeout(a.Ctx, 15*time.Second)
	defer cancel()

	// Check if we have the data cached
	cacheKey := cacheKey(i.GetSymbolSimple(), "recommendation")
	if cachedValue, ok := a.cache[cacheKey]; ok {
		if !time.Now().After(cachedValue.Expire) {
			return cachedValue.Value.(*InstrumentAnalysis), nil
		}
	}
	recommendations, err := a.Client.RecommendationTrends(ctx, i.GetSymbolSimple())
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

	a.cache[cacheKey] = &utils.CacheEntry{instrumentAnalysis, time.Now().Add(24 * time.Hour)}

	return a.cache[cacheKey].Value.(*InstrumentAnalysis), nil
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

func IEXConfigLoad(data []byte) (string, string, error) {
	iexConf := &struct {
		Token string
		URL   string
	}{}

	err := json.Unmarshal(data, iexConf)
	if err != nil {
		return "", "", err
	}

	return iexConf.Token, iexConf.URL, nil
}
