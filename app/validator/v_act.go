package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator/v_data"
)

func (v BsValidator) VAccountList(ctx *gin.Context) {
	var params v_data.VAccountList
	bindData(ctx, &params, (&handlers.Account{}).AccountList)
}

func (v BsValidator) VAccountUpdate(ctx *gin.Context) {
	var params v_data.VAccountUpdate
	bindData(ctx, &params, (&handlers.Account{}).AccountUpdate)
}

func (v BsValidator) VAccountInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Account{}).AccountInfo)
}
