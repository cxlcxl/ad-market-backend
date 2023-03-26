package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator/v_data"
)

func (v BsValidator) VConfigList(ctx *gin.Context) {
	var p v_data.VConfigList
	bindData(ctx, &p, (&handlers.Config{}).Configs)
}

func (v BsValidator) VConfigCreate(ctx *gin.Context) {
	var p v_data.VConfigCreate
	bindData(ctx, &p, (&handlers.Config{}).ConfigCreate)
}

func (v BsValidator) VConfig(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Config{}).Config)
}

func (v BsValidator) VConfigUpdate(ctx *gin.Context) {
	var p v_data.VConfigUpdate
	bindData(ctx, &p, (&handlers.Config{}).ConfigUpdate)
}
