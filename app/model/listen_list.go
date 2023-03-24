package model

import (
	"gorm.io/gorm"
)

type ListenList struct {
	connectDb

	Id       int64  `json:"id"`
	Title    string `json:"title"`
	OrderBy  int    `json:"order_by"`
	ListenId int64  `json:"listen_id"`
}

func (m *ListenList) TableName() string {
	return "listen_lists"
}

func NewListenList(db *gorm.DB) *ListenList {
	return &ListenList{connectDb: connectDb{DB: db}}
}

func (m *ListenList) ListenListCreate(listens []*ListenList) (err error) {
	err = m.Table(m.TableName()).CreateInBatches(listens, 100).Error
	return
}

func (m *ListenList) FindListByListenId(listenId int64) (lists []*ListenList) {
	m.Table(m.TableName()).Where("listen_id = ?", listenId).Order("order_by asc").Find(&lists)
	return
}

func (m *ListenList) DeleteByListenId(listenId int64) (err error) {
	err = m.Exec("delete from listen_lists where listen_id = ?", listenId).Error
	return
}
