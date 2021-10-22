package models

import (
	"fmt"
	"gin-blog/internal/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func InitDB() error {
	var err error
	db, err = gorm.Open(config.Database.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			config.Database.UserName,
			config.Database.Password,
			config.Database.Host,
			config.Database.DBName,
			config.Database.Charset,
			config.Database.ParseTime,
		))
	if err != nil {
		return err
	}

	if config.Server.RunMode == "debug" {
		db.LogMode(true)
	}
	//db.SingularTable(true)
	db.DB().SetMaxIdleConns(config.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.Database.MaxOpenConns)
	return nil
}

func NewQuery() *gorm.DB {
	return db
}
