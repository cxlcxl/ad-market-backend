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
	// header Wechatpay-Signature
	if err := servicepayment.ActionOrder(params); err == nil {
		response.PaySuccess(ctx)
	} else {
		response.PayFail(ctx)
	}
}

func (h *Payment) Order(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VApiOrder)
	rs, sn, err := servicepayment.UserOrder(ctx, params.Mobile)
	if err != nil {
		response.Fail(ctx, "下单失败："+err.Error())
	} else {
		response.Success(ctx, gin.H{"info": "下单成功跳转支付", "order": rs, "order_sn": sn})
	}
}

func (h *Payment) JsApiOrder(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VApiOrder)
	prepayId, sn, err := servicepayment.JsApiOrder(ctx, params.Mobile)
	if err != nil {
		response.Fail(ctx, "下单失败："+err.Error())
	} else {
		response.Success(ctx, gin.H{"info": "下单成功跳转支付", "prepay_id": prepayId, "order_sn": sn})
	}
}

func (h *Payment) OrderQuery(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VApiOrderQuery)
	state, err := servicepayment.OrderQuery(ctx, params.Sn)
	if err != nil {
		response.Fail(ctx, "订单查询失败："+err.Error())
	} else {
		response.Success(ctx, state)
	}
}
