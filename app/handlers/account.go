package handlers

import (
	"github.com/gin-gonic/gin"
	"market/app/model"
	"market/app/response"
	"market/app/service/ali_sms"
	servicepayment "market/app/service/payment"
	"market/app/utils"
	"market/app/validator/v_data"
	"market/app/vars"
	"strconv"
	"time"
)

type Account struct{}

func (h *Account) AccountList(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountList)
	offset := utils.GetPages(params.Page, params.PageSize)
	acts, total, err := model.NewAct(vars.DBMysql).AccountList(params.Mobile, params.AccountName, params.State, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"total": total,
		"list":  acts,
		"state": vars.AccountState,
	})
}

func (h *Account) AccountInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	act, err := model.NewAct(vars.DBMysql).FindAccountById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, act)
}

func (h *Account) AccountUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountUpdate)
	d := map[string]interface{}{
		"account_name": params.AccountName,
		"mobile":       params.Mobile,
		"state":        params.State,
		"remark":       params.Remark,
		"updated_at":   time.Now(),
	}
	if err := model.NewAct(vars.DBMysql).AccountUpdate(d, params.Id); err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (h *Account) AccountSms(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountSms)
	if _, err := ali_sms.BuildAndSend(params.Mobile); err != nil {
		response.Fail(ctx, "验证码发送失败："+err.Error())
	} else {
		response.Success(ctx, nil)
	}
}

func (h *Account) AccountSmsValid(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VAccountSmsValid)
	if len(params.Code) != 4 {
		response.Fail(ctx, "验证码错误")
		return
	}
	if err := ali_sms.ValidSmsCode(params.Mobile, params.Code); err != nil {
		response.Fail(ctx, "验证失败："+err.Error())
	} else {
		if act := model.NewAct(vars.DBMysql).FindAccountStateByMobile(params.Mobile); act != nil {
			if act.State >= vars.AccountStatePaid {
				response.Success(ctx, gin.H{"state": 2, "info": "已支付的手机号，请直接跳转"})
				return
			}
		} else {
			d := map[string]interface{}{"state": vars.AccountStateNoPaid}
			_ = model.NewAct(vars.DBMysql).AccountUpdateByMobile(d, params.Mobile)
		}
		// 验证通过直接支付下单
		rs, outTradeNo, err := servicepayment.UserOrder(ctx, params.Mobile)
		if err != nil {
			response.Fail(ctx, "下单失败请重试："+err.Error())
		} else {
			response.Success(ctx, gin.H{"state": 1, "info": "验证手机号成功，下单成功跳转支付", "order": rs, "order_sn": outTradeNo})
		}
	}
}
