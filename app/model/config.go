package model

import (
	"gorm.io/gorm"
	"market/app/cache"
)

type Config struct {
	connectDb

	Id     int64  `json:"id"`
	Key    string `json:"key" gorm:"column:_k"`
	Val    string `json:"val" gorm:"column:_v"`
	Desc   string `json:"desc" gorm:"column:_desc"`
	State  uint8  `json:"state"`
	Bak1   string `json:"bak1"`
	Bak2   string `json:"bak2"`
	Remark string `json:"remark"`
}

var (
	configKey = "db:configs:"
)

// NewConfig 实例
func NewConfig(db *gorm.DB) *Config {
	return &Config{connectDb: connectDb{DB: db}}
}

func (m *Config) TableName() string {
	return "configs"
}

func (m *Config) List(k, desc string, state uint8, offset, limit int64) (configs []*Config, total int64, err error) {
	query := m.Table(m.TableName()).Where("state = ?", state)
	if k != "" {
		query = query.Where("_k like ?", "%"+k+"%")
	}
	if desc != "" {
		query = query.Where("_desc like ?", "%"+desc+"%")
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	err = query.Offset(int(offset)).Limit(int(limit)).Find(&configs).Error
	return
}

func (m *Config) FindConfigsByKey(k string) (configs []*Config, err error) {
	err = cache.New(m.DB).QueryRow(configKey, &configs, k, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Where("_k = ? and state = 1", id).First(v).Error
	})
	return
}

func (m *Config) FindOneByKey(k string) (config *Config, err error) {
	err = m.Table(m.TableName()).Where("_k = ? and state = 1", k).First(&config).Error
	return
}

func (m *Config) FindOneById(id int64) (config *Config, err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).First(&config).Error
	return
}

func (m *Config) CreateConfig(c Config) (err error) {
	err = m.Table(m.TableName()).Create(&c).Error
	return
}

func (m *Config) UpdateConfig(id int64, v map[string]interface{}) (err error) {
	err = m.Table(m.TableName()).Where("id = ?", id).Updates(v).Error
	return
}

func (m *Config) ApiFindOneByKey(k string) (config *Config, err error) {
	err = cache.New(m.DB).QueryRow("db:config", &config, k, func(db *gorm.DB, v interface{}, id interface{}) error {
		return db.Table(m.TableName()).Select("_v").Where("_k = ?", id).First(v).Error
	})
	return
}
