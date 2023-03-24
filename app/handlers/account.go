package handlers

import (
	"github.com/gin-gonic/gin"
	"market/app/model"
	"market/app/response"
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
