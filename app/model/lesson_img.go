package model

import (
	"gorm.io/gorm"
	"market/app/cache"
)

type LessonImg struct {
	connectDb

	Id     int64  `json:"id"`
	Name   string `json:"name"`
	State  uint8  `json:"state"`
	ImgUrl string `json:"img_url"`
	FCode  string `json:"f_code"`
}

func (m *LessonImg) TableName() string {
	return "lesson_imgs"
}

func NewLessonImg(db *gorm.DB) *LessonImg {
	return &LessonImg{connectDb: connectDb{DB: db}}
}

func (m *LessonImg) LessonImgList(name string, offset, limit int64) (ls []*LessonImg, total int64, err error) {
	tbl := m.Table(m.TableName()).Where("state = 1")
	if len(name) > 0 {
		tbl = tbl.Where("name like ?", "%"+name+"%")
	}

	if err = tbl.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = tbl.Offset(int(offset)).Limit(int(limit)).Find(&ls).Error
	}
	return
}

func (m *LessonImg) LessonImgCreate(lesson *LessonImg) (err error) {
	err = m.Table(m.TableName()).Create(lesson).Error
	return
}

func (m *LessonImg) FindImgById(lessonId int64) (lists []*LessonImg) {
	m.Table(m.TableName()).Where("lesson_id = ?", lessonId).Order("order_by asc").Find(&lists)
	return
}

func (m *LessonImg) DeleteById(id int64) (err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).Update("state", 0).Error
	return
}

func (m *LessonImg) FindImgByCode(code string) (filePath string) {
	var img LessonImg
	err := cache.New(m.DB).QueryRow("db:lsimg", &img, code, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Select("img_url").Where("f_code = ?", id).First(v).Error
	})
	if err != nil {
		return ""
	}
	return img.ImgUrl
}
