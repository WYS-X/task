package service

import (
	"fmt"
	"net/http"
	"task/Task4/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{DB: db}
}

type registerModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname" binding:"required"`
}
type loginModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (s *userService) Register(c *gin.Context) {
	fmt.Println("in register")
	var regModel registerModel
	if err := c.ShouldBind(&regModel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regModel.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	user := model.User{Nickname: regModel.Nickname, Email: regModel.Email, Password: string(hashedPassword)}
	if err := s.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": 0, "message": "create success"})
}
func (s *userService) Login(c *gin.Context) {
	var loginModel loginModel
	if err := c.ShouldBind(&loginModel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	var user model.User
	if err := s.DB.Where("email = ?", loginModel.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginModel.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    1,
			"message": "密码错误",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"nickname": user.Nickname,
		"expire":   time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte("xblog"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    1,
			"message": "生成token失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": tokenString,
	})
}
