package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"market/app/response"
	servicepayment "market/app/service/payment"
	"market/app/validator/v_data"
	"market/app/vars"
)

type Payment struct{}

func (h *Payment) Action(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VApiWxPayAction)
	apiV3Key := vars.YmlConfig.GetString("WxPay.ActionKey")
	fmt.Println(apiV3Key, params)
	if err := servicepayment.ActionOrder(""); err == nil {
		response.PaySuccess(ctx)
	} else {
		response.PayFail(ctx)
	}
}

func (h *Payment) Order(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VApiOrder)
	rs, err := servicepayment.UserOrder(ctx, params.Mobile)
	if err != nil {
		response.Fail(ctx, "下单失败："+err.Error())
	} else {
		response.Success(ctx, rs)
	}
}
