package model

import (
	"gorm.io/gorm"
	"market/app/cache"
)

type Lesson struct {
	connectDb

	Id       int64         `json:"id"`
	Title    string        `json:"title"`
	ImgUrl   string        `json:"img_url"`
	SubTitle string        `json:"sub_title"`
	OrderBy  int           `json:"order_by"`
	State    uint8         `json:"state"`
	Amt      int           `json:"amt"`
	Lists    []*LessonList `json:"lists" gorm:"-"`
}

func (m *Lesson) TableName() string {
	return "lessons"
}

func NewLesson(db *gorm.DB) *Lesson {
	return &Lesson{connectDb: connectDb{DB: db}}
}

func (m *Lesson) FindLessonById(id int64) (lsn *Lesson, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&lsn).Error
	return
}

func (m *Lesson) LessonList(title string, state uint8, offset, limit int64) (ls []*Lesson, total int64, err error) {
	tbl := m.Table(m.TableName()).Where("state = ?", state).Order("order_by asc, id desc")
	if len(title) > 0 {
		tbl = tbl.Where("title like ?", "%"+title+"%")
	}

	if err = tbl.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = tbl.Offset(int(offset)).Limit(int(limit)).Find(&ls).Error
	}
	return
}

func (m *Lesson) LessonCreate(lesson *Lesson, lists []*LessonList) (err error) {
	err = m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Create(lesson).Error; err != nil {
			return err
		}
		for i := range lists {
			lists[i].LessonId = lesson.Id
		}
		if err = NewLessonList(tx).LessonListCreate(lists); err != nil {
			return err
		}
		return nil
	})
	return
}

func (m *Lesson) LessonUpdate(d map[string]interface{}, id int64, lists []*LessonList) (err error) {
	err = m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Where("id = ?", id).Updates(d).Error; err != nil {
			return err
		}
		if err = NewLessonList(tx).DeleteByLessonId(id); err != nil {
			return err
		}
		if err = NewLessonList(tx).LessonListCreate(lists); err != nil {
			return err
		}
		return nil
	})

	return
}

///////// API ////////////

func (m *Lesson) ApiLessonList() (ls []*Lesson, err error) {
	err = cache.New(m.DB).Query("db:lessons", &ls, func(db *gorm.DB, v interface{}) error {
		return db.Table(m.TableName()).Where("state = 1").Select("id,title,sub_title,img_url").
			Order("order_by asc, id desc").Find(v).Error
	})

	return
}

func (m *Lesson) ApiFindLessonById(id int64) (lsn *Lesson, err error) {
	err = cache.New(m.DB).QueryRow("db:lesson", &lsn, id, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("id = ?", id).First(v).Error
	})
	return
}
