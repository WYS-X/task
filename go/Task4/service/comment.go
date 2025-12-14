package service

import (
	"net/http"
	"strconv"

	"task/Task4/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentService struct {
	DB *gorm.DB
}

func NewCommentService(db *gorm.DB) *commentService {
	return &commentService{DB: db}
}

type commentCreateModel struct {
	PostId  int    `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (s *postService) AddComment(c *gin.Context) {
	var p commentCreateModel
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	userId := c.GetInt("userId")
	comment := model.Comment{
		UserId:  userId,
		PostId:  p.PostId,
		Content: p.Content,
	}
	if err := s.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    comment.ID,
	})
}

type commentPost struct {
	Page int `form:"page" binding:"required"`
	Size int `form:"size" binding:"required"`
	ID   int `uri:"ID"`
}
type pageCommentResult struct {
	Page int             `json:"page"`
	Size int             `json:"size"`
	Data []model.Comment `json:"data"`
}

func (s *postService) GetPostComments(c *gin.Context) {
	var pagePost commentPost
	postId, perr := strconv.Atoi(c.Param("ID"))
	if perr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	if err := c.ShouldBindQuery(&pagePost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	db := s.DB.Model(&model.Comment{})
	var result = pageCommentResult{
		Page: pagePost.Page,
		Size: pagePost.Size,
	}
	db.Where("post_id = ?", postId).Order("id desc").Offset((pagePost.Page - 1) * pagePost.Size).Limit(pagePost.Size).Find(&result.Data)
	c.JSON(200, result)
}
