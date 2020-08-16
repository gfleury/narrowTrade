/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body76 struct {
	EmailAddress string `json:"EmailAddress,omitempty"`
	NotifyWithMail bool `json:"NotifyWithMail,omitempty"`
	NotifyWithPopup bool `json:"NotifyWithPopup,omitempty"`
	Sound []string `json:"Sound,omitempty"`
}
