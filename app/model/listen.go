package model

import (
	"gorm.io/gorm"
	"market/app/cache"
)

type Listen struct {
	connectDb

	Id       int64         `json:"id"`
	Title    string        `json:"title"`
	ImgUrl   string        `json:"img_url"`
	SubTitle string        `json:"sub_title"`
	OrderBy  int           `json:"order_by"`
	State    uint8         `json:"state"`
	Amt      int           `json:"amt"`
	Lists    []*ListenList `json:"lists" gorm:"-"`
}

func (m *Listen) TableName() string {
	return "listens"
}

func NewListen(db *gorm.DB) *Listen {
	return &Listen{connectDb: connectDb{DB: db}}
}

func (m *Listen) FindListenById(id int64) (lsn *Listen, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&lsn).Error
	return
}

func (m *Listen) ListenList(title string, state uint8, offset, limit int64) (ls []*Listen, total int64, err error) {
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

func (m *Listen) ListenCreate(listen *Listen, lists []*ListenList) (err error) {
	err = m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Create(listen).Error; err != nil {
			return err
		}
		for i := range lists {
			lists[i].ListenId = listen.Id
		}
		if err = NewListenList(tx).ListenListCreate(lists); err != nil {
			return err
		}
		return nil
	})
	return
}

func (m *Listen) ListenUpdate(d map[string]interface{}, id int64, lists []*ListenList) (err error) {
	err = m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Where("id = ?", id).Updates(d).Error; err != nil {
			return err
		}
		if err = NewListenList(tx).DeleteByListenId(id); err != nil {
			return err
		}
		if err = NewListenList(tx).ListenListCreate(lists); err != nil {
			return err
		}
		return nil
	})

	return
}

///////// API ////////////

func (m *Listen) ApiListenList() (ls []*Listen, err error) {
	err = cache.New(m.DB).Query("db:listens", &ls, func(db *gorm.DB, v interface{}) error {
		return db.Table(m.TableName()).Where("state = 1").Select("id,title,sub_title,img_url").
			Order("order_by asc, id desc").Find(v).Error
	})

	return
}

func (m *Listen) ApiFindListenById(id int64) (lsn *Listen, err error) {
	err = cache.New(m.DB).QueryRow("db:listen", &lsn, id, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("id = ?", id).First(v).Error
	})
	return
}
