package router

import (
	"github.com/gin-gonic/gin"
	"market/app/middleware"
	"market/app/validator"
)

func initConfigApis(g *gin.RouterGroup) {
	group := g.Group("/config", middleware.CheckUserLogin())
	{
		group.POST("/update", (validator.BsValidator{}).VConfigUpdate)
		group.POST("/create", (validator.BsValidator{}).VConfigCreate)

		group.GET("/:id", (validator.BsValidator{}).VConfig)
		group.GET("/list", (validator.BsValidator{}).VConfigList)
	}
}
