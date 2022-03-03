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

type ITestSetController interface {
	RestController
}

type TestSetController struct {
	DB *gorm.DB
}

func NewTestSetController() ITestSetController {
	db := common.InitDB()
	err := db.AutoMigrate(model.TestSet{})
	if err != nil {
		println("数据库初始化失败..")
		return nil
	}
	return TestSetController{DB: db}
}

func (t TestSetController) Create(ctx *gin.Context) {
	str, _ := ctx.Get("user")
	user := str.(model.User)

	var requestTestSet model.TestSet
	err := ctx.Bind(&requestTestSet)

	requestTestSet.Uploader = int(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		response.Fail(ctx, nil, "读取错误，请检查数据格式是否正确")
		return
	}
	if requestTestSet.Input == "" {
		response.Fail(ctx, nil, "数据验证错误，题目输入不能为空")
		return
	}

	t.DB.Create(&requestTestSet)

	response.Success(ctx, gin.H{"testset": requestTestSet}, "创建成功")
}

func (t TestSetController) Update(ctx *gin.Context) {
	var requestTestSet model.TestSet
	err := ctx.Bind(&requestTestSet)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}
	if requestTestSet.Input == "" {
		response.Fail(ctx, nil, "样例输入不能为空")
		return
	}

	TestSetId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if TestSetId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var updateTestSet model.TestSet

	t.DB.First(&updateTestSet, TestSetId)

	if updateTestSet.ID == 0 {
		response.Fail(ctx, nil, "找不到该题目")
		return
	}

	t.DB.Model(&updateTestSet).Update("Input", requestTestSet.Input)

	response.Success(ctx, nil, "修改成功")
}

func (t TestSetController) Show(ctx *gin.Context) {
	TestSetId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if TestSetId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	var TestSet model.TestSet

	t.DB.Where("ID = ?", TestSetId).First(&TestSet)

	if TestSet.ID == 0 {
		response.Fail(ctx, nil, "题目ID错误")
		return
	}

	response.Success(ctx, gin.H{"testset": TestSet}, "获取成功")
}

func (t TestSetController) Delete(ctx *gin.Context) {
	TestSetId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if TestSetId <= 0 {
		response.Fail(ctx, nil, "请求ID错误")
		return
	}

	if err := t.DB.Delete(&model.TestSet{}, TestSetId); err.Error != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
