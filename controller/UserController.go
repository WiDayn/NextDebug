package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/model"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	if len(name) == 0 {
		name = "Undefined"
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于六位"})
		return
	}

	newUser := model.User{
		Name:     name,
		Password: password,
	}
	tx := DB.Create(&newUser)
	if tx.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": tx.Error})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 200, "msg": "注册成功"})
	}
	log.Println(name, password)
}
