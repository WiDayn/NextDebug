package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/model"
	"prmlk.com/nextdebug/response"
	"strconv"
)

type IOnlineJudgeController interface {
	RestController
}

type OnlineJudgeController struct {
	DB *gorm.DB
}

func NewOnlineJudgeController() IOnlineJudgeController {
	db := common.InitDB()
	err := db.AutoMigrate(model.OnlineJudge{})
	if err != nil {
		println("数据库初始化失败..")
		return nil
	}
	return OnlineJudgeController{DB: db}
}

func (c OnlineJudgeController) Create(ctx *gin.Context) {
	var requestOnlineJudge model.OnlineJudge
	err := ctx.Bind(&requestOnlineJudge)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}
	if requestOnlineJudge.Name == "" {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	c.DB.Create(&requestOnlineJudge)

	response.Success(ctx, gin.H{"problem": requestOnlineJudge}, "创建成功")
}

func (c OnlineJudgeController) Update(ctx *gin.Context) {
	var requestOnlineJudge model.OnlineJudge
	err := ctx.Bind(&requestOnlineJudge)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
	}
	if requestOnlineJudge.Name == "" {
		response.Fail(ctx, nil, "OJ名称不能为空")
		return
	}

	OnlineJudgeId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if OnlineJudgeId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var updateOnlineJudge model.OnlineJudge

	c.DB.First(&updateOnlineJudge, OnlineJudgeId)

	if updateOnlineJudge.ID == 0 {
		response.Fail(ctx, nil, "找不到该OJ")
		return
	}

	c.DB.Model(&updateOnlineJudge).Update("name", updateOnlineJudge.Name)

	response.Success(ctx, nil, "修改成功")
}

func (c OnlineJudgeController) Show(ctx *gin.Context) {
	OnlineJudgeId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if OnlineJudgeId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var OnlineJudge model.OnlineJudge

	c.DB.Where("ID = ?", OnlineJudgeId).First(&OnlineJudge)

	if OnlineJudge.ID == 0 {
		response.Fail(ctx, nil, "题目ID错误")
		return
	}

	response.Success(ctx, gin.H{"online_judge": OnlineJudge}, "获取成功")
}

func (c OnlineJudgeController) Delete(ctx *gin.Context) {
	OnlineJudgeId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if OnlineJudgeId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	if err := c.DB.Delete(&model.OnlineJudge{}, OnlineJudgeId); err.Error != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
