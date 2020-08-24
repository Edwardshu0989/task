package dao

import (
	"awesomeProject/Log"
	"awesomeProject/config"
	"github.com/jinzhu/gorm"
)

func New() (db *gorm.DB, cf func(), err error) {
	db, err = gorm.Open("mysql", config.Env.DSN)
	if err != nil {
		Log.Info(err.Error())
	}
	cf = func() { defer db.Close() }
	return

}
