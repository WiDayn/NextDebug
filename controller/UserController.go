package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/dto"
	"prmlk.com/nextdebug/model"
	"prmlk.com/nextdebug/response"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	if len(name) < 3 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能小于三位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于六位")
		return
	}

	if isNameExist(DB, name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已经被注册")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hashedPassword),
	}
	tx := DB.Create(&newUser)
	if tx.Error != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "注册失败")
	} else {
		response.Response(ctx, http.StatusOK, 200, nil, "注册成功")
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
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户名不存在")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户名或密码错误")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
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

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
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
