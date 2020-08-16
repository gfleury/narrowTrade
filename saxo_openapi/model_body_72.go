/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body72 struct {
	AccountKey string `json:"AccountKey,omitempty"`
	DecisionMakerUserID int32 `json:"DecisionMakerUserID,omitempty"`
	ExternalReference string `json:"ExternalReference,omitempty"`
	FieldGroups []interface{} `json:"FieldGroups,omitempty"`
	Legs []interface{} `json:"Legs,omitempty"`
	ManualOrder bool `json:"ManualOrder,omitempty"`
	OrderDuration *interface{} `json:"OrderDuration,omitempty"`
	OrderPrice float64 `json:"OrderPrice,omitempty"`
	OrderType []string `json:"OrderType,omitempty"`
	TraderId string `json:"TraderId,omitempty"`
}
