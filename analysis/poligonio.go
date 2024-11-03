package analysis

import (
	"context"
	"log"
	"time"

	"github.com/gfleury/narrowTrade/models"

	polygon "github.com/polygon-io/client-go/rest"
	polygon_models "github.com/polygon-io/client-go/rest/models"
)

type PolygonIOAnalyser struct {
	c *polygon.Client
}

func NewPolygonIOAnalyser(apiKey string) *PolygonIOAnalyser {
	return &PolygonIOAnalyser{ // init client
		c: polygon.New(apiKey)}
}

func (a *PolygonIOAnalyser) Init() {
	// Initialize any required resources or configurations
}

func getDate(years int, months int, days int) time.Time {
	now := time.Now()
	return now.AddDate(years, months, days)
}

func (a *PolygonIOAnalyser) HistoricalPrices(symbol models.Instrument) ([]float64, error) {
	// Fetch historical prices from the API
	log.Println("Fetching historical prices from Polygon.io for", symbol)

	timeSpan := polygon_models.Day
	seriesType := polygon_models.Close
	order := polygon_models.Asc
	limit := 1000

	results, err := a.c.GetEMA(context.Background(),
		&polygon_models.GetEMAParams{
			Ticker:     symbol.GetSymbolSimple(),
			Timespan:   &timeSpan,
			SeriesType: &seriesType,
			Adjusted:   new(bool),
			Limit:      &limit,
			Order:      &order,
		},
	)

	if err != nil {
		return nil, err
	}

	prices := make([]float64, len(results.Results.Values))
	for _, value := range results.Results.Values {
		prices = append(prices, value.Value)
	}

	return prices, nil
}
