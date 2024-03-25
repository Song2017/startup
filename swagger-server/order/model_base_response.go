/*
 * Order Service
 *
 * Services that provides the capabilities for orders and Odoo
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package order

// BaseResponse - Common Response payload
type BaseResponse struct {

	// code 200/500
	Code int32 `json:"code,omitempty"`

	// wrapped data
	Data map[string]interface{} `json:"data,omitempty"`

	// error message
	Msg string `json:"msg,omitempty"`

	TraceId string `json:"traceId,omitempty"`
}
