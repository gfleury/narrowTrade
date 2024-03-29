/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Openapichartv1chartssubscriptionsArguments struct {
	AssetType   []string      `json:"AssetType,omitempty"`
	Count       int32         `json:"Count,omitempty"`
	FieldGroups []interface{} `json:"FieldGroups,omitempty"`
	Horizon     int32         `json:"Horizon,omitempty"`
	Uic         int32         `json:"Uic,omitempty"`
}
