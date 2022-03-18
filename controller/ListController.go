package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/dto"
	"prmlk.com/nextdebug/model"
	"prmlk.com/nextdebug/response"
)

type IListController interface {
	SortProblem(ctx *gin.Context)
	SortOnlineJudge(ctx *gin.Context)
	SortTestSet(ctx *gin.Context)
}

type ListController struct {
	DB *gorm.DB
}

func NewListController() IListController {
	db := common.GetDB()

	return ListController{DB: db}
}

func (c ListController) SortProblem(ctx *gin.Context) {
	var query model.Query
	err := ctx.Bind(&query)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}
	var problem []*model.Problem
	//var problemSet model.ProblemSet

	c.DB.Where("? <= id and id <= ?", query.From, query.To).Find(&problem)
	//for _, set := range problem {
	//	problemSet.Problems = append(problemSet.Problems, set)
	//}
	//另一种实现方法，暂时先不删

	if len(problem) > 0 {
		response.Response(ctx, http.StatusOK, 200, gin.H{"problems": dto.ToProblemsDto(problem)}, "查询完成")
	} else {
		response.Fail(ctx, nil, "查询不到相关的题目")
	}
}

func (c ListController) SortOnlineJudge(ctx *gin.Context) {
	var query model.Query
	err := ctx.Bind(&query)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
	}
	var onlineJudge []*model.OnlineJudge

	c.DB.Where("? <= id and id <= ?", query.From, query.To).Find(&onlineJudge)

	if len(onlineJudge) > 0 {
		response.Response(ctx, http.StatusOK, 200, gin.H{"online_judge": onlineJudge}, "查询完成")
	}
}

func (c ListController) SortTestSet(ctx *gin.Context) {
	var requestProblem model.Problem
	err := ctx.Bind(&requestProblem)
	if err != nil {
		response.Fail(ctx, nil, "读取错误")
		return
	}
	if requestProblem.ID <= 0 {
		response.Fail(ctx, nil, "ID有误")
		return
	}
	var testSet []*model.TestSet

	c.DB.Where("problem_id = ?", requestProblem.ID).Find(&testSet)

	if len(testSet) > 0 {
		response.Response(ctx, http.StatusOK, 200, gin.H{"test_set": dto.ToTestSetsDto(testSet)}, "查询完成")
	}
}
