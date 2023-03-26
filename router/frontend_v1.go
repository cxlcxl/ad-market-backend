package router

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator"
)

func initFrontV1Apis(g *gin.RouterGroup) {
	group := g.Group("/v1")
	{
		group.POST("/login", (validator.BsValidator{}).VApiLogin)

		group.POST("/sms", (validator.BsValidator{}).VAccountSms)
		group.POST("/sms-valid", (validator.BsValidator{}).VAccountSmsValid)

		group.GET("/asset/:code", (validator.BsValidator{}).VAsset)

		group.GET("/listen", (&handlers.Listen{}).Listen)
		group.GET("/listen/:id", (&handlers.Listen{}).ApiListenInfo)

		group.GET("/config", (&handlers.Config{}).ApiFindKey)
		group.GET("/configs", (&handlers.Config{}).ApiFindKeys)
	}
}
