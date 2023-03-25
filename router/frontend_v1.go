package router

import (
	"github.com/gin-gonic/gin"
	"market/app/validator"
)

func initFrontV1Apis(g *gin.RouterGroup) {
	group := g.Group("/v1")
	{
		group.POST("/sms", (validator.BsValidator{}).VAccountSms)
		group.POST("/sms-valid", (validator.BsValidator{}).VAccountSmsValid)

		group.GET("/asset/:code", (validator.BsValidator{}).VAsset)
	}
}
