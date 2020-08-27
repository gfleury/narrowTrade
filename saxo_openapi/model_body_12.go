/*
 * Swagger Saxo Bank openapi
 *
 * Saxo Bank openapi for trading
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saxo_openapi

type Body12 struct {
	ClientKey   string                                               `json:"ClientKey,omitempty"`
	Documents   []string                                             `json:"Documents,omitempty"`
	RenewalData *Openapicmv1clientrenewalsRenewalEntityIdRenewalData `json:"RenewalData,omitempty"`
}
