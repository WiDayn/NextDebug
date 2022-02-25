package main

import (
	"github.com/gin-gonic/gin"
	"prmlk.com/nextdebug/controller"
	"prmlk.com/nextdebug/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	downloadController := controller.NewDownloadController()
	r.GET("/api/avatar", downloadController.Avatar)

	userRoutes := r.Group("/api/auth")
	userController := controller.NewUserController()
	userRoutes.POST("/register", userController.Register)
	userRoutes.POST("/login", userController.Login)
	userRoutes.GET("/info", middleware.AuthMiddleware(), userController.Info)

	problemRoutes := r.Group("/api/problems")
	problemController := controller.NewProblemController()
	problemRoutes.POST("", problemController.Create)
	problemRoutes.PUT("/:id", problemController.Update)
	problemRoutes.GET("/:id", problemController.Show)
	problemRoutes.DELETE("/:id", problemController.Delete)

	listRoutes := r.Group("/api/list")
	listRoutesController := controller.NewListController()
	listRoutes.POST("/sort_problem", listRoutesController.SortProblem)
	listRoutes.POST("/sort_online_judge", listRoutesController.SortOnlineJudge)

	onlineJudgeRoutes := r.Group("/api/online_judge")
	onlineJudgeContorller := controller.NewOnlineJudgeController()
	onlineJudgeRoutes.POST("", onlineJudgeContorller.Create)
	onlineJudgeRoutes.PUT("/:id", onlineJudgeContorller.Update)
	onlineJudgeRoutes.GET("/:id", onlineJudgeContorller.Show)
	onlineJudgeRoutes.DELETE("/:id", onlineJudgeContorller.Delete)

	uploadAvatarRoutes := r.Group("/api/upload")
	uploadJudgeController := controller.NewUploadController()
	uploadAvatarRoutes.POST("/avatar", middleware.AuthMiddleware(), uploadJudgeController.Avatar)
	return r
}
