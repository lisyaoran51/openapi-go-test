// Package agencyBaseCompiled provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package agencyBaseCompiled

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
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