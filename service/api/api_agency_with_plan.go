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

type AgencyWithPlanAPI interface {


    // AgencyWithPlanAdd Post /agencyWithPlan
     AgencyWithPlanAdd(c *gin.Context)

    // AgencyWithPlanGet Post /agencyWithPlanGet
     AgencyWithPlanGet(c *gin.Context)

    // AgencyWithPlanPageGet Post /agencyWithPlan/pageGet
     AgencyWithPlanPageGet(c *gin.Context)

}