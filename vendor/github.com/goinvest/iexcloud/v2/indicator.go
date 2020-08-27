package iex

type Indicator struct {
	Indicator [][]float64 `json:"indicator"`
	Chart     []struct {
		Date           string  `json:"date"`
		Open           float64 `json:"open"`
		Close          float64 `json:"close"`
		High           float64 `json:"high"`
		Low            float64 `json:"low"`
		Volume         int     `json:"volume"`
		UOpen          float64 `json:"uOpen"`
		UClose         float64 `json:"uClose"`
		UHigh          float64 `json:"uHigh"`
		ULow           float64 `json:"uLow"`
		UVolume        int     `json:"uVolume"`
		Change         int     `json:"change"`
		ChangePercent  int     `json:"changePercent"`
		Label          string  `json:"label"`
		ChangeOverTime int     `json:"changeOverTime"`
	} `json:"chart"`
}
