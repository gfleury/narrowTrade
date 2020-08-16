/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body38 struct {
	AccountValueProtectionLimit float64 `json:"AccountValueProtectionLimit,omitempty"`
	BenchmarkInstrument *Openapiportv1accountsAccountKeyBenchmarkInstrument `json:"BenchmarkInstrument,omitempty"`
	DisplayName string `json:"DisplayName,omitempty"`
	UseCashPositionsAsMarginCollateral bool `json:"UseCashPositionsAsMarginCollateral,omitempty"`
}
