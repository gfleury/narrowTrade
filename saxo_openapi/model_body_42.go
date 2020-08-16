/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body42 struct {
	AccountValueProtectionLimit float64 `json:"AccountValueProtectionLimit,omitempty"`
	ForceOpenDefaultValue bool `json:"ForceOpenDefaultValue,omitempty"`
	NewPositionNettingMode []string `json:"NewPositionNettingMode,omitempty"`
}
