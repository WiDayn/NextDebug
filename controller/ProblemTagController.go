package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/model"
	"prmlk.com/nextdebug/response"
	"strconv"
)

type IProblemTagController interface {
	RestController
}

type ProblemTagController struct {
	DB *gorm.DB
}

func (p ProblemTagController) Create(ctx *gin.Context) {
	var requestProblemTag model.ProblemTag
	err := ctx.ShouldBind(&requestProblemTag)
	if err != nil {
		response.Fail(ctx, nil, "读取错误，请检查数据格式是否正确")
		return
	}
	if requestProblemTag.Name == "" {
		response.Fail(ctx, nil, "读取错误，Tag名称不能为空")
		return
	}
	p.DB.Save(&requestProblemTag)

	response.Success(ctx, gin.H{"problems_tag": requestProblemTag}, "创建成功")
}

func (p ProblemTagController) Update(ctx *gin.Context) {
	var requestProblemTag model.ProblemTag
	err := ctx.Bind(&requestProblemTag)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}
	if requestProblemTag.Name == "" {
		response.Fail(ctx, nil, "Tag标题不能为空")
		return
	}

	ProblemTagId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemTagId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var updateProblemTag model.ProblemTag

	p.DB.First(&updateProblemTag, ProblemTagId)

	if updateProblemTag.ID == 0 {
		response.Fail(ctx, nil, "找不到该题目")
		return
	}

	p.DB.Model(&updateProblemTag).Update("name", requestProblemTag.Name)

	response.Success(ctx, nil, "修改成功")
}

func (p ProblemTagController) Show(ctx *gin.Context) {
	ProblemTagId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemTagId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var ProblemTag model.ProblemTag

	p.DB.Where("ID = ?", ProblemTagId).First(&ProblemTag)
	if ProblemTag.ID == 0 {
		response.Fail(ctx, nil, "题目ID错误")
		return
	}
	p.DB.Model(&ProblemTag).Association("Problems").Find(&ProblemTag.Problems)
	response.Success(ctx, gin.H{"problem_tag": ProblemTag}, "获取成功")
}

func (p ProblemTagController) Delete(ctx *gin.Context) {
	ProblemTagId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemTagId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	if err := p.DB.Delete(&model.ProblemTag{}, ProblemTagId); err.Error != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}

	response.Success(ctx, nil, "删除成功")
}

func NewProblemTagController() IProblemTagController {
	db := common.InitDB()
	err := db.AutoMigrate(&model.ProblemTag{})
	if err != nil {
		println("数据库初始化失败..")
		return nil
	}
	return ProblemTagController{DB: db}
}
