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
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the AgencyAPI part of the API
	AgencyAPI AgencyAPI
	// Routes for the AgencyPlanAPI part of the API
	AgencyPlanAPI AgencyPlanAPI
	// Routes for the AgencyWithPlanAPI part of the API
	AgencyWithPlanAPI AgencyWithPlanAPI
	// Routes for the PlatformAPI part of the API
	PlatformAPI PlatformAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		{
			"AgencyAdd",
			http.MethodPost,
			"/agency",
			handleFunctions.AgencyAPI.AgencyAdd,
		},
		{
			"AgencyBatchUpdate",
			http.MethodPatch,
			"/agency",
			handleFunctions.AgencyAPI.AgencyBatchUpdate,
		},
		{
			"AgencyGet",
			http.MethodPost,
			"/agencyGet",
			handleFunctions.AgencyAPI.AgencyGet,
		},
		{
			"AgencyPageGet",
			http.MethodPost,
			"/agency/pageGet",
			handleFunctions.AgencyAPI.AgencyPageGet,
		},
		{
			"AgencyUpdate",
			http.MethodPatch,
			"/agency/:agencyId",
			handleFunctions.AgencyAPI.AgencyUpdate,
		},
		{
			"AgencyPlanAdd",
			http.MethodPost,
			"/agencyPlan",
			handleFunctions.AgencyPlanAPI.AgencyPlanAdd,
		},
		{
			"AgencyPlanBatchUpdate",
			http.MethodPatch,
			"/agencyPlan",
			handleFunctions.AgencyPlanAPI.AgencyPlanBatchUpdate,
		},
		{
			"AgencyPlanGet",
			http.MethodPost,
			"/agencyPlanGet",
			handleFunctions.AgencyPlanAPI.AgencyPlanGet,
		},
		{
			"AgencyPlanGetByReferralCode",
			http.MethodPost,
			"/agencyPlan/byReferralCodeGet",
			handleFunctions.AgencyPlanAPI.AgencyPlanGetByReferralCode,
		},
		{
			"AgencyPlanUpdate",
			http.MethodPatch,
			"/agencyPlan/:agencyPlanId",
			handleFunctions.AgencyPlanAPI.AgencyPlanUpdate,
		},
		{
			"AgencyWithPlanAdd",
			http.MethodPost,
			"/agencyWithPlan",
			handleFunctions.AgencyWithPlanAPI.AgencyWithPlanAdd,
		},
		{
			"AgencyWithPlanGet",
			http.MethodPost,
			"/agencyWithPlanGet",
			handleFunctions.AgencyWithPlanAPI.AgencyWithPlanGet,
		},
		{
			"AgencyWithPlanPageGet",
			http.MethodPost,
			"/agencyWithPlan/pageGet",
			handleFunctions.AgencyWithPlanAPI.AgencyWithPlanPageGet,
		},
		{
			"PlatformGet",
			http.MethodPost,
			"/platformGet",
			handleFunctions.PlatformAPI.PlatformGet,
		},
	}
}
