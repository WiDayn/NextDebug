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
	"strings"
)

type IUserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Info(ctx *gin.Context)
}

type UserController struct {
	DB *gorm.DB
}

func NewUserController() UserController {
	db := common.InitDB()

	return UserController{DB: db}
}

func (c UserController) Register(ctx *gin.Context) {
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)

	requestUser.NickName = requestUser.Name
	requestUser.Name = strings.ToLower(requestUser.Name)
	requestUser.Email = strings.ToLower(requestUser.Email)

	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户信息读取错误")
		return
	}

	if len(requestUser.Name) < 3 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能小于三位")
		return
	}

	if len(requestUser.Password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于六位")
		return
	}

	if isNameExist(c.DB, requestUser.Name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已经被注册")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Email:    requestUser.Email,
		Name:     requestUser.Name,
		NickName: requestUser.NickName,
		Password: string(hashedPassword),
	}
	tx := c.DB.Create(&newUser)
	if tx.Error != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "注册信息写入失败")
		return
	}
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v\n", err)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "注册成功"})
}

func (c UserController) Login(ctx *gin.Context) {

	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)

	requestUser.Name = strings.ToLower(requestUser.Name)
	requestUser.Email = strings.ToLower(requestUser.Email)

	var user model.User

	c.DB.Where("name = ?", requestUser.Name).First(&user)

	if user.ID <= 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户名不存在")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password))

	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户名或密码错误")
		return
	}
	//发放token
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

func (c UserController) Info(ctx *gin.Context) {
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
