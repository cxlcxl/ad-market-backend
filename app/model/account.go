package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	connectDb

	Id          int64     `json:"id"`
	State       uint8     `json:"state"`        // 状态
	AccountName string    `json:"account_name"` // 账户名
	Mobile      string    `json:"mobile"`       //
	Remark      string    `json:"remark"`       //
	CreatedAt   time.Time `json:"created_at"`   // 添加时间
	UpdatedAt   time.Time `json:"updated_at"`   // 最后一次修改时间
}

func (m *Account) TableName() string {
	return "accounts"
}

func NewAct(db *gorm.DB) *Account {
	return &Account{connectDb: connectDb{DB: db}}
}

func (m *Account) FindAccountById(id int64) (act *Account, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&act).Error
	return
}

func (m *Account) AccountList(mobile, accountName string, state uint8, offset, limit int64) (ls []*Account, total int64, err error) {
	tbl := m.Table(m.TableName()).Order("updated_at desc")
	if len(accountName) > 0 {
		tbl = tbl.Where("account_name like ?", "%"+accountName+"%")
	}
	if mobile != "" {
		tbl = tbl.Where("mobile = ?", mobile)
	}
	if state > 0 {
		tbl = tbl.Where("state = ?", state)
	}
	if err = tbl.Count(&total).Error; err != nil {
		return
	}
	if total > 0 {
		err = tbl.Offset(int(offset)).Limit(int(limit)).Find(&ls).Error
	}
	return
}

func (m *Account) AccountCreate(state int, mobile string) (err error) {
	err = m.Exec("insert ignore into accounts(mobile, state, created_at) values(?, ?, NOW())", mobile, state).Error
	return
}

func (m *Account) AccountUpdate(d map[string]interface{}, id int64) (err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).Updates(d).Error
	return
}

func (m *Account) AccountUpdateByMobile(d map[string]interface{}, mobile string) (err error) {
	err = m.Table(m.TableName()).Where("mobile = ?", mobile).Updates(d).Error
	return
}
