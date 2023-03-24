package router

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/middleware"
	"market/app/validator"
)

func initRbacApis(r *gin.RouterGroup) {
	r.POST("/login", (validator.BsValidator{}).VLogin) //用户登陆

	r.Use(middleware.CheckUserLogin())
	{
		r.POST("/self/update", (validator.BsValidator{}).VSelfUpdate)
		r.GET("/profile", (&handlers.User{}).Profile) //个人信息
		r.POST("/logout", (&handlers.User{}).Logout)
		r.POST("/reset-pass", (validator.BsValidator{}).VResetPass)

		u := r.Group("/user")
		{
			//角色列表
			u.GET("/list", (validator.BsValidator{}).VUserList)
			//用户创建
			u.POST("/create", (validator.BsValidator{}).VUserCreate)
			//用户修改
			u.POST("/update", (validator.BsValidator{}).VUserUpdate)
			//用户信息
			u.GET("/:id", (validator.BsValidator{}).VUserInfo)
			//用户删除
			u.POST("/destroy")
		}
	}
}
