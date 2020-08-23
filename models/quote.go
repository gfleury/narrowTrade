package models

type InstrumentPrice struct {
	AssetType   string `json:"AssetType"`
	LastUpdated string `json:"LastUpdated"`
	PriceSource string `json:"PriceSource"`
	Quote       Quote  `json:"Quote"`
	Uic         int    `json:"Uic"`
}

type Quote struct {
	Amount           int     `json:"Amount"`
	Ask              float64 `json:"Ask"`
	Bid              float64 `json:"Bid"`
	DelayedByMinutes int     `json:"DelayedByMinutes"`
	ErrorCode        string  `json:"ErrorCode"`
	Mid              float64 `json:"Mid"`
	PriceSource      string  `json:"PriceSource"`
	PriceSourceType  string  `json:"PriceSourceType"`
	PriceTypeAsk     string  `json:"PriceTypeAsk"`
	PriceTypeBid     string  `json:"PriceTypeBid"`
}
