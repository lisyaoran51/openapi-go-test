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
	"github.com/gin-gonic/gin"
)

type AgencyPlanAPI interface {


    // AgencyPlanAdd Post /agencyPlan
     AgencyPlanAdd(c *gin.Context)

    // AgencyPlanBatchUpdate Patch /agencyPlan
     AgencyPlanBatchUpdate(c *gin.Context)

    // AgencyPlanGet Post /agencyPlanGet
     AgencyPlanGet(c *gin.Context)

    // AgencyPlanGetByReferralCode Post /agencyPlan/byReferralCodeGet
     AgencyPlanGetByReferralCode(c *gin.Context)

    // AgencyPlanUpdate Patch /agencyPlan/:agencyPlanId
     AgencyPlanUpdate(c *gin.Context)

}