package router

import (
	"task/Task4/service"

	"task/Task4/router/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	router := gin.Default()
	//  router.Use(middleware.)

	userService := service.NewUserService(db)

	group1 := router.Group("/api/v1")
	{
		group1.POST("/register", userService.Register)
		group1.POST("/login", userService.Login)
	}
	group2 := router.Group("/api/v1")
	group2.Use(middleware.Auth())
	{
		//文章
		//评论
	}
}
