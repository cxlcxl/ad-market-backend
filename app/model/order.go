package model

import (
	"gorm.io/gorm"
	"market/app/vars"
	"time"
)

type Order struct {
	connectDb

	Id         int64  `json:"id"`
	Amt        int    `json:"amt"`
	PayAt      string `json:"pay_at"`
	Mobile     string `json:"mobile"`
	OutTradeNo string `json:"out_trade_no"`
	State      uint8  `json:"state"`
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

func (m *Order) ActionOrder(outTradeNo string) (err error) {
	d := map[string]interface{}{
		"pay_at": time.Now().Format(vars.DateTimeFormat),
		"state":  vars.OrderStatePaid,
	}
	actSQL := "update orders left join accounts on orders.mobile = accounts.mobile set accounts.state = ? where orders.out_trade_no = ?"
	return m.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table(m.TableName()).Where("out_trade_no = ?", outTradeNo).UpdateColumns(d).Error; err != nil {
			return err
		}
		if err = tx.Exec(actSQL, vars.AccountStatePaid, outTradeNo).Error; err != nil {
			return err
		}
		return nil
	})
}
