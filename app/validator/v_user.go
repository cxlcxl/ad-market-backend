package validator

import (
	"bs.mobgi.cc/app/handlers"
	"bs.mobgi.cc/app/validator/v_data"
	"github.com/gin-gonic/gin"
)

func (v BsValidator) VUserList(ctx *gin.Context) {
	var params v_data.VUserList
	bindData(ctx, &params, (&handlers.User{}).UserList)
}

func (v BsValidator) VUserCreate(ctx *gin.Context) {
	var params v_data.VUserCreate
	bindData(ctx, &params, (&handlers.User{}).UserCreate)
}

func (v BsValidator) VUserUpdate(ctx *gin.Context) {
	var params v_data.VUserUpdate
	bindData(ctx, &params, (&handlers.User{}).UserUpdate)
}

func (v BsValidator) VUserInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.User{}).UserInfo)
}

func (v BsValidator) VLogin(ctx *gin.Context) {
	var params v_data.VLogin
	bindData(ctx, &params, (&handlers.User{}).Login)
}
