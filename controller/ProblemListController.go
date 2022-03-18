package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/model"
	"prmlk.com/nextdebug/response"
	"strconv"
)

type IProblemListController interface {
	RestController
}

type ProblemListController struct {
	DB *gorm.DB
}

func (p ProblemListController) Create(ctx *gin.Context) {
	str, _ := ctx.Get("user")
	user := str.(model.User)

	var requestProblemList model.ProblemList
	err := ctx.Bind(&requestProblemList)

	requestProblemList.Creator = int(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		response.Fail(ctx, nil, "读取错误，请检查数据格式是否正确")
		return
	}
	if requestProblemList.Problems == nil {
		response.Fail(ctx, nil, "数据验证错误，题目名称不能为空")
		return
	}
	requestProblemList.Vote = 0
	p.DB.Create(&requestProblemList)
	response.Success(ctx, gin.H{"problem": requestProblemList}, "创建成功")
}

func (p ProblemListController) Update(ctx *gin.Context) {
	var requestProblemList model.ProblemList
	err := ctx.Bind(&requestProblemList)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}
	if requestProblemList.Name == "" {
		response.Fail(ctx, nil, "题单标题不能为空")
		return
	}

	ProblemListId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemListId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var updateProblemList model.ProblemList

	p.DB.First(&updateProblemList, ProblemListId)

	if updateProblemList.ID == 0 {
		response.Fail(ctx, nil, "找不到该题目")
		return
	}

	p.DB.Model(&updateProblemList).Update("name", updateProblemList.Name)
	//Problem
	err = p.DB.Model(&updateProblemList).Association("Problems").Clear()
	if err != nil {
		response.Fail(ctx, nil, "关联的题目清除失败")
		return
	}
	err = p.DB.Save(&updateProblemList).Association("Problems").Append(requestProblemList.Problems)
	if err != nil {
		response.Fail(ctx, nil, "题目修改失败")
		return
	}

	response.Success(ctx, nil, "修改成功")
}

func (p ProblemListController) Show(ctx *gin.Context) {
	ProblemListId, _ := strconv.Atoi(ctx.Params.ByName("originalID"))

	if ProblemListId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var ProblemList model.ProblemList

	p.DB.Where("original_id = ?", ProblemListId).First(&ProblemList)
	if ProblemList.ID == 0 {
		response.Fail(ctx, nil, "题单ID错误")
		return
	}
	p.DB.Model(&ProblemList).Association("Problems").Find(&ProblemList.Problems)
	response.Success(ctx, gin.H{"problem_tag": ProblemList}, "获取成功")
}

func (p ProblemListController) Delete(ctx *gin.Context) {
	ProblemListId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemListId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	if err := p.DB.Delete(&model.ProblemTag{}, ProblemListId); err.Error != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}

	response.Success(ctx, nil, "删除成功")
}

func NewProblemListController() IProblemListController {
	db := common.InitDB()
	err := db.AutoMigrate(model.ProblemList{})
	if err != nil {
		println("数据库初始化失败..")
		return nil
	}
	return ProblemListController{DB: db}
}
