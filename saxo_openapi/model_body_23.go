/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body23 struct {
	AccountKey string `json:"AccountKey,omitempty"`
	AdvancedOptions *Openapicsv2cashmanagementwithdrawalsWithdrawIBanSwiftAdvancedOptions `json:"AdvancedOptions,omitempty"`
	Amount float64 `json:"Amount,omitempty"`
	ClearingCode string `json:"ClearingCode,omitempty"`
	Currency string `json:"Currency,omitempty"`
	Iban string `json:"Iban,omitempty"`
	MessageToBeneficiary string `json:"MessageToBeneficiary,omitempty"`
	ReceivingCountryCode string `json:"ReceivingCountryCode,omitempty"`
	Swift string `json:"Swift,omitempty"`
	WithdrawalReasonId int32 `json:"WithdrawalReasonId,omitempty"`
}
