package router

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/middleware"
	"bs.mobgi.cc/app/validator"
	"github.com/gin-gonic/gin"
)

func initRbacApis(r *gin.RouterGroup) {
	r.POST("/login", (validator.BsValidator{}).VLogin) //用户登陆

	r.Use(middleware.CheckUserLogin())
	{
		r.GET("/profile", (&handlers.User{}).Profile) //个人信息
		r.POST("/logout", (&handlers.User{}).Logout)

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
