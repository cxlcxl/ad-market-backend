package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/validator/v_data"
)

func (v BsValidator) VLessonList(ctx *gin.Context) {
	var params v_data.VLessonList
	bindData(ctx, &params, (&handlers.Lesson{}).LessonList)
}

func (v BsValidator) VLessonCreate(ctx *gin.Context) {
	var params v_data.VLessonCreate
	bindData(ctx, &params, (&handlers.Lesson{}).LessonCreate)
}

func (v BsValidator) VLessonUpdate(ctx *gin.Context) {
	var params v_data.VLessonUpdate
	bindData(ctx, &params, (&handlers.Lesson{}).LessonUpdate)
}

func (v BsValidator) VLessonInfo(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Lesson{}).LessonInfo)
}
