package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/dto"
	"prmlk.com/nextdebug/model"
	"prmlk.com/nextdebug/response"
	"strconv"
	"strings"
	"time"
)

type IUploadController interface {
	Avatar(ctx *gin.Context)
}

type UploadController struct {
	DB *gorm.DB
}

func NewUploadController() IUploadController {
	db := common.InitDB()

	return UploadController{DB: db}
}

func (c UploadController) Avatar(ctx *gin.Context) {
	str, _ := ctx.Get("user")
	user := str.(model.User)

	var avatar model.Avatar
	err := ctx.Bind(&avatar)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}

	if avatar.Context == "" {
		response.Fail(ctx, nil, "图片不能为空")
		return
	}
	baseCode := avatar.Context[strings.IndexByte(avatar.Context, ',')+1:] //去除返回中的Base64解释头
	//解码Base64，然后导出为WEBP
	unbased, err := base64.StdEncoding.DecodeString(baseCode)
	if err != nil {
		fmt.Println(err)
		response.Fail(ctx, nil, "Base64解码失败")
		return
	}

	webpName := strconv.FormatInt(time.Now().Unix(), 10)

	if err = ioutil.WriteFile("./data/avatar/"+webpName+".webp", unbased, 0666); err != nil {
		fmt.Println(err)
	}
	//删除原来的头像文件
	_ = os.Remove("./data/avatar/" + user.Avatar + ".webp")

	//写入数据库，赋值新的头像名称并返回
	c.DB.Model(&user).Update("avatar", webpName)
	user.Avatar = webpName
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user)}, "上传成功")
}
