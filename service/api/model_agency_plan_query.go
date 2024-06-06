/*
 * agency base
 *
 * This is my API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api



import (
	"time"
)

type AgencyPlanQuery struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	AgencyId int64 `json:"agencyId,omitempty"`

	AgencyName string `json:"agencyName,omitempty"`

	// 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation int32 `json:"cardAssociation,omitempty"`

	PlatformId int64 `json:"platformId,omitempty"`

	PlatformName string `json:"platformName,omitempty"`

	// 1:phone 2:email 3:username
	RegistrationMethod int32 `json:"registrationMethod,omitempty"`

	// in percentage
	Discount float32 `json:"discount,omitempty"`

	ReferralCode string `json:"referralCode,omitempty"`

	// 1:active 2:inactive
	PlanStatus int32 `json:"planStatus,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	CreatedAtFrom time.Time `json:"createdAtFrom,omitempty"`

	CreatedAtTo time.Time `json:"createdAtTo,omitempty"`

	UpdatedAtFrom time.Time `json:"updatedAtFrom,omitempty"`

	UpdatedAtTo time.Time `json:"updatedAtTo,omitempty"`

	IdIn []int64 `json:"idIn,omitempty"`
}