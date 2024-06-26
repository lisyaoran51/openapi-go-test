// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (PATCH /agency)
	AgencybatchUpdate(c *gin.Context)

	// (POST /agency)
	Agencyadd(c *gin.Context)

	// (POST /agency/pageGet)
	Agencypageget(c *gin.Context)

	// (PATCH /agency/{agencyId})
	Agencyupdate(c *gin.Context, agencyId int64)

	// (POST /agencyGet)
	Agencyget(c *gin.Context)

	// (PATCH /agencyPlan)
	AgencyPlanbatchUpdate(c *gin.Context)

	// (POST /agencyPlan)
	AgencyPlanadd(c *gin.Context)

	// (POST /agencyPlan/byReferralCodeGet)
	AgencyPlangetByReferralCode(c *gin.Context)

	// (PATCH /agencyPlan/{agencyPlanId})
	AgencyPlanupdate(c *gin.Context, agencyPlanId int64)

	// (POST /agencyPlanGet)
	AgencyPlanget(c *gin.Context)

	// (POST /agencyWithPlan)
	AgencyWithPlanadd(c *gin.Context)

	// (POST /agencyWithPlan/pageGet)
	AgencyWithPlanpageget(c *gin.Context)

	// (POST /agencyWithPlanGet)
	AgencyWithPlanget(c *gin.Context)

	// (POST /platformGet)
	Platformget(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AgencybatchUpdate operation middleware
func (siw *ServerInterfaceWrapper) AgencybatchUpdate(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencybatchUpdate(c)
}

// Agencyadd operation middleware
func (siw *ServerInterfaceWrapper) Agencyadd(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Agencyadd(c)
}

// Agencypageget operation middleware
func (siw *ServerInterfaceWrapper) Agencypageget(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Agencypageget(c)
}

// Agencyupdate operation middleware
func (siw *ServerInterfaceWrapper) Agencyupdate(c *gin.Context) {

	var err error

	// ------------- Path parameter "agencyId" -------------
	var agencyId int64

	err = runtime.BindStyledParameter("simple", false, "agencyId", c.Param("agencyId"), &agencyId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter agencyId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Agencyupdate(c, agencyId)
}

// Agencyget operation middleware
func (siw *ServerInterfaceWrapper) Agencyget(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Agencyget(c)
}

// AgencyPlanbatchUpdate operation middleware
func (siw *ServerInterfaceWrapper) AgencyPlanbatchUpdate(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyPlanbatchUpdate(c)
}

// AgencyPlanadd operation middleware
func (siw *ServerInterfaceWrapper) AgencyPlanadd(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyPlanadd(c)
}

// AgencyPlangetByReferralCode operation middleware
func (siw *ServerInterfaceWrapper) AgencyPlangetByReferralCode(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyPlangetByReferralCode(c)
}

// AgencyPlanupdate operation middleware
func (siw *ServerInterfaceWrapper) AgencyPlanupdate(c *gin.Context) {

	var err error

	// ------------- Path parameter "agencyPlanId" -------------
	var agencyPlanId int64

	err = runtime.BindStyledParameter("simple", false, "agencyPlanId", c.Param("agencyPlanId"), &agencyPlanId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter agencyPlanId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyPlanupdate(c, agencyPlanId)
}

// AgencyPlanget operation middleware
func (siw *ServerInterfaceWrapper) AgencyPlanget(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyPlanget(c)
}

// AgencyWithPlanadd operation middleware
func (siw *ServerInterfaceWrapper) AgencyWithPlanadd(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyWithPlanadd(c)
}

// AgencyWithPlanpageget operation middleware
func (siw *ServerInterfaceWrapper) AgencyWithPlanpageget(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyWithPlanpageget(c)
}

// AgencyWithPlanget operation middleware
func (siw *ServerInterfaceWrapper) AgencyWithPlanget(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AgencyWithPlanget(c)
}

// Platformget operation middleware
func (siw *ServerInterfaceWrapper) Platformget(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Platformget(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.PATCH(options.BaseURL+"/agency", wrapper.AgencybatchUpdate)
	router.POST(options.BaseURL+"/agency", wrapper.Agencyadd)
	router.POST(options.BaseURL+"/agency/pageGet", wrapper.Agencypageget)
	router.PATCH(options.BaseURL+"/agency/:agencyId", wrapper.Agencyupdate)
	router.POST(options.BaseURL+"/agencyGet", wrapper.Agencyget)
	router.PATCH(options.BaseURL+"/agencyPlan", wrapper.AgencyPlanbatchUpdate)
	router.POST(options.BaseURL+"/agencyPlan", wrapper.AgencyPlanadd)
	router.POST(options.BaseURL+"/agencyPlan/byReferralCodeGet", wrapper.AgencyPlangetByReferralCode)
	router.PATCH(options.BaseURL+"/agencyPlan/:agencyPlanId", wrapper.AgencyPlanupdate)
	router.POST(options.BaseURL+"/agencyPlanGet", wrapper.AgencyPlanget)
	router.POST(options.BaseURL+"/agencyWithPlan", wrapper.AgencyWithPlanadd)
	router.POST(options.BaseURL+"/agencyWithPlan/pageGet", wrapper.AgencyWithPlanpageget)
	router.POST(options.BaseURL+"/agencyWithPlanGet", wrapper.AgencyWithPlanget)
	router.POST(options.BaseURL+"/platformGet", wrapper.Platformget)
}

type AgencybatchUpdateRequestObject struct {
	Body *AgencybatchUpdateJSONRequestBody
}

type AgencybatchUpdateResponseObject interface {
	VisitAgencybatchUpdateResponse(w http.ResponseWriter) error
}

type AgencybatchUpdate200JSONResponse AgencyBatchUpdateResponse

func (response AgencybatchUpdate200JSONResponse) VisitAgencybatchUpdateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyaddRequestObject struct {
	Body *AgencyaddJSONRequestBody
}

type AgencyaddResponseObject interface {
	VisitAgencyaddResponse(w http.ResponseWriter) error
}

type Agencyadd200JSONResponse AgencyAddResponse

func (response Agencyadd200JSONResponse) VisitAgencyaddResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencypagegetRequestObject struct {
	Body *AgencypagegetJSONRequestBody
}

type AgencypagegetResponseObject interface {
	VisitAgencypagegetResponse(w http.ResponseWriter) error
}

type Agencypageget200JSONResponse AgencyPageResponse

func (response Agencypageget200JSONResponse) VisitAgencypagegetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyupdateRequestObject struct {
	AgencyId int64 `json:"agencyId"`
	Body     *AgencyupdateJSONRequestBody
}

type AgencyupdateResponseObject interface {
	VisitAgencyupdateResponse(w http.ResponseWriter) error
}

type Agencyupdate200JSONResponse AgencyUpdateResponse

func (response Agencyupdate200JSONResponse) VisitAgencyupdateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencygetRequestObject struct {
	Body *AgencygetJSONRequestBody
}

type AgencygetResponseObject interface {
	VisitAgencygetResponse(w http.ResponseWriter) error
}

type Agencyget200JSONResponse AgencyGetResponse

func (response Agencyget200JSONResponse) VisitAgencygetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyPlanbatchUpdateRequestObject struct {
	Body *AgencyPlanbatchUpdateJSONRequestBody
}

type AgencyPlanbatchUpdateResponseObject interface {
	VisitAgencyPlanbatchUpdateResponse(w http.ResponseWriter) error
}

type AgencyPlanbatchUpdate200JSONResponse AgencyPlanBatchUpdateResponse

func (response AgencyPlanbatchUpdate200JSONResponse) VisitAgencyPlanbatchUpdateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyPlanaddRequestObject struct {
	Body *AgencyPlanaddJSONRequestBody
}

type AgencyPlanaddResponseObject interface {
	VisitAgencyPlanaddResponse(w http.ResponseWriter) error
}

type AgencyPlanadd200JSONResponse AgencyPlanAddResponse

func (response AgencyPlanadd200JSONResponse) VisitAgencyPlanaddResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyPlangetByReferralCodeRequestObject struct {
	Body *AgencyPlangetByReferralCodeJSONRequestBody
}

type AgencyPlangetByReferralCodeResponseObject interface {
	VisitAgencyPlangetByReferralCodeResponse(w http.ResponseWriter) error
}

type AgencyPlangetByReferralCode200JSONResponse AgencyPlanGetByReferralCodeResponse

func (response AgencyPlangetByReferralCode200JSONResponse) VisitAgencyPlangetByReferralCodeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyPlanupdateRequestObject struct {
	AgencyPlanId int64 `json:"agencyPlanId"`
	Body         *AgencyPlanupdateJSONRequestBody
}

type AgencyPlanupdateResponseObject interface {
	VisitAgencyPlanupdateResponse(w http.ResponseWriter) error
}

type AgencyPlanupdate200JSONResponse AgencyPlanUpdateResponse

func (response AgencyPlanupdate200JSONResponse) VisitAgencyPlanupdateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyPlangetRequestObject struct {
	Body *AgencyPlangetJSONRequestBody
}

type AgencyPlangetResponseObject interface {
	VisitAgencyPlangetResponse(w http.ResponseWriter) error
}

type AgencyPlanget200JSONResponse AgencyPlanGetResponse

func (response AgencyPlanget200JSONResponse) VisitAgencyPlangetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyWithPlanaddRequestObject struct {
	Body *AgencyWithPlanaddJSONRequestBody
}

type AgencyWithPlanaddResponseObject interface {
	VisitAgencyWithPlanaddResponse(w http.ResponseWriter) error
}

type AgencyWithPlanadd200JSONResponse AgencyWithPlanAddResponse

func (response AgencyWithPlanadd200JSONResponse) VisitAgencyWithPlanaddResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyWithPlanpagegetRequestObject struct {
	Body *AgencyWithPlanpagegetJSONRequestBody
}

type AgencyWithPlanpagegetResponseObject interface {
	VisitAgencyWithPlanpagegetResponse(w http.ResponseWriter) error
}

type AgencyWithPlanpageget200JSONResponse AgencyWithPlanPageResponse

func (response AgencyWithPlanpageget200JSONResponse) VisitAgencyWithPlanpagegetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AgencyWithPlangetRequestObject struct {
	Body *AgencyWithPlangetJSONRequestBody
}

type AgencyWithPlangetResponseObject interface {
	VisitAgencyWithPlangetResponse(w http.ResponseWriter) error
}

type AgencyWithPlanget200JSONResponse AgencyWithPlanGetResponse

func (response AgencyWithPlanget200JSONResponse) VisitAgencyWithPlangetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PlatformgetRequestObject struct {
	Body *PlatformgetJSONRequestBody
}

type PlatformgetResponseObject interface {
	VisitPlatformgetResponse(w http.ResponseWriter) error
}

type Platformget200JSONResponse PlatformGetResponse

func (response Platformget200JSONResponse) VisitPlatformgetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (PATCH /agency)
	AgencybatchUpdate(ctx context.Context, request AgencybatchUpdateRequestObject) (AgencybatchUpdateResponseObject, error)

	// (POST /agency)
	Agencyadd(ctx context.Context, request AgencyaddRequestObject) (AgencyaddResponseObject, error)

	// (POST /agency/pageGet)
	Agencypageget(ctx context.Context, request AgencypagegetRequestObject) (AgencypagegetResponseObject, error)

	// (PATCH /agency/{agencyId})
	Agencyupdate(ctx context.Context, request AgencyupdateRequestObject) (AgencyupdateResponseObject, error)

	// (POST /agencyGet)
	Agencyget(ctx context.Context, request AgencygetRequestObject) (AgencygetResponseObject, error)

	// (PATCH /agencyPlan)
	AgencyPlanbatchUpdate(ctx context.Context, request AgencyPlanbatchUpdateRequestObject) (AgencyPlanbatchUpdateResponseObject, error)

	// (POST /agencyPlan)
	AgencyPlanadd(ctx context.Context, request AgencyPlanaddRequestObject) (AgencyPlanaddResponseObject, error)

	// (POST /agencyPlan/byReferralCodeGet)
	AgencyPlangetByReferralCode(ctx context.Context, request AgencyPlangetByReferralCodeRequestObject) (AgencyPlangetByReferralCodeResponseObject, error)

	// (PATCH /agencyPlan/{agencyPlanId})
	AgencyPlanupdate(ctx context.Context, request AgencyPlanupdateRequestObject) (AgencyPlanupdateResponseObject, error)

	// (POST /agencyPlanGet)
	AgencyPlanget(ctx context.Context, request AgencyPlangetRequestObject) (AgencyPlangetResponseObject, error)

	// (POST /agencyWithPlan)
	AgencyWithPlanadd(ctx context.Context, request AgencyWithPlanaddRequestObject) (AgencyWithPlanaddResponseObject, error)

	// (POST /agencyWithPlan/pageGet)
	AgencyWithPlanpageget(ctx context.Context, request AgencyWithPlanpagegetRequestObject) (AgencyWithPlanpagegetResponseObject, error)

	// (POST /agencyWithPlanGet)
	AgencyWithPlanget(ctx context.Context, request AgencyWithPlangetRequestObject) (AgencyWithPlangetResponseObject, error)

	// (POST /platformGet)
	Platformget(ctx context.Context, request PlatformgetRequestObject) (PlatformgetResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// AgencybatchUpdate operation middleware
func (sh *strictHandler) AgencybatchUpdate(ctx *gin.Context) {
	var request AgencybatchUpdateRequestObject

	var body AgencybatchUpdateJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencybatchUpdate(ctx, request.(AgencybatchUpdateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencybatchUpdate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencybatchUpdateResponseObject); ok {
		if err := validResponse.VisitAgencybatchUpdateResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Agencyadd operation middleware
func (sh *strictHandler) Agencyadd(ctx *gin.Context) {
	var request AgencyaddRequestObject

	var body AgencyaddJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Agencyadd(ctx, request.(AgencyaddRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Agencyadd")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyaddResponseObject); ok {
		if err := validResponse.VisitAgencyaddResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Agencypageget operation middleware
func (sh *strictHandler) Agencypageget(ctx *gin.Context) {
	var request AgencypagegetRequestObject

	var body AgencypagegetJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Agencypageget(ctx, request.(AgencypagegetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Agencypageget")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencypagegetResponseObject); ok {
		if err := validResponse.VisitAgencypagegetResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Agencyupdate operation middleware
func (sh *strictHandler) Agencyupdate(ctx *gin.Context, agencyId int64) {
	var request AgencyupdateRequestObject

	request.AgencyId = agencyId

	var body AgencyupdateJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Agencyupdate(ctx, request.(AgencyupdateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Agencyupdate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyupdateResponseObject); ok {
		if err := validResponse.VisitAgencyupdateResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Agencyget operation middleware
func (sh *strictHandler) Agencyget(ctx *gin.Context) {
	var request AgencygetRequestObject

	var body AgencygetJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Agencyget(ctx, request.(AgencygetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Agencyget")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencygetResponseObject); ok {
		if err := validResponse.VisitAgencygetResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyPlanbatchUpdate operation middleware
func (sh *strictHandler) AgencyPlanbatchUpdate(ctx *gin.Context) {
	var request AgencyPlanbatchUpdateRequestObject

	var body AgencyPlanbatchUpdateJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyPlanbatchUpdate(ctx, request.(AgencyPlanbatchUpdateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyPlanbatchUpdate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyPlanbatchUpdateResponseObject); ok {
		if err := validResponse.VisitAgencyPlanbatchUpdateResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyPlanadd operation middleware
func (sh *strictHandler) AgencyPlanadd(ctx *gin.Context) {
	var request AgencyPlanaddRequestObject

	var body AgencyPlanaddJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyPlanadd(ctx, request.(AgencyPlanaddRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyPlanadd")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyPlanaddResponseObject); ok {
		if err := validResponse.VisitAgencyPlanaddResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyPlangetByReferralCode operation middleware
func (sh *strictHandler) AgencyPlangetByReferralCode(ctx *gin.Context) {
	var request AgencyPlangetByReferralCodeRequestObject

	var body AgencyPlangetByReferralCodeJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyPlangetByReferralCode(ctx, request.(AgencyPlangetByReferralCodeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyPlangetByReferralCode")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyPlangetByReferralCodeResponseObject); ok {
		if err := validResponse.VisitAgencyPlangetByReferralCodeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyPlanupdate operation middleware
func (sh *strictHandler) AgencyPlanupdate(ctx *gin.Context, agencyPlanId int64) {
	var request AgencyPlanupdateRequestObject

	request.AgencyPlanId = agencyPlanId

	var body AgencyPlanupdateJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyPlanupdate(ctx, request.(AgencyPlanupdateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyPlanupdate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyPlanupdateResponseObject); ok {
		if err := validResponse.VisitAgencyPlanupdateResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyPlanget operation middleware
func (sh *strictHandler) AgencyPlanget(ctx *gin.Context) {
	var request AgencyPlangetRequestObject

	var body AgencyPlangetJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyPlanget(ctx, request.(AgencyPlangetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyPlanget")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyPlangetResponseObject); ok {
		if err := validResponse.VisitAgencyPlangetResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyWithPlanadd operation middleware
func (sh *strictHandler) AgencyWithPlanadd(ctx *gin.Context) {
	var request AgencyWithPlanaddRequestObject

	var body AgencyWithPlanaddJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyWithPlanadd(ctx, request.(AgencyWithPlanaddRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyWithPlanadd")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyWithPlanaddResponseObject); ok {
		if err := validResponse.VisitAgencyWithPlanaddResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyWithPlanpageget operation middleware
func (sh *strictHandler) AgencyWithPlanpageget(ctx *gin.Context) {
	var request AgencyWithPlanpagegetRequestObject

	var body AgencyWithPlanpagegetJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyWithPlanpageget(ctx, request.(AgencyWithPlanpagegetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyWithPlanpageget")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyWithPlanpagegetResponseObject); ok {
		if err := validResponse.VisitAgencyWithPlanpagegetResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// AgencyWithPlanget operation middleware
func (sh *strictHandler) AgencyWithPlanget(ctx *gin.Context) {
	var request AgencyWithPlangetRequestObject

	var body AgencyWithPlangetJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AgencyWithPlanget(ctx, request.(AgencyWithPlangetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AgencyWithPlanget")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AgencyWithPlangetResponseObject); ok {
		if err := validResponse.VisitAgencyWithPlangetResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Platformget operation middleware
func (sh *strictHandler) Platformget(ctx *gin.Context) {
	var request PlatformgetRequestObject

	var body PlatformgetJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Platformget(ctx, request.(PlatformgetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Platformget")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PlatformgetResponseObject); ok {
		if err := validResponse.VisitPlatformgetResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
