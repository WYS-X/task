package router

import (
	logger "task/Task4/log"
	"task/Task4/service"
	"time"

	"task/Task4/router/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	router := gin.New()
	router.Use(GinLogger(), GinRecover())

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
	group2 := router.Group("/api/v1", middleware.Auth())
	{
		//文章
		group2.POST("/post", postService.AddPost)
		group2.PUT("/post/:ID", postService.UpdatePost)
		group2.DELETE("/post/:ID", postService.DeletePost)
		group2.GET("/post", postService.GetPosts)
		group2.GET("/post/:ID", postService.GetPost)
		//评论
		group2.POST("/comment", postService.AddComment)
		group2.GET("/post/:ID/comment", postService.GetPostComments)
	}
	router.Run(":8080")
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)
		if raw != "" {
			path = path + "?" + raw
		}
		logger.Log.Info("http_request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.Duration("cost", cost))
	}
}

func GinRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error("panic",
					zap.Any("error", err),
					zap.Stack("stack"),
				)
				c.AbortWithStatus(500)
			}
		}()
	}
}
