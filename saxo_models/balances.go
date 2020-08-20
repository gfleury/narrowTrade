package saxo_models

import "github.com/mitchellh/mapstructure"

type Balance struct {
	CalculationReliability string `json:"CalculationReliability"`
	CashBalance            int    `json:"CashBalance"`
	ChangesScheduled       bool   `json:"ChangesScheduled"`
	ClosedPositionsCount   int    `json:"ClosedPositionsCount"`
	CollateralAvailable    int    `json:"CollateralAvailable"`
	CollateralCreditValue  struct {
		Line           int `json:"Line"`
		UtilizationPct int `json:"UtilizationPct"`
	} `json:"CollateralCreditValue"`
	CostToClosePositions int    `json:"CostToClosePositions"`
	Currency             string `json:"Currency"`
	CurrencyDecimals     int    `json:"CurrencyDecimals"`
	InitialMargin        struct {
		CollateralAvailable   int `json:"CollateralAvailable"`
		CollateralCreditValue struct {
			Line           int `json:"Line"`
			UtilizationPct int `json:"UtilizationPct"`
		} `json:"CollateralCreditValue"`
		MarginAvailable              int `json:"MarginAvailable"`
		MarginCollateralNotAvailable int `json:"MarginCollateralNotAvailable"`
		MarginUsedByCurrentPositions int `json:"MarginUsedByCurrentPositions"`
		MarginUtilizationPct         int `json:"MarginUtilizationPct"`
		NetEquityForMargin           int `json:"NetEquityForMargin"`
	} `json:"InitialMargin"`
	IsPortfolioMarginModelSimple     bool `json:"IsPortfolioMarginModelSimple"`
	MarginAvailableForTrading        int  `json:"MarginAvailableForTrading"`
	MarginCollateralNotAvailable     int  `json:"MarginCollateralNotAvailable"`
	MarginExposureCoveragePct        int  `json:"MarginExposureCoveragePct"`
	MarginNetExposure                int  `json:"MarginNetExposure"`
	MarginUsedByCurrentPositions     int  `json:"MarginUsedByCurrentPositions"`
	MarginUtilizationPct             int  `json:"MarginUtilizationPct"`
	NetEquityForMargin               int  `json:"NetEquityForMargin"`
	NetPositionsCount                int  `json:"NetPositionsCount"`
	NonMarginPositionsValue          int  `json:"NonMarginPositionsValue"`
	OpenPositionsCount               int  `json:"OpenPositionsCount"`
	OptionPremiumsMarketValue        int  `json:"OptionPremiumsMarketValue"`
	OrdersCount                      int  `json:"OrdersCount"`
	OtherCollateral                  int  `json:"OtherCollateral"`
	SettlementValue                  int  `json:"SettlementValue"`
	TotalValue                       int  `json:"TotalValue"`
	TransactionsNotBooked            int  `json:"TransactionsNotBooked"`
	UnrealizedMarginClosedProfitLoss int  `json:"UnrealizedMarginClosedProfitLoss"`
	UnrealizedMarginOpenProfitLoss   int  `json:"UnrealizedMarginOpenProfitLoss"`
	UnrealizedMarginProfitLoss       int  `json:"UnrealizedMarginProfitLoss"`
	UnrealizedPositionsValue         int  `json:"UnrealizedPositionsValue"`
}

func (ma *ModeledAPI) GetBalanceMe() (*Balance, error) {
	b := &Balance{}
	ma.Throttle()
	data, _, err := ma.Client.PortfolioApi.GetBalance(ma.Ctx)
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, b)
	return b, err
}
