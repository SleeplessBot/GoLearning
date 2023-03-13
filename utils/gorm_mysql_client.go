package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Address string // eg. root:root@tcp(my_db:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local
	Debug   bool
}

func NewDbCli(conf DbConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(
		mysql.Open(conf.Address),
		&gorm.Config{
			Logger: NewGoZeroGormLogger(),
		})
	if err != nil {
		return
	}
	if conf.Debug {
		db = db.Debug()
	}
	return
}
