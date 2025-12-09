package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (s *userService) Register(c *gin.Context) {
	var model registerModel
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "参数错误"})
	}

}
func (s *userService) Login(c *gin.Context) {

}
