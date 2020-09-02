package analysis

import "github.com/gfleury/narrowTrade/models"

type IndicatorName string

const (
	BBANDS IndicatorName = "bbands"
	EMA                  = "ema"
)

type InstrumentAnalysis struct {
	BuyRatings  int `json:"ratingBuy"`
	HoldRatings int `json:"ratingHold"`
	SellRatings int `json:"ratingSell"`
}

type Quote struct {
	Symbol                string  `json:"symbol"`
	CompanyName           string  `json:"companyName"`
	CalculationPrice      string  `json:"calculationPrice"`
	Open                  float64 `json:"open"`
	OpenTime              string  `json:"openTime"`
	Close                 float64 `json:"close"`
	CloseTime             string  `json:"closeTime"`
	High                  float64 `json:"high"`
	Low                   float64 `json:"low"`
	LatestPrice           float64 `json:"latestPrice"`
	LatestSource          string  `json:"latestSource"`
	LatestTime            string  `json:"latestTime"`
	LatestUpdate          string  `json:"latestUpdate"`
	LatestVolume          int     `json:"latestVolume"`
	RealtimePrice         float64 `json:"RealtimePrice"`
	RealtimeSize          int     `json:"RealtimeSize"`
	LastUpdated           string  `json:"LastUpdated"`
	DelayedPrice          float64 `json:"delayedPrice"`
	DelayedPriceTime      string  `json:"delayedPriceTime"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	ExtendedPriceTime     string  `json:"extendedPriceTime"`
	PreviousClose         float64 `json:"previousClose"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	MarketPercent         float64 `json:"MarketPercent"`
	Volume                int     `json:"Volume"`
	AvgTotalVolume        int     `json:"avgTotalVolume"`
	BidPrice              float64 `json:"BidPrice"`
	BidSize               int     `json:"BidSize"`
	AskPrice              float64 `json:"AskPrice"`
	AskSize               int     `json:"AskSize"`
	MarketCap             int     `json:"marketCap"`
	Week52High            float64 `json:"week52High"`
	Week52Low             float64 `json:"week52Low"`
	YTDChange             float64 `json:"ytdChange"`
	PERatio               float64 `json:"peRatio"`
}

type InstrumentAnalyser interface {
	Analysis() []InstrumentAnalysis
	Indicator(models.Instrument, IndicatorName) ([]float64, error)
	Quote(models.Instrument) (*Quote, error)
	OneAnalysis(models.Instrument) (*InstrumentAnalysis, error)
	Analyse([]models.Instrument) error
}

func (in IndicatorName) String() string {
	return string(in)
}
