/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body65 struct {
	AccountKey string `json:"AccountKey,omitempty"`
	FieldGroups []string `json:"FieldGroups,omitempty"`
	Legs []interface{} `json:"Legs,omitempty"`
}
