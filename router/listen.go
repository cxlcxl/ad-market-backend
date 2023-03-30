package router

import (
	"github.com/gin-gonic/gin"
	"market/app/middleware"
	"market/app/validator"
)

func initLessonApis(g *gin.RouterGroup) {
	group := g.Group("/lesson", middleware.CheckUserLogin())
	{
		group.POST("/create", (validator.BsValidator{}).VLessonCreate)
		group.POST("/update", (validator.BsValidator{}).VLessonUpdate)
		group.GET("/:id", (validator.BsValidator{}).VLessonInfo)
		group.GET("/list", (validator.BsValidator{}).VLessonList)
	}
	asset := g.Group("/asset", middleware.CheckUserLogin())
	{
		asset.POST("/upload", (validator.BsValidator{}).VAssetUpload)
		asset.POST("/delete/:id", (validator.BsValidator{}).VAssetDel)
		asset.GET("/list", (validator.BsValidator{}).VAssetList)
	}
}
