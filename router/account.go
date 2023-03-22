package router

import (
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initAccountApis(g *gin.RouterGroup) {
	group := g.Group("/account", middleware.CheckUserLogin())
	{
		group.POST("/update", (validator.BsValidator{}).VAccountUpdate)

		group.GET("/:id", (validator.BsValidator{}).VAccountInfo)
		group.GET("/list", (validator.BsValidator{}).VAccountList)
	}
}
