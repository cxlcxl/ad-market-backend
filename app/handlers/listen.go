package handlers

import (
	"github.com/gin-gonic/gin"
	"market/app/model"
	"market/app/response"
	"market/app/utils"
	"market/app/validator/v_data"
	"market/app/vars"
	"strconv"
)

type Listen struct{}

func (l *Listen) ListenList(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VListenList)
	offset := utils.GetPages(params.Page, params.PageSize)
	listens, total, err := model.NewListen(vars.DBMysql).ListenList(params.Title, params.State, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": listens})
}

func (l *Listen) ListenInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	listen, err := model.NewListen(vars.DBMysql).FindListenById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	listen.Lists = model.NewListenList(vars.DBMysql).FindListByListenId(id)
	response.Success(ctx, listen)
}

func (l *Listen) ListenCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VListenCreate)
	listen := &model.Listen{
		Title:    params.Title,
		ImgUrl:   params.ImgUrl,
		SubTitle: params.SubTitle,
		OrderBy:  params.OrderBy,
		Amt:      params.Amt,
		State:    1,
	}
	lists := make([]*model.ListenList, 0)
	for _, list := range params.Lists {
		lists = append(lists, &model.ListenList{Title: list.Title, OrderBy: list.OrderBy})
	}
	err := model.NewListen(vars.DBMysql).ListenCreate(listen, lists)
	if err != nil {
		response.Fail(ctx, "创建失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *Listen) ListenUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VListenUpdate)
	d := map[string]interface{}{
		"title":     params.Title,
		"order_by":  params.OrderBy,
		"state":     params.State,
		"sub_title": params.SubTitle,
		"img_url":   params.ImgUrl,
		"amt":       params.Amt,
	}
	lists := make([]*model.ListenList, 0)
	for _, list := range params.Lists {
		lists = append(lists, &model.ListenList{Title: list.Title, OrderBy: list.OrderBy, ListenId: params.Id})
	}
	err := model.NewListen(vars.DBMysql).ListenUpdate(d, params.Id, lists)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *Listen) Listen(ctx *gin.Context) {
	listens, err := model.NewListen(vars.DBMysql).ApiListenList()
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, listens)
}

func (l *Listen) ApiListenInfo(ctx *gin.Context) {
	v := ctx.Param("id")
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	listen, err := model.NewListen(vars.DBMysql).ApiFindListenById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	listen.Lists = model.NewListenList(vars.DBMysql).ApiFindListByListenId(id)
	response.Success(ctx, listen)
}
