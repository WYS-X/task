package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1,
				"message": "未登录",
			})
			return
		}
		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			return []byte("xblog"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1,
				"message": "令牌错误",
			})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1,
				"message": "签名无效",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if exp, ok := claims["expire"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    1,
					"message": "令牌过期",
				})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1,
				"message": "时间戳无效",
			})
			return
		}

		c.Set("userId", int(claims["id"].(float64)))
		c.Set("nickname", claims["nickname"].(string))
		c.Next()
	}
}
