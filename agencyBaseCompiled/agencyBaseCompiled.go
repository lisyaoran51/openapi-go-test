package agencyBaseCompiled

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Agency defines model for agency.
type Agency struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int       `json:"agencyStatus,omitempty"`
	CountryCode         *int       `json:"countryCode,omitempty"`
	CreatedAt           *time.Time `json:"createdAt,omitempty"`
	Id                  *int64     `json:"id,omitempty"`
	Name                *string    `json:"name,omitempty"`
	NationalPhoneNumber *string    `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64     `json:"platformId,omitempty"`
	PlatformName        *string    `json:"platformName,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyAddResponse defines model for agencyAddResponse.
type AgencyAddResponse struct {
	Code      *int    `json:"code,omitempty"`
	Data      *Id     `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyBatchUpdateRequest defines model for agencyBatchUpdateRequest.
type AgencyBatchUpdateRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int         `json:"agencyStatus,omitempty"`
	CountryCode         *int         `json:"countryCode,omitempty"`
	CreatedAt           *time.Time   `json:"createdAt,omitempty"`
	Id                  *int64       `json:"id,omitempty"`
	Name                *string      `json:"name,omitempty"`
	NationalPhoneNumber *string      `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64       `json:"platformId,omitempty"`
	PlatformName        *string      `json:"platformName,omitempty"`
	Query               *AgencyQuery `json:"query,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyBatchUpdateResponse defines model for agencyBatchUpdateResponse.
type AgencyBatchUpdateResponse struct {
	Code      *int                    `json:"code,omitempty"`
	Data      *map[string]interface{} `json:"data,omitempty"`
	Message   *string                 `json:"message,omitempty"`
	RequestId *string                 `json:"requestId,omitempty"`
	Success   *bool                   `json:"success,omitempty"`
}

// AgencyGetResponse defines model for agencyGetResponse.
type AgencyGetResponse struct {
	Code      *int      `json:"code,omitempty"`
	Data      *[]Agency `json:"data,omitempty"`
	Message   *string   `json:"message,omitempty"`
	RequestId *string   `json:"requestId,omitempty"`
	Success   *bool     `json:"success,omitempty"`
}

// AgencyPageRequest defines model for agencyPageRequest.
type AgencyPageRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int       `json:"agencyStatus,omitempty"`
	CountryCode         *int       `json:"countryCode,omitempty"`
	CreatedAt           *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom       *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo         *time.Time `json:"createdAtTo,omitempty"`
	Id                  *int64     `json:"id,omitempty"`
	IdIn                *[]int64   `json:"idIn,omitempty"`
	Name                *string    `json:"name,omitempty"`
	NationalPhoneNumber *string    `json:"nationalPhoneNumber,omitempty"`
	PageCurrent         *int       `json:"pageCurrent,omitempty"`
	PageSize            *int       `json:"pageSize,omitempty"`
	PlatformId          *int64     `json:"platformId,omitempty"`
	PlatformName        *string    `json:"platformName,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAtFrom      *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo        *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPageResponse defines model for agencyPageResponse.
type AgencyPageResponse struct {
	Code *int `json:"code,omitempty"`
	Data *struct {
		Current *int64    `json:"current,omitempty"`
		Pages   *int64    `json:"pages,omitempty"`
		Records *[]Agency `json:"records,omitempty"`
		Size    *int64    `json:"size,omitempty"`
		Total   *int64    `json:"total,omitempty"`
	} `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyPlan defines model for agencyPlan.
type AgencyPlan struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanAddResponse defines model for agencyPlanAddResponse.
type AgencyPlanAddResponse struct {
	Code      *int    `json:"code,omitempty"`
	Data      *Id     `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyPlanBatchUpdateRequest defines model for agencyPlanBatchUpdateRequest.
type AgencyPlanBatchUpdateRequest struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int             `json:"planStatus,omitempty"`
	PlatformId   *int64           `json:"platformId,omitempty"`
	PlatformName *string          `json:"platformName,omitempty"`
	Query        *AgencyPlanQuery `json:"query,omitempty"`
	ReferralCode *string          `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanBatchUpdateResponse defines model for agencyPlanBatchUpdateResponse.
type AgencyPlanBatchUpdateResponse struct {
	Code      *int                    `json:"code,omitempty"`
	Data      *map[string]interface{} `json:"data,omitempty"`
	Message   *string                 `json:"message,omitempty"`
	RequestId *string                 `json:"requestId,omitempty"`
	Success   *bool                   `json:"success,omitempty"`
}

// AgencyPlanGetByReferralCodeRequest defines model for agencyPlanGetByReferralCodeRequest.
type AgencyPlanGetByReferralCodeRequest struct {
	ReferralCode *string `json:"referralCode,omitempty"`
	UserId       *int64  `json:"userId,omitempty"`
}

// AgencyPlanGetByReferralCodeResponse defines model for agencyPlanGetByReferralCodeResponse.
type AgencyPlanGetByReferralCodeResponse struct {
	Code      *int        `json:"code,omitempty"`
	Data      *AgencyPlan `json:"data,omitempty"`
	Message   *string     `json:"message,omitempty"`
	RequestId *string     `json:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty"`
}

// AgencyPlanGetResponse defines model for agencyPlanGetResponse.
type AgencyPlanGetResponse struct {
	Code      *int          `json:"code,omitempty"`
	Data      *[]AgencyPlan `json:"data,omitempty"`
	Message   *string       `json:"message,omitempty"`
	RequestId *string       `json:"requestId,omitempty"`
	Success   *bool         `json:"success,omitempty"`
}

// AgencyPlanQuery defines model for agencyPlanQuery.
type AgencyPlanQuery struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom   *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo     *time.Time `json:"createdAtTo,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	IdIn     *[]int64 `json:"idIn,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAtFrom      *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo        *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanUpdateRequest defines model for agencyPlanUpdateRequest.
type AgencyPlanUpdateRequest struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int             `json:"planStatus,omitempty"`
	PlatformId   *int64           `json:"platformId,omitempty"`
	PlatformName *string          `json:"platformName,omitempty"`
	Query        *AgencyPlanQuery `json:"query,omitempty"`
	ReferralCode *string          `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanUpdateResponse defines model for agencyPlanUpdateResponse.
type AgencyPlanUpdateResponse struct {
	Code      *int        `json:"code,omitempty"`
	Data      *AgencyPlan `json:"data,omitempty"`
	Message   *string     `json:"message,omitempty"`
	RequestId *string     `json:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty"`
}

// AgencyQuery defines model for agencyQuery.
type AgencyQuery struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int       `json:"agencyStatus,omitempty"`
	CountryCode         *int       `json:"countryCode,omitempty"`
	CreatedAt           *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom       *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo         *time.Time `json:"createdAtTo,omitempty"`
	Id                  *int64     `json:"id,omitempty"`
	IdIn                *[]int64   `json:"idIn,omitempty"`
	Name                *string    `json:"name,omitempty"`
	NationalPhoneNumber *string    `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64     `json:"platformId,omitempty"`
	PlatformName        *string    `json:"platformName,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAtFrom      *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo        *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyUpdateRequest defines model for agencyUpdateRequest.
type AgencyUpdateRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int         `json:"agencyStatus,omitempty"`
	CountryCode         *int         `json:"countryCode,omitempty"`
	CreatedAt           *time.Time   `json:"createdAt,omitempty"`
	Id                  *int64       `json:"id,omitempty"`
	Name                *string      `json:"name,omitempty"`
	NationalPhoneNumber *string      `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64       `json:"platformId,omitempty"`
	PlatformName        *string      `json:"platformName,omitempty"`
	Query               *AgencyQuery `json:"query,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyUpdateResponse defines model for agencyUpdateResponse.
