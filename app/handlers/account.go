package handlers

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/response"
	"bs.mobgi.cc/app/utils"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
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
		"state": vars.CommonState,
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
