package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator/v_data"
)

func (v BsValidator) VListenList(ctx *gin.Context) {
	var params v_data.VListenList
	bindData(ctx, &params, (&handlers.Listen{}).ListenList)
}

func (v BsValidator) VListenCreate(ctx *gin.Context) {
	var params v_data.VListenCreate
	bindData(ctx, &params, (&handlers.Listen{}).ListenCreate)
}

func (v BsValidator) VListenUpdate(ctx *gin.Context) {
	var params v_data.VListenUpdate
	bindData(ctx, &params, (&handlers.Listen{}).ListenUpdate)
}

func (v BsValidator) VListenInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Listen{}).ListenInfo)
}
