/*
 * agency base
 *
 * This is my API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type AgencyAddResponse struct {

	Code int32 `json:"code,omitempty"`

	Message string `json:"message,omitempty"`

	RequestId string `json:"requestId,omitempty"`

	Success bool `json:"success,omitempty"`

	Data Id `json:"data,omitempty"`
}
