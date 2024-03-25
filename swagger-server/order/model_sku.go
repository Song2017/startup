/*
 * Order Service
 *
 * Services that provides the capabilities for orders and Odoo
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package order

type Sku struct {

	Barcode string `json:"barcode,omitempty"`

	SkuNumber string `json:"skuNumber,omitempty"`
}