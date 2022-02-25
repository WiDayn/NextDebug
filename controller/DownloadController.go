package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"prmlk.com/nextdebug/common"
)

type IDownloadController interface {
	Avatar(ctx *gin.Context)
}

type DownloadController struct {
	DB *gorm.DB
}

func (c DownloadController) Avatar(ctx *gin.Context) {
	//用于直接返回头像的webp
	avatarName := ctx.Query("src")
	file, _ := ioutil.ReadFile("./data/avatar/" + avatarName + ".webp")
	_, _ = ctx.Writer.WriteString(string(file))
}

func NewDownloadController() DownloadController {
	db := common.InitDB()

	return DownloadController{DB: db}
}
