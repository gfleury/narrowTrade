package saxo_models

import (
	"github.com/mitchellh/mapstructure"
)

type Accounts struct {
	Data []struct {
		AccountGroupKey                       string   `json:"AccountGroupKey"`
		AccountID                             string   `json:"AccountId"`
		AccountKey                            string   `json:"AccountKey"`
		AccountSubType                        string   `json:"AccountSubType"`
		AccountType                           string   `json:"AccountType"`
		Active                                bool     `json:"Active"`
		CanUseCashPositionsAsMarginCollateral bool     `json:"CanUseCashPositionsAsMarginCollateral"`
		CfdBorrowingCostsActive               bool     `json:"CfdBorrowingCostsActive"`
		ClientID                              string   `json:"ClientId"`
		ClientKey                             string   `json:"ClientKey"`
		CreationDate                          string   `json:"CreationDate"`
		Currency                              string   `json:"Currency"`
		CurrencyDecimals                      int      `json:"CurrencyDecimals"`
		DirectMarketAccess                    bool     `json:"DirectMarketAccess"`
		IndividualMargining                   bool     `json:"IndividualMargining"`
		IsCurrencyConversionAtSettlementTime  bool     `json:"IsCurrencyConversionAtSettlementTime"`
		IsMarginTradingAllowed                bool     `json:"IsMarginTradingAllowed"`
		IsShareable                           bool     `json:"IsShareable"`
		IsTrialAccount                        bool     `json:"IsTrialAccount"`
		LegalAssetTypes                       []string `json:"LegalAssetTypes"`
		MarginCalculationMethod               string   `json:"MarginCalculationMethod"`
		Sharing                               []string `json:"Sharing"`
		SupportsAccountValueProtectionLimit   bool     `json:"SupportsAccountValueProtectionLimit"`
		UseCashPositionsAsMarginCollateral    bool     `json:"UseCashPositionsAsMarginCollateral"`
	} `json:"Data"`
}

func (ma *ModeledAPI) GetAccounts() (*Accounts, error) {
	a := &Accounts{}
	ma.Throttle()
	data, _, err := ma.Client.PortfolioApi.GetAccounts(ma.Ctx, nil)
	defer ma.UpdateLastCall()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(data, a)
	return a, err
}

func (a *Accounts) GetClientKeyMe() string {
	return a.Data[0].ClientKey
}

func (a *Accounts) GetClientKey(idx int) string {
	return a.Data[idx].ClientKey
}

func (a *Accounts) GetAccountKeyMe() string {
	return a.Data[0].AccountKey
}

func (a *Accounts) GetAccountKey(idx int) string {
	return a.Data[idx].AccountKey
}
