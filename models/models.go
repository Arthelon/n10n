package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func Init() error {
	var err error = nil
	db, err = gorm.Open("mysql", Conf.DbPath)
	if err != nil {
		return err
	}
	db.SingularTable(true)
	db.LogMode(true)
	return nil
}