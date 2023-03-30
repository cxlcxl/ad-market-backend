package router

import (
	"market/app/handlers"
	"market/app/middleware"
	"market/app/validator"

	"github.com/gin-gonic/gin"
)

func initFrontV1Apis(g *gin.RouterGroup) {
	g.GET("/v1/asset/:code", (validator.BsValidator{}).VAsset)
	group := g.Group("/v1", middleware.CheckApiSecret()) //
	{
		group.POST("/pay", (validator.BsValidator{}).VApiOrder)

		group.POST("/login", (validator.BsValidator{}).VApiLogin)

		group.POST("/sms", (validator.BsValidator{}).VAccountSms)
		group.POST("/sms-valid", (validator.BsValidator{}).VAccountSmsValid)

		group.GET("/listen", (&handlers.Listen{}).Listen)
		group.GET("/listen/:id", (&handlers.Listen{}).ApiListenInfo)

		group.GET("/config", (&handlers.Config{}).ApiFindKey)
		group.GET("/configs", (&handlers.Config{}).ApiFindKeys)
	}
}
