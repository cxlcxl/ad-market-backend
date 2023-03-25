package model

import (
	"gorm.io/gorm"
	"market/app/cache"
)

type ListenImg struct {
	connectDb

	Id     int64  `json:"id"`
	Name   string `json:"name"`
	State  uint8  `json:"state"`
	ImgUrl string `json:"img_url"`
	FCode  string `json:"f_code"`
}

func (m *ListenImg) TableName() string {
	return "listen_imgs"
}

func NewListenImg(db *gorm.DB) *ListenImg {
	return &ListenImg{connectDb: connectDb{DB: db}}
}

func (m *ListenImg) ListenImgList(name string, offset, limit int64) (ls []*ListenImg, total int64, err error) {
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

func (m *ListenImg) ListenImgCreate(listen *ListenImg) (err error) {
	err = m.Table(m.TableName()).Create(listen).Error
	return
}

func (m *ListenImg) FindImgById(listenId int64) (lists []*ListenImg) {
	m.Table(m.TableName()).Where("listen_id = ?", listenId).Order("order_by asc").Find(&lists)
	return
}

func (m *ListenImg) DeleteById(id int64) (err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).Update("state", 0).Error
	return
}

func (m *ListenImg) FindImgByCode(code string) (filePath string) {
	var img ListenImg
	err := cache.New(m.DB).QueryRow("db:lsimg", &img, code, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Select("img_url").Where("f_code = ?", id).First(v).Error
	})
	if err != nil {
		return ""
	}
	return img.ImgUrl
}
