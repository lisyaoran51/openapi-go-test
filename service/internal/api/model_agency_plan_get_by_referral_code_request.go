/*
 * agency base
 *
 * This is my API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type AgencyPlanGetByReferralCodeRequest struct {

	UserId int64 `json:"userId,omitempty"`

	ReferralCode string `json:"referralCode,omitempty"`
}
