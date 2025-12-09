package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("验证用户")
		c.Next()
	}
}
