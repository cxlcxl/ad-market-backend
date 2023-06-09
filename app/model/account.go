package model

import (
	"gorm.io/gorm"
	"market/app/vars"
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

func (m *Account) FindAccountStateByMobile(mobile string) (act *Account) {
	m.Table(m.TableName()).Where("mobile = ?", mobile).Select("state").First(&act)
	return
}

func (m *Account) FindAccountByMobile(mobile string) (ct int64) {
	m.Table(m.TableName()).Where("mobile = ?", mobile).Count(&ct)
	return
}

func (m *Account) AccountList(mobile, accountName string, state uint8, offset, limit int64) (ls []*Account, total int64, err error) {
	tbl := m.Table(m.TableName()).Order("id desc")
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
	/*if ct := m.FindAccountByMobile(mobile); ct > 0 {
		return nil
	}*/
	datetime := time.Now().Format(vars.DateTimeFormat)
	err = m.Exec("insert into accounts(mobile, state, created_at) values(?, ?, ?)", mobile, state, datetime).Error
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
