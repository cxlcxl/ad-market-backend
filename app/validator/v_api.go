package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator/v_data"
)

func (v BsValidator) VApiLogin(ctx *gin.Context) {
	var params v_data.VApiLogin
	bindData(ctx, &params, (&handlers.Api{}).Login)
}
