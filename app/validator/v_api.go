package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator/v_data"
)

func (v BsValidator) VApiWxPayAction(ctx *gin.Context) {
	var params v_data.VApiWxPayAction
	// 解密
	bindData(ctx, &params, (&handlers.Payment{}).Action)
}

func (v BsValidator) VApiLogin(ctx *gin.Context) {
	var params v_data.VApiLogin
	bindData(ctx, &params, (&handlers.Api{}).Login)
}

func (v BsValidator) VApiOrder(ctx *gin.Context) {
	var params v_data.VApiOrder
	bindData(ctx, &params, (&handlers.Payment{}).Order)
}

func (v BsValidator) VApiOrderQuery(ctx *gin.Context) {
	var params v_data.VApiOrderQuery
	bindData(ctx, &params, (&handlers.Payment{}).OrderQuery)
}

func (v BsValidator) VAccountSms(ctx *gin.Context) {
	var params v_data.VAccountSms
	bindData(ctx, &params, (&handlers.Account{}).AccountSms)
}

func (v BsValidator) VAccountSmsValid(ctx *gin.Context) {
	var params v_data.VAccountSmsValid
	bindData(ctx, &params, (&handlers.Account{}).AccountSmsValid)
}
