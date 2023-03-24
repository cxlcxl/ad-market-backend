package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator/v_data"
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

func (v BsValidator) VSelfUpdate(ctx *gin.Context) {
	var params v_data.VSelfUpdate
	bindData(ctx, &params, (&handlers.User{}).VSelfUpdate, fillUser)
}

func (v BsValidator) VUserInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.User{}).UserInfo)
}

func (v BsValidator) VLogin(ctx *gin.Context) {
	var params v_data.VLogin
	bindData(ctx, &params, (&handlers.User{}).Login)
}

func (v BsValidator) VResetPass(ctx *gin.Context) {
	var params v_data.VResetPass
	bindData(ctx, &params, (&handlers.User{}).ResetPass, fillUser)
}
