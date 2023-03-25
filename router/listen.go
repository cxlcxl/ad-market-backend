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
	asset := g.Group("/asset", middleware.CheckUserLogin())
	{
		asset.POST("/upload", (validator.BsValidator{}).VAssetUpload)
		asset.POST("/delete/:id", (validator.BsValidator{}).VAssetDel)
		asset.GET("/list", (validator.BsValidator{}).VAssetList)
	}
}
