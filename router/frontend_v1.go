package router

import (
	"market/app/handlers"
	"market/app/middleware"
	"market/app/validator"

	"github.com/gin-gonic/gin"
)

func initFrontV1Apis(g *gin.RouterGroup) {
	g.POST("/wxpay/action", (validator.BsValidator{}).VApiWxPayAction)
	g.GET("/v1/asset/:code", (validator.BsValidator{}).VAsset)
	//g.GET("/v1/urllink", (&handlers.Api{}).GetUrlLink)

	group := g.Group("/v1", middleware.CheckApiSecret()) //
	{
		group.POST("/pay", (validator.BsValidator{}).VApiOrder)
		//group.POST("/pay", (validator.BsValidator{}).VApiJsApiOrder)
		group.POST("/query-order", (validator.BsValidator{}).VApiOrderQuery)

		group.POST("/login", (validator.BsValidator{}).VApiLogin)

		group.POST("/sms", (validator.BsValidator{}).VAccountSms)
		group.POST("/sms-valid", (validator.BsValidator{}).VAccountSmsValid)

		group.GET("/lesson", (&handlers.Lesson{}).Lesson)
		group.GET("/lesson/:id", (&handlers.Lesson{}).ApiLessonInfo)

		group.GET("/config", (&handlers.Config{}).ApiFindKey)
		group.GET("/configs", (&handlers.Config{}).ApiFindKeys)

		group.POST("/xcx-sdk", (&handlers.Api{}).XcxSdk)
		group.GET("/urllink", (&handlers.Api{}).GetUrlLink)
	}
}
