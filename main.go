package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	ID int64
	Name string `gorm:"type:varchar(20); not null"`
	Password string `gorm:"type:varchar(100); not null"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		password := ctx.PostForm("password")

		if len(name) == 0 {
			name = "Undefined"
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于六位"})
			return
		}

		newuser := User{
			Name: name,
			Password: password,
		}
		tx := db.Create(&newuser)
		if tx.Error != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg":tx.Error})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 200, "msg": "注册成功"})
		}
		log.Println(name, password)
	})
	r.Run()
}

func InitDB() *gorm.DB  {
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
		db.AutoMigrate(&User{})
		return db
	}
	return nil
}