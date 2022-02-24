package common

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"prmlk.com/nextdebug/model"
	"time"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("Fail to read config:", err)
	}

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		cfg.Section("sql").Key("username").String(),
		cfg.Section("sql").Key("password").String(),
		cfg.Section("sql").Key("host").String(),
		cfg.Section("sql").Key("port").String(),
		cfg.Section("sql").Key("database").String(),
		cfg.Section("sql").Key("charset").String(),
		cfg.Section("sql").Key("loc").String())

	if cfg.Section("sql").Key("driverName").String() == "mysql" {
		db, _ := gorm.Open(mysql.Open(args), &gorm.Config{})
		if err != nil {
			fmt.Println("Fail to connect a database: ", err)
		}
		err = db.AutoMigrate(&model.User{})
		if err != nil {
			fmt.Println("Fail to automigrate: ", err)
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
		DB = db
		return db
	}
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
