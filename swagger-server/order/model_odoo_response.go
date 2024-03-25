/*
 * Order Service
 *
 * Services that provides the capabilities for orders and Odoo
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package order

type OdooResponse struct {

	Code int32 `json:"code,omitempty"`

	ConnectorResponses []map[string]interface{} `json:"connectorResponses,omitempty"`

	Data DataPayload `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	ResponseTime int32 `json:"responseTime,omitempty"`
}