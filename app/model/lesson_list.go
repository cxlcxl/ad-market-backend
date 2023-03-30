package model

import (
	"gorm.io/gorm"
	"market/app/cache"
)

type LessonList struct {
	connectDb

	Id       int64  `json:"id"`
	Title    string `json:"title"`
	OrderBy  int    `json:"order_by"`
	LessonId int64  `json:"lesson_id"`
}

func (m *LessonList) TableName() string {
	return "lesson_lists"
}

func NewLessonList(db *gorm.DB) *LessonList {
	return &LessonList{connectDb: connectDb{DB: db}}
}

func (m *LessonList) LessonListCreate(lessons []*LessonList) (err error) {
	err = m.Table(m.TableName()).CreateInBatches(lessons, 100).Error
	return
}

func (m *LessonList) FindListByLessonId(lessonId int64) (lists []*LessonList) {
	m.Table(m.TableName()).Where("lesson_id = ?", lessonId).Order("order_by asc").Find(&lists)
	return
}

func (m *LessonList) DeleteByLessonId(lessonId int64) (err error) {
	err = m.Exec("delete from lesson_lists where lesson_id = ?", lessonId).Error
	return
}

func (m *LessonList) ApiFindListByLessonId(lessonId int64) (lists []*LessonList) {
	_ = cache.New(m.DB).QueryRow("db:lessonlist", &lists, lessonId, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("lesson_id = ?", id).Order("order_by asc").Find(v).Error
	})
	return
}
