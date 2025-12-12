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
	postService := service.NewPostService(db)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "xblog")
	})

	group1 := router.Group("/api/v1")
	{
		group1.POST("/register", userService.Register)
		group1.POST("/login", userService.Login)
	}
	group2 := router.Group("/api/v1")
	group2.Use(middleware.Auth())
	{
		//文章
		group2.POST("/post", postService.AddPost)
		group2.PUT("/post/:id", postService.UpdatePost)
		group2.DELETE("/post/:id", postService.AddPost)
		group2.GET("/post", postService.GetPosts)
		group2.GET("/post/:id", postService.GetPosts)
		//评论
	}
	router.Run(":8080")
}
