package model

import (
	"gorm.io/gorm"
)

type Order struct {
	connectDb

	Id         int64  `json:"id"`
	Amt        int    `json:"amt"`
	PayAt      string `json:"pay_at"`
	Mobile     string `json:"mobile"`
	OutTradeNo string `json:"out_trade_no"`
}

// NewOrder 实例
func NewOrder(db *gorm.DB) *Order {
	return &Order{connectDb: connectDb{DB: db}}
}

func (m *Order) TableName() string {
	return "orders"
}

func (m *Order) CreateOrder(order *Order) error {
	return m.Table(m.TableName()).Create(order).Error
}
