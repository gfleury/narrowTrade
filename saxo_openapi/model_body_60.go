/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body60 struct {
	Expiries []interface{} `json:"Expiries,omitempty"`
	MaxStrikesPerExpiry int32 `json:"MaxStrikesPerExpiry,omitempty"`
}
