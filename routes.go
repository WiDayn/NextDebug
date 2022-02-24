package main

import (
	"github.com/gin-gonic/gin"
	"prmlk.com/nextdebug/controller"
	"prmlk.com/nextdebug/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)

	problemRoutes := r.Group("/api/problems")
	problemController := controller.NewProblemController()
	problemRoutes.POST("", problemController.Create)
	problemRoutes.PUT("/:id", problemController.Update)
	problemRoutes.GET("/:id", problemController.Show)
	problemRoutes.DELETE("/:id", problemController.Delete)

	listRoutes := r.Group("/api/list")
	listRoutesController := controller.NewListController()
	listRoutes.POST("/sort", listRoutesController.Sort)

	return r
}
