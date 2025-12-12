package service

import (
	"net/http"

	"task/Task4/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *postService {
	return &postService{DB: db}
}

type postModel struct {
	ID      int    `json:"id"`
	Title   string `json:"title";binding:"required"`
	Content string `json:"content";binding:"required"`
}

func (s *postService) AddPost(c *gin.Context) {
	var p postModel
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	post := model.Post{
		Title:   p.Title,
		Content: p.Content,
	}
	if err := s.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    post.ID,
	})
}
func (s *postService) UpdatePost(c *gin.Context) {

}
func (s *postService) DeletePost(c *gin.Context) {

}
func (s *postService) GetPosts(c *gin.Context) {

}
