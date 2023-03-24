package router

import (
	"github.com/gin-gonic/gin"
	"market/app/middleware"
	"market/app/validator"
)

func initListenApis(g *gin.RouterGroup) {
	group := g.Group("/listen", middleware.CheckUserLogin())
	{
		group.POST("/create", (validator.BsValidator{}).VListenCreate)
		group.POST("/update", (validator.BsValidator{}).VListenUpdate)
		group.GET("/:id", (validator.BsValidator{}).VListenInfo)
		group.GET("/list", (validator.BsValidator{}).VListenList)
	}
}
