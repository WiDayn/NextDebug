package common

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"prmlk.com/nextdebug/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("Fail to read config:", err)
	}

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		cfg.Section("sql").Key("username").String(),
		cfg.Section("sql").Key("password").String(),
		cfg.Section("sql").Key("host").String(),
		cfg.Section("sql").Key("port").String(),
		cfg.Section("sql").Key("database").String(),
		cfg.Section("sql").Key("charset").String())

	if cfg.Section("sql").Key("driverName").String() == "mysql" {
		db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
		if err != nil {
			fmt.Println("Fail to connect a database: ", err)
		}
		err = db.AutoMigrate(&model.User{})
		if err != nil {
			fmt.Println("Fail to automigrate: ", err)
		}
		DB = db
		return db
	}
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
