package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/model"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	if len(name) < 3 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名不能小于三位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于六位"})
		return
	}

	if isNameExist(DB, name) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名已经被注册"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
	}
	newUser := model.User{
		Name:     name,
		Password: string(hashedPassword),
	}
	tx := DB.Create(&newUser)
	if tx.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": tx.Error})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 200, "msg": "注册成功"})
	}
	log.Println(name, password)
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	var user model.User

	DB.Where("name = ?", name).First(&user)

	if user.ID <= 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "用户名不存在"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": "400", "msg": "用户名或密码错误"})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error : %v\n", err)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登陆成功"})

	return
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}

func isNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	if user.ID > 0 {
		return true
	} else {
		return false
	}
}
