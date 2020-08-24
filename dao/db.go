package dao

import (
	"awesomeProject/Log"
	"awesomeProject/config"
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
)

var (
	models = []interface{}{
		&model.Product{},
	}
)

func New() (db *gorm.DB, cf func(), err error) {
	db, err = gorm.Open("mysql", config.Env.DSN)
	if err != nil {
		Log.Info(err.Error())
	}
	cf = func() { defer db.Close() }
	return
}

// 检查如果表不存在则创建
func (dao *Db) CreateTable() {
	for _, model := range models {
		//if !dao.db.HasTable(model) {
		//	dao.db.CreateTable(model)
		//}
		dao.db.AutoMigrate(model)
	}
}