type AgencyUpdateResponse struct {
	Code      *int    `json:"code,omitempty"`
	Data      *Agency `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyWithPlan defines model for agencyWithPlan.
type AgencyWithPlan struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus *int `json:"agencyStatus,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CountryCode     *int       `json:"countryCode,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount            *float32 `json:"discount,omitempty"`
	Id                  *int64   `json:"id,omitempty"`
	Name                *string  `json:"name,omitempty"`
	NationalPhoneNumber *string  `json:"nationalPhoneNumber,omitempty"`
	PlanId              *int64   `json:"planId,omitempty"`
	PlanName            *string  `json:"planName,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyWithPlanAddResponse defines model for agencyWithPlanAddResponse.
type AgencyWithPlanAddResponse struct {
	Code *int `json:"code,omitempty"`
	Data *struct {
		AgencyId     *int64 `json:"agencyId,omitempty"`
		AgencyPlanId *int64 `json:"agencyPlanId,omitempty"`
	} `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyWithPlanGetResponse defines model for agencyWithPlanGetResponse.
type AgencyWithPlanGetResponse struct {
	Code      *int              `json:"code,omitempty"`
	Data      *[]AgencyWithPlan `json:"data,omitempty"`
	Message   *string           `json:"message,omitempty"`
	RequestId *string           `json:"requestId,omitempty"`
	Success   *bool             `json:"success,omitempty"`
}

// AgencyWithPlanPageRequest defines model for agencyWithPlanPageRequest.
type AgencyWithPlanPageRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus *int `json:"agencyStatus,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CountryCode     *int       `json:"countryCode,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom   *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo     *time.Time `json:"createdAtTo,omitempty"`

	// Discount in percentage
	Discount            *float32 `json:"discount,omitempty"`
	Id                  *int64   `json:"id,omitempty"`
	Name                *string  `json:"name,omitempty"`
	NationalPhoneNumber *string  `json:"nationalPhoneNumber,omitempty"`
	OrderByColumn       *string  `json:"orderByColumn,omitempty"`

	// OrderByDirection 1:asc 2:desc
	OrderByDirection *int    `json:"orderByDirection,omitempty"`
	PageCurrent      *int    `json:"pageCurrent,omitempty"`
	PageSize         *int    `json:"pageSize,omitempty"`
	PlanId           *int64  `json:"planId,omitempty"`
	PlanName         *string `json:"planName,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int `json:"registrationMethod,omitempty"`

	// Strategy 1:newest 2:oldest 3:newest active 4:oldest active
	Strategy      *int       `json:"strategy,omitempty"`
	UpdatedAtFrom *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo   *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

// AgencyWithPlanPageResponse defines model for agencyWithPlanPageResponse.
type AgencyWithPlanPageResponse struct {
	Code *int `json:"code,omitempty"`
	Data *struct {
		Current *int64            `json:"current,omitempty"`
		Pages   *int64            `json:"pages,omitempty"`
		Records *[]AgencyWithPlan `json:"records,omitempty"`
		Size    *int64            `json:"size,omitempty"`
		Total   *int64            `json:"total,omitempty"`
	} `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyWithPlanQuery defines model for agencyWithPlanQuery.
type AgencyWithPlanQuery struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus *int `json:"agencyStatus,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CountryCode     *int       `json:"countryCode,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom   *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo     *time.Time `json:"createdAtTo,omitempty"`

	// Discount in percentage
	Discount            *float32 `json:"discount,omitempty"`
	Id                  *int64   `json:"id,omitempty"`
	Name                *string  `json:"name,omitempty"`
	NationalPhoneNumber *string  `json:"nationalPhoneNumber,omitempty"`
	OrderByColumn       *string  `json:"orderByColumn,omitempty"`

	// OrderByDirection 1:asc 2:desc
	OrderByDirection *int    `json:"orderByDirection,omitempty"`
	PlanId           *int64  `json:"planId,omitempty"`
	PlanName         *string `json:"planName,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int `json:"registrationMethod,omitempty"`

	// Strategy 1:newest 2:oldest 3:newest active 4:oldest active
	Strategy      *int       `json:"strategy,omitempty"`
	UpdatedAtFrom *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo   *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

// Id defines model for id.
type Id struct {
	Id *int64 `json:"id,omitempty"`
}

// Platform defines model for platform.
type Platform struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *int64     `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// PlatformGetResponse defines model for platformGetResponse.
type PlatformGetResponse struct {
	Code      *int        `json:"code,omitempty"`
	Data      *[]Platform `json:"data,omitempty"`
	Message   *string     `json:"message,omitempty"`
	RequestId *string     `json:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty"`
}

// Result defines model for result.
type Result struct {
	Code      *int    `json:"code,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencybatchUpdateJSONRequestBody defines body for AgencybatchUpdate for application/json ContentType.
type AgencybatchUpdateJSONRequestBody = AgencyBatchUpdateRequest

// AgencyaddJSONRequestBody defines body for Agencyadd for application/json ContentType.
type AgencyaddJSONRequestBody = Agency

// AgencypagegetJSONRequestBody defines body for Agencypageget for application/json ContentType.
type AgencypagegetJSONRequestBody = AgencyPageRequest

// AgencyupdateJSONRequestBody defines body for Agencyupdate for application/json ContentType.
type AgencyupdateJSONRequestBody = AgencyUpdateRequest

// AgencygetJSONRequestBody defines body for Agencyget for application/json ContentType.
type AgencygetJSONRequestBody = AgencyQuery

// AgencyPlanbatchUpdateJSONRequestBody defines body for AgencyPlanbatchUpdate for application/json ContentType.
type AgencyPlanbatchUpdateJSONRequestBody = AgencyPlanBatchUpdateRequest

// AgencyPlanaddJSONRequestBody defines body for AgencyPlanadd for application/json ContentType.
type AgencyPlanaddJSONRequestBody = AgencyPlan

// AgencyPlangetByReferralCodeJSONRequestBody defines body for AgencyPlangetByReferralCode for application/json ContentType.
type AgencyPlangetByReferralCodeJSONRequestBody = AgencyPlanGetByReferralCodeRequest

// AgencyPlanupdateJSONRequestBody defines body for AgencyPlanupdate for application/json ContentType.
type AgencyPlanupdateJSONRequestBody = AgencyPlanUpdateRequest

// AgencyPlangetJSONRequestBody defines body for AgencyPlanget for application/json ContentType.
type AgencyPlangetJSONRequestBody = AgencyPlanQuery

// AgencyWithPlanaddJSONRequestBody defines body for AgencyWithPlanadd for application/json ContentType.
type AgencyWithPlanaddJSONRequestBody = AgencyWithPlan

// AgencyWithPlanpagegetJSONRequestBody defines body for AgencyWithPlanpageget for application/json ContentType.
type AgencyWithPlanpagegetJSONRequestBody = AgencyWithPlanPageRequest

// AgencyWithPlangetJSONRequestBody defines body for AgencyWithPlanget for application/json ContentType.
type AgencyWithPlangetJSONRequestBody = AgencyWithPlanQuery

// PlatformgetJSONRequestBody defines body for Platformget for application/json ContentType.
type PlatformgetJSONRequestBody = Platform

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// AgencybatchUpdateWithBody request with any body
	AgencybatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencybatchUpdate(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyaddWithBody request with any body
	AgencyaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencyadd(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencypagegetWithBody request with any body
	AgencypagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencypageget(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyupdateWithBody request with any body
	AgencyupdateWithBody(ctx context.Context, agencyId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencyupdate(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencygetWithBody request with any body
	AgencygetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencyget(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlanbatchUpdateWithBody request with any body
	AgencyPlanbatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanbatchUpdate(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlanaddWithBody request with any body
	AgencyPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanadd(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlangetByReferralCodeWithBody request with any body
	AgencyPlangetByReferralCodeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlangetByReferralCode(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlanupdateWithBody request with any body
	AgencyPlanupdateWithBody(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanupdate(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlangetWithBody request with any body
	AgencyPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanget(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyWithPlanaddWithBody request with any body
	AgencyWithPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyWithPlanadd(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyWithPlanpagegetWithBody request with any body
	AgencyWithPlanpagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyWithPlanpageget(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyWithPlangetWithBody request with any body
	AgencyWithPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyWithPlanget(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PlatformgetWithBody request with any body
	PlatformgetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Platformget(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AgencybatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencybatchUpdateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencybatchUpdate(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencybatchUpdateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyaddRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencyadd(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyaddRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencypagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencypagegetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencypageget(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencypagegetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyupdateWithBody(ctx context.Context, agencyId int64, cont
	operation="generate"

[Info  - 12:04:50 PM] 2024/06/06 12:04:50 entType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyupdateRequestWithBody(c.Server, agencyId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencyupdate(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyupdateRequest(c.Server, agencyId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencygetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencygetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencyget(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencygetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanbatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanbatchUpdateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanbatchUpdate(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanbatchUpdateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanaddRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanadd(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanaddRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlangetByReferralCodeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetByReferralCodeRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlangetByReferralCode(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetByReferralCodeRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanupdateWithBody(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanupdateRequestWithBody(c.Server, agencyPlanId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanupdate(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanupdateRequest(c.Server, agencyPlanId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanget(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanaddRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanadd(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanaddRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanpagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanpagegetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanpageget(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanpagegetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlangetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanget(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlangetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PlatformgetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlatformgetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Platformget(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlatformgetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAgencybatchUpdateRequest calls the generic AgencybatchUpdate builder with application/json body
func NewAgencybatchUpdateRequest(server string, body AgencybatchUpdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencybatchUpdateRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencybatchUpdateRequestWithBody generates requests for AgencybatchUpdate with any type of body
func NewAgencybatchUpdateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyaddRequest calls the generic Agencyadd builder with application/json body
func NewAgencyaddRequest(server string, body AgencyaddJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyaddRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyaddRequestWithBody generates requests for Agencyadd with any type of body
func NewAgencyaddRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencypagegetRequest calls the generic Agencypageget builder with application/json body
func NewAgencypagegetRequest(server string, body AgencypagegetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencypagegetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencypagegetRequestWithBody generates requests for Agencypageget with any type of body
func NewAgencypagegetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency/pageGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyupdateRequest calls the generic Agencyupdate builder with application/json body
func NewAgencyupdateRequest(server string, agencyId int64, body AgencyupdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyupdateRequestWithBody(server, agencyId, "application/json", bodyReader)
}

// NewAgencyupdateRequestWithBody generates requests for Agencyupdate with any type of body
func NewAgencyupdateRequestWithBody(server string, agencyId int64, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "agencyId", runtime.ParamLocationPath, agencyId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencygetRequest calls the generic Agencyget builder with application/json body
func NewAgencygetRequest(server string, body AgencygetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencygetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencygetRequestWithBody generates requests for Agencyget with any type of body
func NewAgencygetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlanbatchUpdateRequest calls the generic AgencyPlanbatchUpdate builder with application/json body
func NewAgencyPlanbatchUpdateRequest(server string, body AgencyPlanbatchUpdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlanbatchUpdateRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlanbatchUpdateRequestWithBody generates requests for AgencyPlanbatchUpdate with any type of body
func NewAgencyPlanbatchUpdateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlanaddRequest calls the generic AgencyPlanadd builder with application/json body
func NewAgencyPlanaddRequest(server string, body AgencyPlanaddJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlanaddRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlanaddRequestWithBody generates requests for AgencyPlanadd with any type of body
func NewAgencyPlanaddRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlangetByReferralCodeRequest calls the generic AgencyPlangetByReferralCode builder with application/json body
func NewAgencyPlangetByReferralCodeRequest(server string, body AgencyPlangetByReferralCodeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlangetByReferralCodeRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlangetByReferralCodeRequestWithBody generates requests for AgencyPlangetByReferralCode with any type of body
func NewAgencyPlangetByReferralCodeRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan/byReferralCodeGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlanupdateRequest calls the generic AgencyPlanupdate builder with application/json body
func NewAgencyPlanupdateRequest(server string, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlanupdateRequestWithBody(server, agencyPlanId, "application/json", bodyReader)
}

// NewAgencyPlanupdateRequestWithBody generates requests for AgencyPlanupdate with any type of body
func NewAgencyPlanupdateRequestWithBody(server string, agencyPlanId int64, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "agencyPlanId", runtime.ParamLocationPath, agencyPlanId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlangetRequest calls the generic AgencyPlanget builder with application/json body
func NewAgencyPlangetRequest(server string, body AgencyPlangetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlangetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlangetRequestWithBody generates requests for AgencyPlanget with any type of body
func NewAgencyPlangetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlanGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyWithPlanaddRequest calls the generic AgencyWithPlanadd builder with application/json body
func NewAgencyWithPlanaddRequest(server string, body AgencyWithPlanaddJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyWithPlanaddRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyWithPlanaddRequestWithBody generates requests for AgencyWithPlanadd with any type of body
func NewAgencyWithPlanaddRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyWithPlan")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyWithPlanpagegetRequest calls the generic AgencyWithPlanpageget builder with application/json body
func NewAgencyWithPlanpagegetRequest(server string, body AgencyWithPlanpagegetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyWithPlanpagegetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyWithPlanpagegetRequestWithBody generates requests for AgencyWithPlanpageget with any type of body
func NewAgencyWithPlanpagegetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyWithPlan/pageGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyWithPlangetRequest calls the generic AgencyWithPlanget builder with application/json body
func NewAgencyWithPlangetRequest(server string, body AgencyWithPlangetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyWithPlangetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyWithPlangetRequestWithBody generates requests for AgencyWithPlanget with any type of body
func NewAgencyWithPlangetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyWithPlanGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPlatformgetRequest calls the generic Platformget builder with application/json body
func NewPlatformgetRequest(server string, body PlatformgetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPlatformgetRequestWithBody(server, "application/json", bodyReader)
}

// NewPlatformgetRequestWithBody generates requests for Platformget with any type of body
func NewPlatformgetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/platformGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// AgencybatchUpdateWithBodyWithResponse request with any body
	AgencybatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error)

	AgencybatchUpdateWithResponse(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error)

	// AgencyaddWithBodyWithResponse request with any body
	AgencyaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error)

	AgencyaddWithResponse(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error)

	// AgencypagegetWithBodyWithResponse request with any body
	AgencypagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error)

	AgencypagegetWithResponse(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error)

	// AgencyupdateWithBodyWithResponse request with any body
	AgencyupdateWithBodyWithResponse(ctx context.Context, agencyId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error)

	AgencyupdateWithResponse(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error)

	// AgencygetWithBodyWithResponse request with any body
	AgencygetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencygetResponse, error)

	AgencygetWithResponse(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencygetResponse, error)

	// AgencyPlanbatchUpdateWithBodyWithResponse request with any body
	AgencyPlanbatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error)

	AgencyPlanbatchUpdateWithResponse(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error)

	// AgencyPlanaddWithBodyWithResponse request with any body
	AgencyPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error)

	AgencyPlanaddWithResponse(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error)

	// AgencyPlangetByReferralCodeWithBodyWithResponse request with any body
	AgencyPlangetByReferralCodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error)

	AgencyPlangetByReferralCodeWithResponse(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error)

	// AgencyPlanupdateWithBodyWithResponse request with any body
	AgencyPlanupdateWithBodyWithResponse(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error)

	AgencyPlanupdateWithResponse(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error)

	// AgencyPlangetWithBodyWithResponse request with any body
	AgencyPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error)

	AgencyPlangetWithResponse(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error)

	// AgencyWithPlanaddWithBodyWithResponse request with any body
	AgencyWithPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error)

	AgencyWithPlanaddWithResponse(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error)

	// AgencyWithPlanpagegetWithBodyWithResponse request with any body
	AgencyWithPlanpagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error)

	AgencyWithPlanpagegetWithResponse(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error)

	// AgencyWithPlangetWithBodyWithResponse request with any body
	AgencyWithPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error)

	AgencyWithPlangetWithResponse(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error)

	// PlatformgetWithBodyWithResponse request with any body
	PlatformgetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error)

	PlatformgetWithResponse(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error)
}

type AgencybatchUpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyBatchUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencybatchUpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencybatchUpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyaddResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyAddResponse
}

// Status returns HTTPResponse.Status
func (r AgencyaddResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyaddResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencypagegetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPageResponse
}

// Status returns HTTPResponse.Status
func (r AgencypagegetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencypagegetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyupdateResponse struct {
	Body         []byte
	
	operation="generate"

[Info  - 12:04:50 PM] 2024/06/06 12:04:50 HTTPResponse *http.Response
	JSON200      *AgencyUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencyupdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyupdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencygetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyGetResponse
}

// Status returns HTTPResponse.Status
func (r AgencygetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencygetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlanbatchUpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanBatchUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlanbatchUpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlanbatchUpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlanaddResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanAddResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlanaddResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlanaddResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlangetByReferralCodeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanGetByReferralCodeResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlangetByReferralCodeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlangetByReferralCodeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlanupdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlanupdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlanupdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlangetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanGetResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlangetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlangetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyWithPlanaddResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyWithPlanAddResponse
}

// Status returns HTTPResponse.Status
func (r AgencyWithPlanaddResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyWithPlanaddResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyWithPlanpagegetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyWithPlanPageResponse
}

// Status returns HTTPResponse.Status
func (r AgencyWithPlanpagegetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyWithPlanpagegetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyWithPlangetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyWithPlanGetResponse
}

// Status returns HTTPResponse.Status
func (r AgencyWithPlangetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyWithPlangetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PlatformgetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PlatformGetResponse
}

// Status returns HTTPResponse.Status
func (r PlatformgetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PlatformgetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AgencybatchUpdateWithBodyWithResponse request with arbitrary body returning *AgencybatchUpdateResponse
func (c *ClientWithResponses) AgencybatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error) {
	rsp, err := c.AgencybatchUpdateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencybatchUpdateResponse(rsp)
}

func (c *ClientWithResponses) AgencybatchUpdateWithResponse(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error) {
	rsp, err := c.AgencybatchUpdate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencybatchUpdateResponse(rsp)
}

// AgencyaddWithBodyWithResponse request with arbitrary body returning *AgencyaddResponse
func (c *ClientWithResponses) AgencyaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error) {
	rsp, err := c.AgencyaddWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyaddResponse(rsp)
}

func (c *ClientWithResponses) AgencyaddWithResponse(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error) {
	rsp, err := c.Agencyadd(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyaddResponse(rsp)
}

// AgencypagegetWithBodyWithResponse request with arbitrary body returning *AgencypagegetResponse
func (c *ClientWithResponses) AgencypagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error) {
	rsp, err := c.AgencypagegetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencypagegetResponse(rsp)
}

func (c *ClientWithResponses) AgencypagegetWithResponse(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error) {
	rsp, err := c.Agencypageget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencypagegetResponse(rsp)
}

// AgencyupdateWithBodyWithResponse request with arbitrary body returning *AgencyupdateResponse
func (c *ClientWithResponses) AgencyupdateWithBodyWithResponse(ctx context.Context, agencyId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error) {
	rsp, err := c.AgencyupdateWithBody(ctx, agencyId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyupdateResponse(rsp)
}

func (c *ClientWithResponses) AgencyupdateWithResponse(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error) {
	rsp, err := c.Agencyupdate(ctx, agencyId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyupdateResponse(rsp)
}

// AgencygetWithBodyWithResponse request with arbitrary body returning *AgencygetResponse
func (c *ClientWithResponses) AgencygetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencygetResponse, error) {
	rsp, err := c.AgencygetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencygetResponse(rsp)
}

func (c *ClientWithResponses) AgencygetWithResponse(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencygetResponse, error) {
	rsp, err := c.Agencyget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencygetResponse(rsp)
}

// AgencyPlanbatchUpdateWithBodyWithResponse request with arbitrary body returning *AgencyPlanbatchUpdateResponse
func (c *ClientWithResponses) AgencyPlanbatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error) {
	rsp, err := c.AgencyPlanbatchUpdateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanbatchUpdateResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlanbatchUpdateWithResponse(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error) {
	rsp, err := c.AgencyPlanbatchUpdate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanbatchUpdateResponse(rsp)
}

// AgencyPlanaddWithBodyWithResponse request with arbitrary body returning *AgencyPlanaddResponse
func (c *ClientWithResponses) AgencyPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error) {
	rsp, err := c.AgencyPlanaddWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanaddResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlanaddWithResponse(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error) {
	rsp, err := c.AgencyPlanadd(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanaddResponse(rsp)
}

// AgencyPlangetByReferralCodeWithBodyWithResponse request with arbitrary body returning *AgencyPlangetByReferralCodeResponse
func (c *ClientWithResponses) AgencyPlangetByReferralCodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error) {
	rsp, err := c.AgencyPlangetByReferralCodeWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetByReferralCodeResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlangetByReferralCodeWithResponse(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error) {
	rsp, err := c.AgencyPlangetByReferralCode(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetByReferralCodeResponse(rsp)
}

// AgencyPlanupdateWithBodyWithResponse request with arbitrary body returning *AgencyPlanupdateResponse
func (c *ClientWithResponses) AgencyPlanupdateWithBodyWithResponse(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error) {
	rsp, err := c.AgencyPlanupdateWithBody(ctx, agencyPlanId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanupdateResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlanupdateWithResponse(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error) {
	rsp, err := c.AgencyPlanupdate(ctx, agencyPlanId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanupdateResponse(rsp)
}

// AgencyPlangetWithBodyWithResponse request with arbitrary body returning *AgencyPlangetResponse
func (c *ClientWithResponses) AgencyPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error) {
	rsp, err := c.AgencyPlangetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlangetWithResponse(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error) {
	rsp, err := c.AgencyPlanget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetResponse(rsp)
}

// AgencyWithPlanaddWithBodyWithResponse request with arbitrary body returning *AgencyWithPlanaddResponse
func (c *ClientWithResponses) AgencyWithPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error) {
	rsp, err := c.AgencyWithPlanaddWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanaddResponse(rsp)
}

func (c *ClientWithResponses) AgencyWithPlanaddWithResponse(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error) {
	rsp, err := c.AgencyWithPlanadd(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanaddResponse(rsp)
}

// AgencyWithPlanpagegetWithBodyWithResponse request with arbitrary body returning *AgencyWithPlanpagegetResponse
func (c *ClientWithResponses) AgencyWithPlanpagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error) {
	rsp, err := c.AgencyWithPlanpagegetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanpagegetResponse(rsp)
}

func (c *ClientWithResponses) AgencyWithPlanpagegetWithResponse(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error) {
	rsp, err := c.AgencyWithPlanpageget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanpagegetResponse(rsp)
}

// AgencyWithPlangetWithBodyWithResponse request with arbitrary body returning *AgencyWithPlangetResponse
func (c *ClientWithResponses) AgencyWithPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error) {
	rsp, err := c.AgencyWithPlangetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlangetResponse(rsp)
}

func (c *ClientWithResponses) AgencyWithPlangetWithResponse(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error) {
	rsp, err := c.AgencyWithPlanget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlangetResponse(rsp)
}

// PlatformgetWithBodyWithResponse request with arbitrary body returning *PlatformgetResponse
func (c *ClientWithResponses) PlatformgetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error) {
	rsp, err := c.PlatformgetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlatformgetResponse(rsp)
}

func (c *ClientWithResponses) PlatformgetWithResponse(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error) {
	rsp, err := c.Platformget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlatformgetResponse(rsp)
}

// ParseAgencybatchUpdateResponse parses an HTTP response from a AgencybatchUpdateWithResponse call
func ParseAgencybatchUpdateResponse(rsp *http.Response) (*AgencybatchUpdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencybatchUpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyBatchUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyaddResponse parses an HTTP response from a AgencyaddWithResponse call
func ParseAgencyaddResponse(rsp *http.Response) (*AgencyaddResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyaddResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyAddResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencypagegetResponse parses an HTTP response from a AgencypagegetWithResponse call
func ParseAgencypagegetResponse(rsp *http.Response) (*AgencypagegetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencypagegetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyupdateResponse parses an HTTP response from a AgencyupdateWithResponse call
func ParseAgencyupdateResponse(rsp *http.Response) (*AgencyupdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyupdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencygetResponse parses an HTTP response from a AgencygetWithResponse call
func ParseAgencygetResponse(rsp *http.Response) (*AgencygetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencygetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlanbatchUpdateResponse parses an HTTP response from a AgencyPlanbatchUpdateWithResponse call
func ParseAgencyPlanbatchUpdateResponse(rsp *http.Response) (*AgencyPlanbatchUpdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlanbatchUpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanBatchUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlanaddResponse parses an HTTP response from a AgencyPlanaddWithResponse call
func ParseAgencyPlanaddResponse(rsp *http.Response) (*AgencyPlanaddResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlanaddResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanAddResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlangetByReferralCodeResponse parses an HTTP response from a AgencyPlangetByReferralCodeWithResponse call
func ParseAgencyPlangetByReferralCodeResponse(rsp *http.Response) (*AgencyPlangetByReferralCodeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlangetByReferralCodeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanGetByReferralCodeResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlanupdateResponse parses an HTTP response from a AgencyPlanupdateWithResponse call
func ParseAgencyPlanupdateResponse(rsp *http.Response) (*AgencyPlanupdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlanupdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlangetResponse parses an HTTP response from a AgencyPlangetWithResponse call
func ParseAgencyPlangetResponse(rsp *http.Response) (*AgencyPlangetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlangetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyWithPlanaddResponse parses an HTTP response from a AgencyWithPlanaddWithResponse call
func ParseAgencyWithPlanaddResponse(rsp *http.Response) (*AgencyWithPlanaddResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyWithPlanaddResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyWithPlanAddResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyWithPlanpagegetResponse parses an HTTP response from a AgencyWithPlanpagegetWithResponse call
func ParseAgencyWithPlanpagegetResponse(rsp *http.Response) (*AgencyWithPlanpagegetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyWithPlanpagegetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyWithPlanPageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyWithPlangetResponse parses an HTTP response from a AgencyWithPlangetWithResponse call
func ParseAgencyWithPlangetResponse(rsp *http.Response) (*AgencyWithPlangetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyWithPlangetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyWithPlanGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePlatformgetResponse parses an HTTP response from a PlatformgetWithResponse call
func ParsePlatformgetResponse(rsp *http.Response) (*PlatformgetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PlatformgetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PlatformGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (PATCH /agency)
	AgencybatchUpdate(ctx echo.Context) error

	// (POST /agency)
	Agencyadd(ctx echo.Context) error

	// (POST /agency/pageGet)
	Agencypageget(ctx echo.Context) error

	// (PATCH /agency/{agencyId})
	Agencyupdate(ctx echo.Context, agencyId int64) error

	// (POST /agencyGet)
	Agencyget(ctx echo.Context) error

	// (PATCH /agencyPlan)
	AgencyPlanbatchUpdate(ctx echo.Context) error

	// (POST /agencyPlan)
	AgencyPlanadd(ctx echo.Context) error

	// (POST /agencyPlan/byReferralCodeGet)
	AgencyPlangetByReferralCode(ctx echo.Context) error

	// (PATCH /agencyPlan/{agencyPlanId})
	AgencyPlanupdate(ctx echo.Context, agencyPlanId int64) error

	// (POST /agencyPlanGet)
	AgencyPlanget(ctx echo.Context) error

	// (POST /agencyWithPlan)
	AgencyWithPlanadd(ctx echo.Context) error

	// (POST /agencyWithPlan/pageGet)
	AgencyWithPlanpageget(ctx echo.Context) error

	// (POST /agencyWithPlanGet)
	AgencyWithPlanget(ctx echo.Context) error

	// (POST /platformGet)
	Platformget(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AgencybatchUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) AgencybatchUpdate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencybatchUpdate(ctx)
	return err
}

// Agencyadd converts echo context to params.
func (w *ServerInterfaceWrapper) Agencyadd(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencyadd(ctx)
	return err
}

// Agencypageget converts echo context to params.
func (w *ServerInterfaceWrapper) Agencypageget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencypageget(ctx)
	return err
}

// Agencyupdate converts echo context to params.
func (w *ServerInterfaceWrapper) Agencyupdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "agencyId" -------------
	var agencyId int64

	err = runtime.BindStyledParameterWithOptions("simple", "agencyId", ctx.Param("agencyId"), &agencyId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter agencyId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencyupdate(ctx, agencyId)
	return err
}

// Agencyget converts echo context to params.
func (w *ServerInterfaceWrapper) Agencyget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencyget(ctx)
	return err
}

// AgencyPlanbatchUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanbatchUpdate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanbatchUpdate(ctx)
	return err
}

// AgencyPlanadd converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanadd(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanadd(ctx)
	return err
}

// AgencyPlangetByReferralCode converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlangetByReferralCode(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlangetByReferralCode(ctx)
	return err
}

// AgencyPlanupdate converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanupdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "agencyPlanId" -------------
	var agencyPlanId int64

	err = runtime.BindStyledParameterWithOptions("simple", "agencyPlanId", ctx.Param("agencyPlanId"), &agencyPlanId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter agencyPlanId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanupdate(ctx, agencyPlanId)
	return err
}

// AgencyPlanget converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanget(ctx)
	return err
}

// AgencyWithPlanadd converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyWithPlanadd(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyWithPlanadd(ctx)
	return err
}

// AgencyWithPlanpageget converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyWithPlanpageget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyWithPlanpageget(ctx)
	return err
}

// AgencyWithPlanget converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyWithPlanget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyWithPlanget(ctx)
	return err
}

// Platformget converts echo context to params.
func (w *ServerInterfaceWrapper) Platformget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Platformget(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandler
	operation="generate"

[Info  - 12:04:50 PM] 2024/06/06 12:04:50 sWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.PATCH(baseURL+"/agency", wrapper.AgencybatchUpdate)
	router.POST(baseURL+"/agency", wrapper.Agencyadd)
	router.POST(baseURL+"/agency/pageGet", wrapper.Agencypageget)
	router.PATCH(baseURL+"/agency/:agencyId", wrapper.Agencyupdate)
	router.POST(baseURL+"/agencyGet", wrapper.Agencyget)
	router.PATCH(baseURL+"/agencyPlan", wrapper.AgencyPlanbatchUpdate)
	router.POST(baseURL+"/agencyPlan", wrapper.AgencyPlanadd)
	router.POST(baseURL+"/agencyPlan/byReferralCodeGet", wrapper.AgencyPlangetByReferralCode)
	router.PATCH(baseURL+"/agencyPlan/:agencyPlanId", wrapper.AgencyPlanupdate)
	router.POST(baseURL+"/agencyPlanGet", wrapper.AgencyPlanget)
	router.POST(baseURL+"/agencyWithPlan", wrapper.AgencyWithPlanadd)
	router.POST(baseURL+"/agencyWithPlan/pageGet", wrapper.AgencyWithPlanpageget)
	router.POST(baseURL+"/agencyWithPlanGet", wrapper.AgencyWithPlanget)
	router.POST(baseURL+"/platformGet", wrapper.Platformget)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xaT4/bRBT/KtbAAaTdZNktHCwhsbuFakG0S7sVh6pCE8/bZCp7xp0Zl4ZVDnAoRUK0",
	"B7QIISEoEuqxElwqEJ+GpPRboBnbiRPb8SRZx1vKqel4/rz5/d68v3uCPB6EnAFTErknSHo9CLD5ibvA",
	"vL7+FQoeglAUMuPXFFaR+T8B6QkaKsoZctEbLvYUvQPOtktZ/BNtINUPAbmIMgVdEGiwgTweMSX6+5yA",
	"3qRgggCsgOwq/fmYiwAr5CKCFWwqGmQ2lUpQ1tVLKJmaS5l660L+8A10d/OYgk824a4SeBMzxhVOxH/H",
	"40Grg3lACY9aQb+DFZWhH8nWZFrrCHd8OCCv6Y2dtx2bJQfkqB9Ca/f60ZXXtagMB9l7T+7AzHzsH/Y4",
	"g8tR0AFROC/0sdI3PbC6c2bB5bKTBXSpVMKc/yGoHidF7IZaLmfbhQBT39lxIwnCXKbozCjUfJFPsDWJ",
	"g/EI79wCT+ldYoXbJeQqyJAzacTHvn/lGLk3TtCrAo6Ri15pTzS5nahxW4CMfL3JrBITrLD+d95iSork",
	"uTmWaA8rr3fdXPEq3I5AKnvBkseVF+x2BKJfJVm8+iMz1V7EMwdv5th5glwCVYMAVEEg7cBCE+mwELgC",
	"t0PcXZbThJWczCHuwn4kBDBVbPD0hGv0s0JzWC3rmWM7PepNJLexNbgbL7OYK8DjgsiV6dxAMgHP4lDF",
	"Ffat5uaRn8uFj1mZx7S21PH0UjvtYUF2peQeTZxW3kjfoRI7226ApQKxjwVxdtzdAAT1MHv3bihASueC",
	"+/7+nvOmGzHK2SHuF3vpxZ0wodL49rxclDkhCA+Ywt3MUhY7uf+I/w59zFaNjerw7ccgBPZnwq0Xxvnr",
	"d3X+AgAt1epBgLEZKwUCegeLYCAnbpMBgRbmEqi9/tWMamYQnD6gUn+1Jlo+l/laViDTWnUuqxMW+DUb",
	"VCWqu0hgNVbVs3klYwf1nuCBvZMaLzvii6SXB2wKIJtAYyZGSYziovKOl9nLW8XCi2OzajNXZ/EOl1Lm",
	"/xX5jBT5fGffDSrufME+pqqXpiqr6u0as5EVUgsdli8STrPLdQb4FeGMDXX1RMNnkLce2iO9WG6d3rzx",
	"qGf8ehaKfNJVK5SW0i3WXmKalv0lKzWV030uSk7TSrGkRjUckHBBQOz197kfBazQ6iYzLlIBXpmfwdJz",
	"tl09WkiAKW9At1+0lMGnIJWz7XKf6B876Uhizy+kH+YY9aYjorh6Nk0iXT4lTmtKBQ/05ezMndsSWcpU",
	"g45xrCyWLjE5Na9ape3gAKTUsV0xM8adxlFH7quMPA+kzHzrcO5DcWKnhyg75nlqj3pUOlQ6Qd/ZPTzQ",
	"zFDlm4saS+p0sNR83QEhE11obbW2jO0KgeGQIhfttLZaO0g7LdUzArUzTXasvF7+2NFXT59/+WD0w++j",
	"0yd///HLs4f3kNkx1kN95USAzqSyh8aQ7HHSj3FlKvGpOAx96pnV7VsytqQxjXbusKDgaVDTJ1IBBLlK",
	"RBBzbHTRXHR7a6tOORKtN4JM43flg5hm3JXIvZH+UYPWwZDLgkxidPpk+OjHuUhjQmpFuBk8swnFAjgO",
	"NlIlbutQ7BLEr7oQ2+GD0+Ff38XYDu/fe/7T5yUI6526oGpFORuHNwL4VDC9HOInaYY2mGdBqm1HlJqN",
	"EAscgAIhjeOY3igxdAcXkbaRyDVmDKUec5ItzkK5kYGlOhS5WSfp58BuLW+yJsRbv7ISvut+W2ltrAF4",
	"s0HQUtiOW/i2/nh0+nT08/0SpPVu6/LMJe3IZoxbSa/RlhLDgp2XrsS/bn+dNAkaQnkZv52gO6307c5U",
	"33OujXn27ePRN4//Of312aM/s/amkovubHu1dmZKm8yNMVbeYl6Nv5NsAdY2IqhkbKHYIPQxqwgQkvrw",
	"OQ4S8t3RxlRlNes5rR/WQYPNI6792TYYQsw2GZZGPdtmq/Zjw4dfj77/Ih6ZS0G6bd1+bVIcboKFokaX",
	"LRNjyYvYWDBPNazEIzErczPX9JB1ZLBFHaVGqVoqo53L1Qo0VRC0LnIatGNFDdMlWcmUmKv4GD79bfjg",
	"SQ79dIf6cJ/UoNcLdlH5vRrmsbQ3B4PB4N8AAAD//zblOeGHNQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

	operation="generate"

[Error - 12:04:50 PM] 2024/06/06 12:04:50 tidy: diagnosing file:///Users/josh_sinrui/Documents/josh/personal/openapi-go-test/go.mod: err: exit status 1: stderr: go: downloading github.com/kr/pretty v0.3.1
go: finding module for package github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
go: github.com/lisyaoran51/openapi-go-test/tools imports
	github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen: cannot find module providing package github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen: module lookup disabled by GOPROXY=off
go: github.com/lisyaoran51/openapi-go-test/service/api imports
	github.com/gin-gonic/gin imports
	github.com/gin-gonic/gin/binding imports
	gopkg.in/yaml.v3 tested by
	gopkg.in/yaml.v3.test imports
	gopkg.in/check.v1 imports
	github.com/kr/pretty: module lookup disabled by GOPROXY=off


[Info  - 12:04:56 PM] 2024/06/06 12:04:56 background imports cache refresh starting

[Info  - 12:04:56 PM] 2024/06/06 12:04:56 background refresh finished after 17.489834ms

[Info  - 12:05:31 PM] 2024/06/06 12:05:31 background imports cache refresh starting

[Info  - 12:05:32 PM] 2024/06/06 12:05:32 background refresh finished after 20.572709ms

[Info  - 12:06:06 PM] 2024/06/06 12:06:06 go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen agencyBaseCompiled.yml

	operation="generate"

[Info  - 12:06:07 PM] 2024/06/06 12:06:07 // Package agencyBaseCompiled provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package agencyBaseCompiled

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Agency defines model for agency.
type Agency struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int       `json:"agencyStatus,omitempty"`
	CountryCode         *int       `json:"countryCode,omitempty"`
	CreatedAt           *time.Time `json:"createdAt,omitempty"`
	Id                  *int64     `json:"id,omitempty"`
	Name                *string    `json:"name,omitempty"`
	NationalPhoneNumber *string    `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64     `json:"platformId,omitempty"`
	PlatformName        *string    `json:"platformName,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyAddResponse defines model for agencyAddResponse.
type AgencyAddResponse struct {
	Code      *int    `json:"code,omitempty"`
	Data      *Id     `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyBatchUpdateRequest defines model for agencyBatchUpdateRequest.
type AgencyBatchUpdateRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int         `json:"agencyStatus,omitempty"`
	CountryCode         *int         `json:"countryCode,omitempty"`
	CreatedAt           *time.Time   `json:"createdAt,omitempty"`
	Id                  *int64       `json:"id,omitempty"`
	Name                *string      `json:"name,omitempty"`
	NationalPhoneNumber *string      `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64       `json:"platformId,omitempty"`
	PlatformName        *string      `json:"platformName,omitempty"`
	Query               *AgencyQuery `json:"query,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyBatchUpdateResponse defines model for agencyBatchUpdateResponse.
type AgencyBatchUpdateResponse struct {
	Code      *int                    `json:"code,omitempty"`
	Data      *map[string]interface{} `json:"data,omitempty"`
	Message   *string                 `json:"message,omitempty"`
	RequestId *string                 `json:"requestId,omitempty"`
	Success   *bool                   `json:"success,omitempty"`
}

// AgencyGetResponse defines model for agencyGetResponse.
type AgencyGetResponse struct {
	Code      *int      `json:"code,omitempty"`
	Data      *[]Agency `json:"data,omitempty"`
	Message   *string   `json:"message,omitempty"`
	RequestId *string   `json:"requestId,omitempty"`
	Success   *bool     `json:"success,omitempty"`
}

// AgencyPageRequest defines model for agencyPageRequest.
type AgencyPageRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int       `json:"agencyStatus,omitempty"`
	CountryCode         *int       `json:"countryCode,omitempty"`
	CreatedAt           *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom       *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo         *time.Time `json:"createdAtTo,omitempty"`
	Id                  *int64     `json:"id,omitempty"`
	IdIn                *[]int64   `json:"idIn,omitempty"`
	Name                *string    `json:"name,omitempty"`
	NationalPhoneNumber *string    `json:"nationalPhoneNumber,omitempty"`
	PageCurrent         *int       `json:"pageCurrent,omitempty"`
	PageSize            *int       `json:"pageSize,omitempty"`
	PlatformId          *int64     `json:"platformId,omitempty"`
	PlatformName        *string    `json:"platformName,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAtFrom      *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo        *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPageResponse defines model for agencyPageResponse.
type AgencyPageResponse struct {
	Code *int `json:"code,omitempty"`
	Data *struct {
		Current *int64    `json:"current,omitempty"`
		Pages   *int64    `json:"pages,omitempty"`
		Records *[]Agency `json:"records,omitempty"`
		Size    *int64    `json:"size,omitempty"`
		Total   *int64    `json:"total,omitempty"`
	} `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyPlan defines model for agencyPlan.
type AgencyPlan struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanAddResponse defines model for agencyPlanAddResponse.
type AgencyPlanAddResponse struct {
	Code      *int    `json:"code,omitempty"`
	Data      *Id     `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyPlanBatchUpdateRequest defines model for agencyPlanBatchUpdateRequest.
type AgencyPlanBatchUpdateRequest struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int             `json:"planStatus,omitempty"`
	PlatformId   *int64           `json:"platformId,omitempty"`
	PlatformName *string          `json:"platformName,omitempty"`
	Query        *AgencyPlanQuery `json:"query,omitempty"`
	ReferralCode *string          `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanBatchUpdateResponse defines model for agencyPlanBatchUpdateResponse.
type AgencyPlanBatchUpdateResponse struct {
	Code      *int                    `json:"code,omitempty"`
	Data      *map[string]interface{} `json:"data,omitempty"`
	Message   *string                 `json:"message,omitempty"`
	RequestId *string                 `json:"requestId,omitempty"`
	Success   *bool                   `json:"success,omitempty"`
}

// AgencyPlanGetByReferralCodeRequest defines model for agencyPlanGetByReferralCodeRequest.
type AgencyPlanGetByReferralCodeRequest struct {
	ReferralCode *string `json:"referralCode,omitempty"`
	UserId       *int64  `json:"userId,omitempty"`
}

// AgencyPlanGetByReferralCodeResponse defines model for agencyPlanGetByReferralCodeResponse.
type AgencyPlanGetByReferralCodeResponse struct {
	Code      *int        `json:"code,omitempty"`
	Data      *AgencyPlan `json:"data,omitempty"`
	Message   *string     `json:"message,omitempty"`
	RequestId *string     `json:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty"`
}

// AgencyPlanGetResponse defines model for agencyPlanGetResponse.
type AgencyPlanGetResponse struct {
	Code      *int          `json:"code,omitempty"`
	Data      *[]AgencyPlan `json:"data,omitempty"`
	Message   *string       `json:"message,omitempty"`
	RequestId *string       `json:"requestId,omitempty"`
	Success   *bool         `json:"success,omitempty"`
}

// AgencyPlanQuery defines model for agencyPlanQuery.
type AgencyPlanQuery struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom   *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo     *time.Time `json:"createdAtTo,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	IdIn     *[]int64 `json:"idIn,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAtFrom      *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo        *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanUpdateRequest defines model for agencyPlanUpdateRequest.
type AgencyPlanUpdateRequest struct {
	AgencyId   *int64  `json:"agencyId,omitempty"`
	AgencyName *string `json:"agencyName,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount *float32 `json:"discount,omitempty"`
	Id       *int64   `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int             `json:"planStatus,omitempty"`
	PlatformId   *int64           `json:"platformId,omitempty"`
	PlatformName *string          `json:"platformName,omitempty"`
	Query        *AgencyPlanQuery `json:"query,omitempty"`
	ReferralCode *string          `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyPlanUpdateResponse defines model for agencyPlanUpdateResponse.
type AgencyPlanUpdateResponse struct {
	Code      *int        `json:"code,omitempty"`
	Data      *AgencyPlan `json:"data,omitempty"`
	Message   *string     `json:"message,omitempty"`
	RequestId *string     `json:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty"`
}

// AgencyQuery defines model for agencyQuery.
type AgencyQuery struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int       `json:"agencyStatus,omitempty"`
	CountryCode         *int       `json:"countryCode,omitempty"`
	CreatedAt           *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom       *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo         *time.Time `json:"createdAtTo,omitempty"`
	Id                  *int64     `json:"id,omitempty"`
	IdIn                *[]int64   `json:"idIn,omitempty"`
	Name                *string    `json:"name,omitempty"`
	NationalPhoneNumber *string    `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64     `json:"platformId,omitempty"`
	PlatformName        *string    `json:"platformName,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAtFrom      *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo        *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyUpdateRequest defines model for agencyUpdateRequest.
type AgencyUpdateRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus        *int         `json:"agencyStatus,omitempty"`
	CountryCode         *int         `json:"countryCode,omitempty"`
	CreatedAt           *time.Time   `json:"createdAt,omitempty"`
	Id                  *int64       `json:"id,omitempty"`
	Name                *string      `json:"name,omitempty"`
	NationalPhoneNumber *string      `json:"nationalPhoneNumber,omitempty"`
	PlatformId          *int64       `json:"platformId,omitempty"`
	PlatformName        *string      `json:"platformName,omitempty"`
	Query               *AgencyQuery `json:"query,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyUpdateResponse defines model for agencyUpdateResponse.
type AgencyUpdateResponse struct {
	Code      *int    `json:"code,omitempty"`
	Data      *Agency `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyWithPlan defines model for agencyWithPlan.
type AgencyWithPlan struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus *int `json:"agencyStatus,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CountryCode     *int       `json:"countryCode,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`

	// Discount in percentage
	Discount            *float32 `json:"discount,omitempty"`
	Id                  *int64   `json:"id,omitempty"`
	Name                *string  `json:"name,omitempty"`
	NationalPhoneNumber *string  `json:"nationalPhoneNumber,omitempty"`
	PlanId              *int64   `json:"planId,omitempty"`
	PlanName            *string  `json:"planName,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AgencyWithPlanAddResponse defines model for agencyWithPlanAddResponse.
type AgencyWithPlanAddResponse struct {
	Code *int `json:"code,omitempty"`
	Data *struct {
		AgencyId     *int64 `json:"agencyId,omitempty"`
		AgencyPlanId *int64 `json:"agencyPlanId,omitempty"`
	} `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyWithPlanGetResponse defines model for agencyWithPlanGetResponse.
type AgencyWithPlanGetResponse struct {
	Code      *int              `json:"code,omitempty"`
	Data      *[]AgencyWithPlan `json:"data,omitempty"`
	Message   *string           `json:"message,omitempty"`
	RequestId *string           `json:"requestId,omitempty"`
	Success   *bool             `json:"success,omitempty"`
}

// AgencyWithPlanPageRequest defines model for agencyWithPlanPageRequest.
type AgencyWithPlanPageRequest struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus *int `json:"agencyStatus,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CountryCode     *int       `json:"countryCode,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom   *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo     *time.Time `json:"createdAtTo,omitempty"`

	// Discount in percentage
	Discount            *float32 `json:"discount,omitempty"`
	Id                  *int64   `json:"id,omitempty"`
	Name                *string  `json:"name,omitempty"`
	NationalPhoneNumber *string  `json:"nationalPhoneNumber,omitempty"`
	OrderByColumn       *string  `json:"orderByColumn,omitempty"`

	// OrderByDirection 1:asc 2:desc
	OrderByDirection *int    `json:"orderByDirection,omitempty"`
	PageCurrent      *int    `json:"pageCurrent,omitempty"`
	PageSize         *int    `json:"pageSize,omitempty"`
	PlanId           *int64  `json:"planId,omitempty"`
	PlanName         *string `json:"planName,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int `json:"registrationMethod,omitempty"`

	// Strategy 1:newest 2:oldest 3:newest active 4:oldest active
	Strategy      *int       `json:"strategy,omitempty"`
	UpdatedAtFrom *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo   *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

// AgencyWithPlanPageResponse defines model for agencyWithPlanPageResponse.
type AgencyWithPlanPageResponse struct {
	Code *int `json:"code,omitempty"`
	Data *struct {
		Current *int64            `json:"current,omitempty"`
		Pages   *int64            `json:"pages,omitempty"`
		Records *[]AgencyWithPlan `json:"records,omitempty"`
		Size    *int64            `json:"size,omitempty"`
		Total   *int64            `json:"total,omitempty"`
	} `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencyWithPlanQuery defines model for agencyWithPlanQuery.
type AgencyWithPlanQuery struct {
	// AgencyStatus 1:active 2:inactive
	AgencyStatus *int `json:"agencyStatus,omitempty"`

	// CardAssociation 1:visa 2:masterCard 3:AmericanExpress 4:JCB 5:unionPay
	CardAssociation *int       `json:"cardAssociation,omitempty"`
	CountryCode     *int       `json:"countryCode,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	CreatedAtFrom   *time.Time `json:"createdAtFrom,omitempty"`
	CreatedAtTo     *time.Time `json:"createdAtTo,omitempty"`

	// Discount in percentage
	Discount            *float32 `json:"discount,omitempty"`
	Id                  *int64   `json:"id,omitempty"`
	Name                *string  `json:"name,omitempty"`
	NationalPhoneNumber *string  `json:"nationalPhoneNumber,omitempty"`
	OrderByColumn       *string  `json:"orderByColumn,omitempty"`

	// OrderByDirection 1:asc 2:desc
	OrderByDirection *int    `json:"orderByDirection,omitempty"`
	PlanId           *int64  `json:"planId,omitempty"`
	PlanName         *string `json:"planName,omitempty"`

	// PlanStatus 1:active 2:inactive
	PlanStatus   *int    `json:"planStatus,omitempty"`
	PlatformId   *int64  `json:"platformId,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	ReferralCode *string `json:"referralCode,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int `json:"registrationMethod,omitempty"`

	// Strategy 1:newest 2:oldest 3:newest active 4:oldest active
	Strategy      *int       `json:"strategy,omitempty"`
	UpdatedAtFrom *time.Time `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo   *time.Time `json:"updatedAtTo,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

// Id defines model for id.
type Id struct {
	Id *int64 `json:"id,omitempty"`
}

// Platform defines model for platform.
type Platform struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *int64     `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`

	// RegistrationMethod 1:phone 2:email 3:username
	RegistrationMethod *int       `json:"registrationMethod,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// PlatformGetResponse defines model for platformGetResponse.
type PlatformGetResponse struct {
	Code      *int        `json:"code,omitempty"`
	Data      *[]Platform `json:"data,omitempty"`
	Message   *string     `json:"message,omitempty"`
	RequestId *string     `json:"requestId,omitempty"`
	Success   *bool       `json:"success,omitempty"`
}

// Result defines model for result.
type Result struct {
	Code      *int    `json:"code,omitempty"`
	Message   *string `json:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

// AgencybatchUpdateJSONRequestBody defines body for AgencybatchUpdate for application/json ContentType.
type AgencybatchUpdateJSONRequestBody = AgencyBatchUpdateRequest

// AgencyaddJSONRequestBody defines body for Agencyadd for application/json ContentType.
type AgencyaddJSONRequestBody = Agency

// AgencypagegetJSONRequestBody defines body for Agencypageget for application/json ContentType.
type AgencypagegetJSONRequestBody = AgencyPageRequest

// AgencyupdateJSONRequestBody defines body for Agencyupdate for application/json ContentType.
type AgencyupdateJSONRequestBody = AgencyUpdateRequest

// AgencygetJSONRequestBody defines body for Agencyget for application/json ContentType.
type AgencygetJSONRequestBody = AgencyQuery

// AgencyPlanbatchUpdateJSONRequestBody defines body for AgencyPlanbatchUpdate for application/json ContentType.
type AgencyPlanbatchUpdateJSONRequestBody = AgencyPlanBatchUpdateRequest

// AgencyPlanaddJSONRequestBody defines body for AgencyPlanadd for application/json ContentType.
type AgencyPlanaddJSONRequestBody = AgencyPlan

// AgencyPlangetByReferralCodeJSONRequestBody defines body for AgencyPlangetByReferralCode for application/json ContentType.
type AgencyPlangetByReferralCodeJSONRequestBody = AgencyPlanGetByReferralCodeRequest

// AgencyPlanupdateJSONRequestBody defines body for AgencyPlanupdate for application/json ContentType.
type AgencyPlanupdateJSONRequestBody = AgencyPlanUpdateRequest

// AgencyPlangetJSONRequestBody defines body for AgencyPlanget for application/json ContentType.
type AgencyPlangetJSONRequestBody = AgencyPlanQuery

// AgencyWithPlanaddJSONRequestBody defines body for AgencyWithPlanadd for application/json ContentType.
type AgencyWithPlanaddJSONRequestBody = AgencyWithPlan

// AgencyWithPlanpagegetJSONRequestBody defines body for AgencyWithPlanpageget for application/json ContentType.
type AgencyWithPlanpagegetJSONRequestBody = AgencyWithPlanPageRequest

// AgencyWithPlangetJSONRequestBody defines body for AgencyWithPlanget for application/json ContentType.
type AgencyWithPlangetJSONRequestBody = AgencyWithPlanQuery

// PlatformgetJSONRequestBody defines body for Platformget for application/json ContentType.
type PlatformgetJSONRequestBody = Platform

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// AgencybatchUpdateWithBody request with any body
	AgencybatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencybatchUpdate(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyaddWithBody request with any body
	AgencyaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencyadd(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencypagegetWithBody request with any body
	AgencypagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencypageget(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyupdateWithBody request with any body
	AgencyupdateWithBody(ctx context.Context, agencyId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencyupdate(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencygetWithBody request with any body
	AgencygetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Agencyget(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlanbatchUpdateWithBody request with any body
	AgencyPlanbatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanbatchUpdate(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlanaddWithBody request with any body
	AgencyPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanadd(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlangetByReferralCodeWithBody request with any body
	AgencyPlangetByReferralCodeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlangetByReferralCode(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlanupdateWithBody request with any body
	AgencyPlanupdateWithBody(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanupdate(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyPlangetWithBody request with any body
	AgencyPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyPlanget(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyWithPlanaddWithBody request with any body
	AgencyWithPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyWithPlanadd(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyWithPlanpagegetWithBody request with any body
	AgencyWithPlanpagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyWithPlanpageget(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AgencyWithPlangetWithBody request with any body
	AgencyWithPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AgencyWithPlanget(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PlatformgetWithBody request with any body
	PlatformgetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Platformget(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AgencybatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencybatchUpdateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencybatchUpdate(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencybatchUpdateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyaddRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencyadd(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyaddRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencypagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencypagegetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencypageget(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencypagegetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyupdateWithBody(ctx context.Context, agencyId int64, cont
	operation="generate"

[Info  - 12:06:07 PM] 2024/06/06 12:06:07 entType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyupdateRequestWithBody(c.Server, agencyId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencyupdate(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyupdateRequest(c.Server, agencyId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencygetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencygetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Agencyget(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencygetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanbatchUpdateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanbatchUpdateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanbatchUpdate(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanbatchUpdateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanaddRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanadd(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanaddRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlangetByReferralCodeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetByReferralCodeRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlangetByReferralCode(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetByReferralCodeRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanupdateWithBody(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanupdateRequestWithBody(c.Server, agencyPlanId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanupdate(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlanupdateRequest(c.Server, agencyPlanId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyPlanget(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyPlangetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanaddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanaddRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanadd(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanaddRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanpagegetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanpagegetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanpageget(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlanpagegetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlangetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlangetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AgencyWithPlanget(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAgencyWithPlangetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PlatformgetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlatformgetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Platformget(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlatformgetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAgencybatchUpdateRequest calls the generic AgencybatchUpdate builder with application/json body
func NewAgencybatchUpdateRequest(server string, body AgencybatchUpdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencybatchUpdateRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencybatchUpdateRequestWithBody generates requests for AgencybatchUpdate with any type of body
func NewAgencybatchUpdateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyaddRequest calls the generic Agencyadd builder with application/json body
func NewAgencyaddRequest(server string, body AgencyaddJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyaddRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyaddRequestWithBody generates requests for Agencyadd with any type of body
func NewAgencyaddRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencypagegetRequest calls the generic Agencypageget builder with application/json body
func NewAgencypagegetRequest(server string, body AgencypagegetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencypagegetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencypagegetRequestWithBody generates requests for Agencypageget with any type of body
func NewAgencypagegetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency/pageGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyupdateRequest calls the generic Agencyupdate builder with application/json body
func NewAgencyupdateRequest(server string, agencyId int64, body AgencyupdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyupdateRequestWithBody(server, agencyId, "application/json", bodyReader)
}

// NewAgencyupdateRequestWithBody generates requests for Agencyupdate with any type of body
func NewAgencyupdateRequestWithBody(server string, agencyId int64, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "agencyId", runtime.ParamLocationPath, agencyId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agency/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencygetRequest calls the generic Agencyget builder with application/json body
func NewAgencygetRequest(server string, body AgencygetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencygetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencygetRequestWithBody generates requests for Agencyget with any type of body
func NewAgencygetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlanbatchUpdateRequest calls the generic AgencyPlanbatchUpdate builder with application/json body
func NewAgencyPlanbatchUpdateRequest(server string, body AgencyPlanbatchUpdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlanbatchUpdateRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlanbatchUpdateRequestWithBody generates requests for AgencyPlanbatchUpdate with any type of body
func NewAgencyPlanbatchUpdateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlanaddRequest calls the generic AgencyPlanadd builder with application/json body
func NewAgencyPlanaddRequest(server string, body AgencyPlanaddJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlanaddRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlanaddRequestWithBody generates requests for AgencyPlanadd with any type of body
func NewAgencyPlanaddRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlangetByReferralCodeRequest calls the generic AgencyPlangetByReferralCode builder with application/json body
func NewAgencyPlangetByReferralCodeRequest(server string, body AgencyPlangetByReferralCodeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlangetByReferralCodeRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlangetByReferralCodeRequestWithBody generates requests for AgencyPlangetByReferralCode with any type of body
func NewAgencyPlangetByReferralCodeRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan/byReferralCodeGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlanupdateRequest calls the generic AgencyPlanupdate builder with application/json body
func NewAgencyPlanupdateRequest(server string, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlanupdateRequestWithBody(server, agencyPlanId, "application/json", bodyReader)
}

// NewAgencyPlanupdateRequestWithBody generates requests for AgencyPlanupdate with any type of body
func NewAgencyPlanupdateRequestWithBody(server string, agencyPlanId int64, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "agencyPlanId", runtime.ParamLocationPath, agencyPlanId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlan/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyPlangetRequest calls the generic AgencyPlanget builder with application/json body
func NewAgencyPlangetRequest(server string, body AgencyPlangetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyPlangetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyPlangetRequestWithBody generates requests for AgencyPlanget with any type of body
func NewAgencyPlangetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyPlanGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyWithPlanaddRequest calls the generic AgencyWithPlanadd builder with application/json body
func NewAgencyWithPlanaddRequest(server string, body AgencyWithPlanaddJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyWithPlanaddRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyWithPlanaddRequestWithBody generates requests for AgencyWithPlanadd with any type of body
func NewAgencyWithPlanaddRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyWithPlan")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyWithPlanpagegetRequest calls the generic AgencyWithPlanpageget builder with application/json body
func NewAgencyWithPlanpagegetRequest(server string, body AgencyWithPlanpagegetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyWithPlanpagegetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyWithPlanpagegetRequestWithBody generates requests for AgencyWithPlanpageget with any type of body
func NewAgencyWithPlanpagegetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyWithPlan/pageGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewAgencyWithPlangetRequest calls the generic AgencyWithPlanget builder with application/json body
func NewAgencyWithPlangetRequest(server string, body AgencyWithPlangetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAgencyWithPlangetRequestWithBody(server, "application/json", bodyReader)
}

// NewAgencyWithPlangetRequestWithBody generates requests for AgencyWithPlanget with any type of body
func NewAgencyWithPlangetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/agencyWithPlanGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPlatformgetRequest calls the generic Platformget builder with application/json body
func NewPlatformgetRequest(server string, body PlatformgetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPlatformgetRequestWithBody(server, "application/json", bodyReader)
}

// NewPlatformgetRequestWithBody generates requests for Platformget with any type of body
func NewPlatformgetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/platformGet")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// AgencybatchUpdateWithBodyWithResponse request with any body
	AgencybatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error)

	AgencybatchUpdateWithResponse(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error)

	// AgencyaddWithBodyWithResponse request with any body
	AgencyaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error)

	AgencyaddWithResponse(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error)

	// AgencypagegetWithBodyWithResponse request with any body
	AgencypagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error)

	AgencypagegetWithResponse(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error)

	// AgencyupdateWithBodyWithResponse request with any body
	AgencyupdateWithBodyWithResponse(ctx context.Context, agencyId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error)

	AgencyupdateWithResponse(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error)

	// AgencygetWithBodyWithResponse request with any body
	AgencygetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencygetResponse, error)

	AgencygetWithResponse(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencygetResponse, error)

	// AgencyPlanbatchUpdateWithBodyWithResponse request with any body
	AgencyPlanbatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error)

	AgencyPlanbatchUpdateWithResponse(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error)

	// AgencyPlanaddWithBodyWithResponse request with any body
	AgencyPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error)

	AgencyPlanaddWithResponse(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error)

	// AgencyPlangetByReferralCodeWithBodyWithResponse request with any body
	AgencyPlangetByReferralCodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error)

	AgencyPlangetByReferralCodeWithResponse(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error)

	// AgencyPlanupdateWithBodyWithResponse request with any body
	AgencyPlanupdateWithBodyWithResponse(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error)

	AgencyPlanupdateWithResponse(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error)

	// AgencyPlangetWithBodyWithResponse request with any body
	AgencyPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error)

	AgencyPlangetWithResponse(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error)

	// AgencyWithPlanaddWithBodyWithResponse request with any body
	AgencyWithPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error)

	AgencyWithPlanaddWithResponse(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error)

	// AgencyWithPlanpagegetWithBodyWithResponse request with any body
	AgencyWithPlanpagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error)

	AgencyWithPlanpagegetWithResponse(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error)

	// AgencyWithPlangetWithBodyWithResponse request with any body
	AgencyWithPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error)

	AgencyWithPlangetWithResponse(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error)

	// PlatformgetWithBodyWithResponse request with any body
	PlatformgetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error)

	PlatformgetWithResponse(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error)
}

type AgencybatchUpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyBatchUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencybatchUpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencybatchUpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyaddResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyAddResponse
}

// Status returns HTTPResponse.Status
func (r AgencyaddResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyaddResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencypagegetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPageResponse
}

// Status returns HTTPResponse.Status
func (r AgencypagegetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencypagegetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyupdateResponse struct {
	Body         []byte
	
	operation="generate"

[Info  - 12:06:07 PM] 2024/06/06 12:06:07 HTTPResponse *http.Response
	JSON200      *AgencyUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencyupdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyupdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencygetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyGetResponse
}

// Status returns HTTPResponse.Status
func (r AgencygetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencygetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlanbatchUpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanBatchUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlanbatchUpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlanbatchUpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlanaddResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanAddResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlanaddResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlanaddResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlangetByReferralCodeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanGetByReferralCodeResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlangetByReferralCodeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlangetByReferralCodeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlanupdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanUpdateResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlanupdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlanupdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyPlangetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyPlanGetResponse
}

// Status returns HTTPResponse.Status
func (r AgencyPlangetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyPlangetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyWithPlanaddResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyWithPlanAddResponse
}

// Status returns HTTPResponse.Status
func (r AgencyWithPlanaddResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyWithPlanaddResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyWithPlanpagegetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyWithPlanPageResponse
}

// Status returns HTTPResponse.Status
func (r AgencyWithPlanpagegetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyWithPlanpagegetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AgencyWithPlangetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AgencyWithPlanGetResponse
}

// Status returns HTTPResponse.Status
func (r AgencyWithPlangetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AgencyWithPlangetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PlatformgetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PlatformGetResponse
}

// Status returns HTTPResponse.Status
func (r PlatformgetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PlatformgetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AgencybatchUpdateWithBodyWithResponse request with arbitrary body returning *AgencybatchUpdateResponse
func (c *ClientWithResponses) AgencybatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error) {
	rsp, err := c.AgencybatchUpdateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencybatchUpdateResponse(rsp)
}

func (c *ClientWithResponses) AgencybatchUpdateWithResponse(ctx context.Context, body AgencybatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencybatchUpdateResponse, error) {
	rsp, err := c.AgencybatchUpdate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencybatchUpdateResponse(rsp)
}

// AgencyaddWithBodyWithResponse request with arbitrary body returning *AgencyaddResponse
func (c *ClientWithResponses) AgencyaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error) {
	rsp, err := c.AgencyaddWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyaddResponse(rsp)
}

func (c *ClientWithResponses) AgencyaddWithResponse(ctx context.Context, body AgencyaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyaddResponse, error) {
	rsp, err := c.Agencyadd(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyaddResponse(rsp)
}

// AgencypagegetWithBodyWithResponse request with arbitrary body returning *AgencypagegetResponse
func (c *ClientWithResponses) AgencypagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error) {
	rsp, err := c.AgencypagegetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencypagegetResponse(rsp)
}

func (c *ClientWithResponses) AgencypagegetWithResponse(ctx context.Context, body AgencypagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencypagegetResponse, error) {
	rsp, err := c.Agencypageget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencypagegetResponse(rsp)
}

// AgencyupdateWithBodyWithResponse request with arbitrary body returning *AgencyupdateResponse
func (c *ClientWithResponses) AgencyupdateWithBodyWithResponse(ctx context.Context, agencyId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error) {
	rsp, err := c.AgencyupdateWithBody(ctx, agencyId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyupdateResponse(rsp)
}

func (c *ClientWithResponses) AgencyupdateWithResponse(ctx context.Context, agencyId int64, body AgencyupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyupdateResponse, error) {
	rsp, err := c.Agencyupdate(ctx, agencyId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyupdateResponse(rsp)
}

// AgencygetWithBodyWithResponse request with arbitrary body returning *AgencygetResponse
func (c *ClientWithResponses) AgencygetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencygetResponse, error) {
	rsp, err := c.AgencygetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencygetResponse(rsp)
}

func (c *ClientWithResponses) AgencygetWithResponse(ctx context.Context, body AgencygetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencygetResponse, error) {
	rsp, err := c.Agencyget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencygetResponse(rsp)
}

// AgencyPlanbatchUpdateWithBodyWithResponse request with arbitrary body returning *AgencyPlanbatchUpdateResponse
func (c *ClientWithResponses) AgencyPlanbatchUpdateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error) {
	rsp, err := c.AgencyPlanbatchUpdateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanbatchUpdateResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlanbatchUpdateWithResponse(ctx context.Context, body AgencyPlanbatchUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanbatchUpdateResponse, error) {
	rsp, err := c.AgencyPlanbatchUpdate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanbatchUpdateResponse(rsp)
}

// AgencyPlanaddWithBodyWithResponse request with arbitrary body returning *AgencyPlanaddResponse
func (c *ClientWithResponses) AgencyPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error) {
	rsp, err := c.AgencyPlanaddWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanaddResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlanaddWithResponse(ctx context.Context, body AgencyPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanaddResponse, error) {
	rsp, err := c.AgencyPlanadd(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanaddResponse(rsp)
}

// AgencyPlangetByReferralCodeWithBodyWithResponse request with arbitrary body returning *AgencyPlangetByReferralCodeResponse
func (c *ClientWithResponses) AgencyPlangetByReferralCodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error) {
	rsp, err := c.AgencyPlangetByReferralCodeWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetByReferralCodeResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlangetByReferralCodeWithResponse(ctx context.Context, body AgencyPlangetByReferralCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetByReferralCodeResponse, error) {
	rsp, err := c.AgencyPlangetByReferralCode(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetByReferralCodeResponse(rsp)
}

// AgencyPlanupdateWithBodyWithResponse request with arbitrary body returning *AgencyPlanupdateResponse
func (c *ClientWithResponses) AgencyPlanupdateWithBodyWithResponse(ctx context.Context, agencyPlanId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error) {
	rsp, err := c.AgencyPlanupdateWithBody(ctx, agencyPlanId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanupdateResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlanupdateWithResponse(ctx context.Context, agencyPlanId int64, body AgencyPlanupdateJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlanupdateResponse, error) {
	rsp, err := c.AgencyPlanupdate(ctx, agencyPlanId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlanupdateResponse(rsp)
}

// AgencyPlangetWithBodyWithResponse request with arbitrary body returning *AgencyPlangetResponse
func (c *ClientWithResponses) AgencyPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error) {
	rsp, err := c.AgencyPlangetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetResponse(rsp)
}

func (c *ClientWithResponses) AgencyPlangetWithResponse(ctx context.Context, body AgencyPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyPlangetResponse, error) {
	rsp, err := c.AgencyPlanget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyPlangetResponse(rsp)
}

// AgencyWithPlanaddWithBodyWithResponse request with arbitrary body returning *AgencyWithPlanaddResponse
func (c *ClientWithResponses) AgencyWithPlanaddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error) {
	rsp, err := c.AgencyWithPlanaddWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanaddResponse(rsp)
}

func (c *ClientWithResponses) AgencyWithPlanaddWithResponse(ctx context.Context, body AgencyWithPlanaddJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanaddResponse, error) {
	rsp, err := c.AgencyWithPlanadd(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanaddResponse(rsp)
}

// AgencyWithPlanpagegetWithBodyWithResponse request with arbitrary body returning *AgencyWithPlanpagegetResponse
func (c *ClientWithResponses) AgencyWithPlanpagegetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error) {
	rsp, err := c.AgencyWithPlanpagegetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanpagegetResponse(rsp)
}

func (c *ClientWithResponses) AgencyWithPlanpagegetWithResponse(ctx context.Context, body AgencyWithPlanpagegetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlanpagegetResponse, error) {
	rsp, err := c.AgencyWithPlanpageget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlanpagegetResponse(rsp)
}

// AgencyWithPlangetWithBodyWithResponse request with arbitrary body returning *AgencyWithPlangetResponse
func (c *ClientWithResponses) AgencyWithPlangetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error) {
	rsp, err := c.AgencyWithPlangetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlangetResponse(rsp)
}

func (c *ClientWithResponses) AgencyWithPlangetWithResponse(ctx context.Context, body AgencyWithPlangetJSONRequestBody, reqEditors ...RequestEditorFn) (*AgencyWithPlangetResponse, error) {
	rsp, err := c.AgencyWithPlanget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAgencyWithPlangetResponse(rsp)
}

// PlatformgetWithBodyWithResponse request with arbitrary body returning *PlatformgetResponse
func (c *ClientWithResponses) PlatformgetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error) {
	rsp, err := c.PlatformgetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlatformgetResponse(rsp)
}

func (c *ClientWithResponses) PlatformgetWithResponse(ctx context.Context, body PlatformgetJSONRequestBody, reqEditors ...RequestEditorFn) (*PlatformgetResponse, error) {
	rsp, err := c.Platformget(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlatformgetResponse(rsp)
}

// ParseAgencybatchUpdateResponse parses an HTTP response from a AgencybatchUpdateWithResponse call
func ParseAgencybatchUpdateResponse(rsp *http.Response) (*AgencybatchUpdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencybatchUpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyBatchUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyaddResponse parses an HTTP response from a AgencyaddWithResponse call
func ParseAgencyaddResponse(rsp *http.Response) (*AgencyaddResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyaddResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyAddResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencypagegetResponse parses an HTTP response from a AgencypagegetWithResponse call
func ParseAgencypagegetResponse(rsp *http.Response) (*AgencypagegetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencypagegetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyupdateResponse parses an HTTP response from a AgencyupdateWithResponse call
func ParseAgencyupdateResponse(rsp *http.Response) (*AgencyupdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyupdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencygetResponse parses an HTTP response from a AgencygetWithResponse call
func ParseAgencygetResponse(rsp *http.Response) (*AgencygetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencygetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlanbatchUpdateResponse parses an HTTP response from a AgencyPlanbatchUpdateWithResponse call
func ParseAgencyPlanbatchUpdateResponse(rsp *http.Response) (*AgencyPlanbatchUpdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlanbatchUpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanBatchUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlanaddResponse parses an HTTP response from a AgencyPlanaddWithResponse call
func ParseAgencyPlanaddResponse(rsp *http.Response) (*AgencyPlanaddResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlanaddResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanAddResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlangetByReferralCodeResponse parses an HTTP response from a AgencyPlangetByReferralCodeWithResponse call
func ParseAgencyPlangetByReferralCodeResponse(rsp *http.Response) (*AgencyPlangetByReferralCodeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlangetByReferralCodeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanGetByReferralCodeResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlanupdateResponse parses an HTTP response from a AgencyPlanupdateWithResponse call
func ParseAgencyPlanupdateResponse(rsp *http.Response) (*AgencyPlanupdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlanupdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanUpdateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyPlangetResponse parses an HTTP response from a AgencyPlangetWithResponse call
func ParseAgencyPlangetResponse(rsp *http.Response) (*AgencyPlangetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyPlangetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyPlanGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyWithPlanaddResponse parses an HTTP response from a AgencyWithPlanaddWithResponse call
func ParseAgencyWithPlanaddResponse(rsp *http.Response) (*AgencyWithPlanaddResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyWithPlanaddResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyWithPlanAddResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyWithPlanpagegetResponse parses an HTTP response from a AgencyWithPlanpagegetWithResponse call
func ParseAgencyWithPlanpagegetResponse(rsp *http.Response) (*AgencyWithPlanpagegetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyWithPlanpagegetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyWithPlanPageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAgencyWithPlangetResponse parses an HTTP response from a AgencyWithPlangetWithResponse call
func ParseAgencyWithPlangetResponse(rsp *http.Response) (*AgencyWithPlangetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AgencyWithPlangetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AgencyWithPlanGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePlatformgetResponse parses an HTTP response from a PlatformgetWithResponse call
func ParsePlatformgetResponse(rsp *http.Response) (*PlatformgetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PlatformgetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PlatformGetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (PATCH /agency)
	AgencybatchUpdate(ctx echo.Context) error

	// (POST /agency)
	Agencyadd(ctx echo.Context) error

	// (POST /agency/pageGet)
	Agencypageget(ctx echo.Context) error

	// (PATCH /agency/{agencyId})
	Agencyupdate(ctx echo.Context, agencyId int64) error

	// (POST /agencyGet)
	Agencyget(ctx echo.Context) error

	// (PATCH /agencyPlan)
	AgencyPlanbatchUpdate(ctx echo.Context) error

	// (POST /agencyPlan)
	AgencyPlanadd(ctx echo.Context) error

	// (POST /agencyPlan/byReferralCodeGet)
	AgencyPlangetByReferralCode(ctx echo.Context) error

	// (PATCH /agencyPlan/{agencyPlanId})
	AgencyPlanupdate(ctx echo.Context, agencyPlanId int64) error

	// (POST /agencyPlanGet)
	AgencyPlanget(ctx echo.Context) error

	// (POST /agencyWithPlan)
	AgencyWithPlanadd(ctx echo.Context) error

	// (POST /agencyWithPlan/pageGet)
	AgencyWithPlanpageget(ctx echo.Context) error

	// (POST /agencyWithPlanGet)
	AgencyWithPlanget(ctx echo.Context) error

	// (POST /platformGet)
	Platformget(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AgencybatchUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) AgencybatchUpdate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencybatchUpdate(ctx)
	return err
}

// Agencyadd converts echo context to params.
func (w *ServerInterfaceWrapper) Agencyadd(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencyadd(ctx)
	return err
}

// Agencypageget converts echo context to params.
func (w *ServerInterfaceWrapper) Agencypageget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencypageget(ctx)
	return err
}

// Agencyupdate converts echo context to params.
func (w *ServerInterfaceWrapper) Agencyupdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "agencyId" -------------
	var agencyId int64

	err = runtime.BindStyledParameterWithOptions("simple", "agencyId", ctx.Param("agencyId"), &agencyId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter agencyId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencyupdate(ctx, agencyId)
	return err
}

// Agencyget converts echo context to params.
func (w *ServerInterfaceWrapper) Agencyget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Agencyget(ctx)
	return err
}

// AgencyPlanbatchUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanbatchUpdate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanbatchUpdate(ctx)
	return err
}

// AgencyPlanadd converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanadd(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanadd(ctx)
	return err
}

// AgencyPlangetByReferralCode converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlangetByReferralCode(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlangetByReferralCode(ctx)
	return err
}

// AgencyPlanupdate converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanupdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "agencyPlanId" -------------
	var agencyPlanId int64

	err = runtime.BindStyledParameterWithOptions("simple", "agencyPlanId", ctx.Param("agencyPlanId"), &agencyPlanId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter agencyPlanId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanupdate(ctx, agencyPlanId)
	return err
}

// AgencyPlanget converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyPlanget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyPlanget(ctx)
	return err
}

// AgencyWithPlanadd converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyWithPlanadd(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyWithPlanadd(ctx)
	return err
}

// AgencyWithPlanpageget converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyWithPlanpageget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyWithPlanpageget(ctx)
	return err
}

// AgencyWithPlanget converts echo context to params.
func (w *ServerInterfaceWrapper) AgencyWithPlanget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AgencyWithPlanget(ctx)
	return err
}

// Platformget converts echo context to params.
func (w *ServerInterfaceWrapper) Platformget(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Platformget(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandler
	operation="generate"

[Info  - 12:06:07 PM] 2024/06/06 12:06:07 sWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.PATCH(baseURL+"/agency", wrapper.AgencybatchUpdate)
	router.POST(baseURL+"/agency", wrapper.Agencyadd)
	router.POST(baseURL+"/agency/pageGet", wrapper.Agencypageget)
	router.PATCH(baseURL+"/agency/:agencyId", wrapper.Agencyupdate)
	router.POST(baseURL+"/agencyGet", wrapper.Agencyget)
	router.PATCH(baseURL+"/agencyPlan", wrapper.AgencyPlanbatchUpdate)
	router.POST(baseURL+"/agencyPlan", wrapper.AgencyPlanadd)
	router.POST(baseURL+"/agencyPlan/byReferralCodeGet", wrapper.AgencyPlangetByReferralCode)
	router.PATCH(baseURL+"/agencyPlan/:agencyPlanId", wrapper.AgencyPlanupdate)
	router.POST(baseURL+"/agencyPlanGet", wrapper.AgencyPlanget)
	router.POST(baseURL+"/agencyWithPlan", wrapper.AgencyWithPlanadd)
	router.POST(baseURL+"/agencyWithPlan/pageGet", wrapper.AgencyWithPlanpageget)
	router.POST(baseURL+"/agencyWithPlanGet", wrapper.AgencyWithPlanget)
	router.POST(baseURL+"/platformGet", wrapper.Platformget)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xaT4/bRBT/KtbAAaTdZNktHCwhsbuFakG0S7sVh6pCE8/bZCp7xp0Zl4ZVDnAoRUK0",
	"B7QIISEoEuqxElwqEJ+GpPRboBnbiRPb8SRZx1vKqel4/rz5/d68v3uCPB6EnAFTErknSHo9CLD5ibvA",
	"vL7+FQoeglAUMuPXFFaR+T8B6QkaKsoZctEbLvYUvQPOtktZ/BNtINUPAbmIMgVdEGiwgTweMSX6+5yA",
	"3qRgggCsgOwq/fmYiwAr5CKCFWwqGmQ2lUpQ1tVLKJmaS5l660L+8A10d/OYgk824a4SeBMzxhVOxH/H",
	"40Grg3lACY9aQb+DFZWhH8nWZFrrCHd8OCCv6Y2dtx2bJQfkqB9Ca/f60ZXXtagMB9l7T+7AzHzsH/Y4",
	"g8tR0AFROC/0sdI3PbC6c2bB5bKTBXSpVMKc/yGoHidF7IZaLmfbhQBT39lxIwnCXKbozCjUfJFPsDWJ",
	"g/EI79wCT+ldYoXbJeQqyJAzacTHvn/lGLk3TtCrAo6Ri15pTzS5nahxW4CMfL3JrBITrLD+d95iSork",
	"uTmWaA8rr3fdXPEq3I5AKnvBkseVF+x2BKJfJVm8+iMz1V7EMwdv5th5glwCVYMAVEEg7cBCE+mwELgC",
	"t0PcXZbThJWczCHuwn4kBDBVbPD0hGv0s0JzWC3rmWM7PepNJLexNbgbL7OYK8DjgsiV6dxAMgHP4lDF",
	"Ffat5uaRn8uFj1mZx7S21PH0UjvtYUF2peQeTZxW3kjfoRI7226ApQKxjwVxdtzdAAT1MHv3bihASueC",
	"+/7+nvOmGzHK2SHuF3vpxZ0wodL49rxclDkhCA+Ywt3MUhY7uf+I/w59zFaNjerw7ccgBPZnwq0Xxvnr",
	"d3X+AgAt1epBgLEZKwUCegeLYCAnbpMBgRbmEqi9/tWMamYQnD6gUn+1Jlo+l/laViDTWnUuqxMW+DUb",
	"VCWqu0hgNVbVs3klYwf1nuCBvZMaLzvii6SXB2wKIJtAYyZGSYziovKOl9nLW8XCi2OzajNXZ/EOl1Lm",
	"/xX5jBT5fGffDSrufME+pqqXpiqr6u0as5EVUgsdli8STrPLdQb4FeGMDXX1RMNnkLce2iO9WG6d3rzx",
	"qGf8ehaKfNJVK5SW0i3WXmKalv0lKzWV030uSk7TSrGkRjUckHBBQOz197kfBazQ6iYzLlIBXpmfwdJz",
	"tl09WkiAKW9At1+0lMGnIJWz7XKf6B876Uhizy+kH+YY9aYjorh6Nk0iXT4lTmtKBQ/05ezMndsSWcpU",
	"g45xrCyWLjE5Na9ape3gAKTUsV0xM8adxlFH7quMPA+kzHzrcO5DcWKnhyg75nlqj3pUOlQ6Qd/ZPTzQ",
	"zFDlm4saS+p0sNR83QEhE11obbW2jO0KgeGQIhfttLZaO0g7LdUzArUzTXasvF7+2NFXT59/+WD0w++j",
	"0yd///HLs4f3kNkx1kN95USAzqSyh8aQ7HHSj3FlKvGpOAx96pnV7VsytqQxjXbusKDgaVDTJ1IBBLlK",
	"RBBzbHTRXHR7a6tOORKtN4JM43flg5hm3JXIvZH+UYPWwZDLgkxidPpk+OjHuUhjQmpFuBk8swnFAjgO",
	"NlIlbutQ7BLEr7oQ2+GD0+Ff38XYDu/fe/7T5yUI6526oGpFORuHNwL4VDC9HOInaYY2mGdBqm1HlJqN",
	"EAscgAIhjeOY3igxdAcXkbaRyDVmDKUec5ItzkK5kYGlOhS5WSfp58BuLW+yJsRbv7ISvut+W2ltrAF4",
	"s0HQUtiOW/i2/nh0+nT08/0SpPVu6/LMJe3IZoxbSa/RlhLDgp2XrsS/bn+dNAkaQnkZv52gO6307c5U",
	"33OujXn27ePRN4//Of312aM/s/amkovubHu1dmZKm8yNMVbeYl6Nv5NsAdY2IqhkbKHYIPQxqwgQkvrw",
	"OQ4S8t3RxlRlNes5rR/WQYPNI6792TYYQsw2GZZGPdtmq/Zjw4dfj77/Ih6ZS0G6bd1+bVIcboKFokaX",
	"LRNjyYvYWDBPNazEIzErczPX9JB1ZLBFHaVGqVoqo53L1Qo0VRC0LnIatGNFDdMlWcmUmKv4GD79bfjg",
	"SQ79dIf6cJ/UoNcLdlH5vRrmsbQ3B4PB4N8AAAD//zblOeGHNQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}