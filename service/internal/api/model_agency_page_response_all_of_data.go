/*
 * agency base
 *
 * This is my API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type AgencyPageResponseAllOfData struct {

	Total int64 `json:"total,omitempty"`

	Size int64 `json:"size,omitempty"`

	Current int64 `json:"current,omitempty"`

	Pages int64 `json:"pages,omitempty"`

	Records []Agency `json:"records,omitempty"`
}
