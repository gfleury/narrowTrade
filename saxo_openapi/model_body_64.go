/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body64 struct {
	AccountKey string `json:"AccountKey,omitempty"`
	Amount int32 `json:"Amount,omitempty"`
	AppHint int32 `json:"AppHint,omitempty"`
	AssetType []string `json:"AssetType,omitempty"`
	Uic int32 `json:"Uic,omitempty"`
}
