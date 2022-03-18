package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"prmlk.com/nextdebug/controller"
	"prmlk.com/nextdebug/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	cfg, _ := ini.Load("config/app.ini")
	r.Use(middleware.CORSMiddleware())

	if cfg.Section("").Key("app_mode").String() != "development" {
		r.Use(middleware.TLSMiddleware())
	}

	downloadController := controller.NewDownloadController()
	r.GET("/api/avatar", downloadController.Avatar)

	userRoutes := r.Group("/api/auth")
	userController := controller.NewUserController()
	userRoutes.POST("/register", userController.Register)
	userRoutes.POST("/login", userController.Login)
	userRoutes.GET("/info", middleware.AuthMiddleware(), userController.Info)
	userRoutes.POST("/userDetail", userController.Show)

	problemRoutes := r.Group("/api/problems")
	problemController := controller.NewProblemController()
	problemRoutes.POST("", middleware.AuthMiddleware(), problemController.Create)
	problemRoutes.PUT("/:id", middleware.AuthMiddleware(), problemController.Update)
	problemRoutes.GET("/:originalID", problemController.Show)
	problemRoutes.DELETE("/:id", middleware.AuthMiddleware(), problemController.Delete)

	problemListRoutes := r.Group("api/problems_list")
	problemListController := controller.NewProblemListController()
	problemListRoutes.POST("", middleware.AuthMiddleware(), problemListController.Create)
	problemListRoutes.PUT("/:id", middleware.AuthMiddleware(), problemListController.Update)
	problemListRoutes.GET("/:id", problemListController.Show)
	problemListRoutes.DELETE("/:id", middleware.AuthMiddleware(), problemListController.Delete)

	problemTagRoutes := r.Group("api/problems_tag")
	problemTagController := controller.NewProblemTagController()
	problemTagRoutes.POST("", problemTagController.Create)
	problemTagRoutes.PUT("/:id", middleware.AuthMiddleware(), problemTagController.Update)
	problemTagRoutes.GET("/:id", problemTagController.Show)
	problemTagRoutes.DELETE("/:id", middleware.AuthMiddleware(), problemTagController.Delete)

	listRoutes := r.Group("/api/list")
	listRoutesController := controller.NewListController()
	listRoutes.POST("/sort_problem", listRoutesController.SortProblem)
	listRoutes.POST("/sort_online_judge", listRoutesController.SortOnlineJudge)
	listRoutes.POST("/sort_test_set", listRoutesController.SortTestSet)

	onlineJudgeRoutes := r.Group("/api/online_judge")
	onlineJudgeController := controller.NewOnlineJudgeController()
	onlineJudgeRoutes.POST("", middleware.AuthMiddleware(), onlineJudgeController.Create)
	onlineJudgeRoutes.PUT("/:id", middleware.AuthMiddleware(), onlineJudgeController.Update)
	onlineJudgeRoutes.GET("/:id", onlineJudgeController.Show)
	onlineJudgeRoutes.DELETE("/:id", middleware.AuthMiddleware(), onlineJudgeController.Delete)

	uploadAvatarRoutes := r.Group("/api/upload")
	uploadJudgeController := controller.NewUploadController()
	uploadAvatarRoutes.POST("/avatar", middleware.AuthMiddleware(), uploadJudgeController.Avatar)

	testSetRoutes := r.Group("/api/test_set")
	testsController := controller.NewTestSetController()
	testSetRoutes.POST("", middleware.AuthMiddleware(), testsController.Create)
	testSetRoutes.PUT("/:id", middleware.AuthMiddleware(), testsController.Update)
	testSetRoutes.GET("/:id", testsController.Show)
	testSetRoutes.DELETE("/:id", middleware.AuthMiddleware(), testsController.Delete)
	return r
}
