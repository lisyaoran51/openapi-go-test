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

type PlatformAPI interface {


    // PlatformGet Post /platformGet
     PlatformGet(c *gin.Context)

}