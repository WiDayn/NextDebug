package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/dto"
	"prmlk.com/nextdebug/model"
	"prmlk.com/nextdebug/response"
	"strconv"
)

type IProblemController interface {
	RestController
}

type ProblemController struct {
	DB *gorm.DB
}

func NewProblemController() IProblemController {
	db := common.InitDB()
	err := db.AutoMigrate(&model.Problem{})
	if err != nil {
		println("数据库初始化失败..")
		return nil
	}
	if err != nil {
		println("数据库初始化失败..")
		return nil
	}
	return ProblemController{DB: db}
}

func (c ProblemController) Create(ctx *gin.Context) {
	str, _ := ctx.Get("user")
	user := str.(model.User)

	var requestProblem model.Problem
	err := ctx.Bind(&requestProblem)

	requestProblem.Uploader = int(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		response.Fail(ctx, nil, "读取错误，请检查数据格式是否正确")
		return
	}
	if requestProblem.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，题目名称不能为空")
		return
	}

	c.DB.Save(&requestProblem)

	response.Success(ctx, gin.H{"problem": requestProblem}, "创建成功")
}

func (c ProblemController) Update(ctx *gin.Context) {
	var requestProblem model.Problem
	err := ctx.Bind(&requestProblem)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}
	if requestProblem.Name == "" {
		response.Fail(ctx, nil, "题目标题不能为空")
		return
	}

	ProblemId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var updateProblem model.Problem

	c.DB.First(&updateProblem, ProblemId)

	if updateProblem.ID == 0 {
		response.Fail(ctx, nil, "找不到该题目")
		return
	}
	c.DB.Model(&updateProblem).Update("name", requestProblem.Name)
	//TAG
	err = c.DB.Model(&updateProblem).Association("ProblemTag").Clear()
	if err != nil {
		response.Fail(ctx, nil, "关联的Tag清除失败")
		return
	}
	err = c.DB.Save(&updateProblem).Association("ProblemTag").Append(requestProblem.ProblemTag)
	if err != nil {
		response.Fail(ctx, nil, "Tag修改失败")
		return
	}
	//List
	err = c.DB.Model(&updateProblem).Association("ProblemList").Clear()
	if err != nil {
		response.Fail(ctx, nil, "关联的题单清除失败")
		return
	}
	err = c.DB.Save(&updateProblem).Association("ProblemList").Append(requestProblem.ProblemList)
	if err != nil {
		response.Fail(ctx, nil, "题单修改失败")
		return
	}
	//Related
	err = c.DB.Model(&updateProblem).Association("RelatedProblem").Clear()
	if err != nil {
		response.Fail(ctx, nil, "关联的问题清除失败")
		return
	}
	err = c.DB.Save(&updateProblem).Association("RelatedProblem").Append(requestProblem.RelatedProblem)
	if err != nil {
		response.Fail(ctx, nil, "问题修改失败")
		return
	}

	response.Success(ctx, nil, "修改成功")
}

func (c ProblemController) Show(ctx *gin.Context) {
	ProblemId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var problem model.Problem

	c.DB.Where("ID = ?", ProblemId).First(&problem)
	c.DB.Model(&problem).Association("ProblemTag").Find(&problem.ProblemTag)
	c.DB.Model(&problem).Association("ProblemList").Find(&problem.ProblemList)
	c.DB.Model(&problem).Association("RelatedProblem").Find(&problem.RelatedProblem)
	if problem.ID == 0 {
		response.Fail(ctx, nil, "题目ID错误")
		return
	}

	response.Success(ctx, gin.H{"problem": dto.ToProblemDetailDto(problem)}, "获取成功")
}

func (c ProblemController) Delete(ctx *gin.Context) {
	ProblemId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if ProblemId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	if err := c.DB.Delete(&model.Problem{}, ProblemId); err.Error != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
