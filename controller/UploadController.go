package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"prmlk.com/nextdebug/common"
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

func (c UploadController) Avatar(ctx *gin.Context) {
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
	//解码Base64，然后导出为WEBP，把图片名写入数据库
	unbased, err := base64.StdEncoding.DecodeString(baseCode)
	if err != nil {
		fmt.Println(err)
		response.Fail(ctx, nil, "Base64解码失败")
		return
	}

	webpName := strconv.FormatInt(time.Now().Unix(), 10)

	if err = ioutil.WriteFile("/data/avatar/"+webpName+".webp", unbased, 0666); err != nil {
		fmt.Println(err)
	}

	response.Success(ctx, gin.H{}, "上传成功")
}

func NewUploadController() IUploadController {
	db := common.InitDB()

	return UploadController{DB: db}
}
