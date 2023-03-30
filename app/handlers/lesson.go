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

type Lesson struct{}

func (l *Lesson) LessonList(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VLessonList)
	offset := utils.GetPages(params.Page, params.PageSize)
	lessons, total, err := model.NewLesson(vars.DBMysql).LessonList(params.Title, params.State, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": lessons})
}

func (l *Lesson) LessonInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	lesson, err := model.NewLesson(vars.DBMysql).FindLessonById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	lesson.Lists = model.NewLessonList(vars.DBMysql).FindListByLessonId(id)
	response.Success(ctx, lesson)
}

func (l *Lesson) LessonCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VLessonCreate)
	lesson := &model.Lesson{
		Title:    params.Title,
		ImgUrl:   params.ImgUrl,
		SubTitle: params.SubTitle,
		OrderBy:  params.OrderBy,
		Amt:      params.Amt,
		State:    1,
	}
	lists := make([]*model.LessonList, 0)
	for _, list := range params.Lists {
		lists = append(lists, &model.LessonList{Title: list.Title, OrderBy: list.OrderBy})
	}
	err := model.NewLesson(vars.DBMysql).LessonCreate(lesson, lists)
	if err != nil {
		response.Fail(ctx, "创建失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *Lesson) LessonUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VLessonUpdate)
	d := map[string]interface{}{
		"title":     params.Title,
		"order_by":  params.OrderBy,
		"state":     params.State,
		"sub_title": params.SubTitle,
		"img_url":   params.ImgUrl,
		"amt":       params.Amt,
	}
	lists := make([]*model.LessonList, 0)
	for _, list := range params.Lists {
		lists = append(lists, &model.LessonList{Title: list.Title, OrderBy: list.OrderBy, LessonId: params.Id})
	}
	err := model.NewLesson(vars.DBMysql).LessonUpdate(d, params.Id, lists)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *Lesson) Lesson(ctx *gin.Context) {
	lessons, err := model.NewLesson(vars.DBMysql).ApiLessonList()
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, lessons)
}

func (l *Lesson) ApiLessonInfo(ctx *gin.Context) {
	v := ctx.Param("id")
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	lesson, err := model.NewLesson(vars.DBMysql).ApiFindLessonById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	lesson.Lists = model.NewLessonList(vars.DBMysql).ApiFindListByLessonId(id)
	response.Success(ctx, lesson)
}
