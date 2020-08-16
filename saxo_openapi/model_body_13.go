/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body13 struct {
	AccountInformation *Openapicmv1signupsAccountInformation `json:"AccountInformation,omitempty"`
	BankInformation *Openapicmv1signupsBankInformation `json:"BankInformation,omitempty"`
	FinlandData *Openapicmv1signupsFinlandData `json:"FinlandData,omitempty"`
	ItalyData *Openapicmv1signupsItalyData `json:"ItalyData,omitempty"`
	OnboardingInformation *Openapicmv1signupsOnboardingInformation `json:"OnboardingInformation,omitempty"`
	PensionData *Openapicmv1signupsPensionData `json:"PensionData,omitempty"`
	PersonalInformation *Openapicmv1signupsPersonalInformation `json:"PersonalInformation,omitempty"`
	ProfileInformation *Openapicmv1signupsProfileInformation `json:"ProfileInformation,omitempty"`
	RegulatoryInformation *Openapicmv1signupsRegulatoryInformation `json:"RegulatoryInformation,omitempty"`
	SingaporeData *Openapicmv1signupsSingaporeData `json:"SingaporeData,omitempty"`
	SwitzerlandData *Openapicmv1signupsSwitzerlandData `json:"SwitzerlandData,omitempty"`
}
