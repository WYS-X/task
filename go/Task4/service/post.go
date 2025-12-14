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

type postCreateModel struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (s *postService) AddPost(c *gin.Context) {
	var p postCreateModel
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	userId := c.GetInt("userId")
	post := model.Post{
		UserID:  userId,
		Title:   p.Title,
		Content: p.Content,
	}
	if err := s.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
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
func (s *postService) GetPost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindUri(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	if err := s.DB.Preload("User").First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "文章不存在",
		})
		return
	}
	c.JSON(200, post)
}
func (s *postService) UpdatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindUri(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	var postContent postCreateModel
	if err := c.ShouldBind(&postContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	userID := c.GetInt("userId")
	post.Title = postContent.Title
	post.Content = postContent.Content
	postResult := s.DB.Model(&post).Where("user_id = ?", userID).Updates(post)
	if postResult.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "文章不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新成功",
	})
}
func (s *postService) DeletePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindUri(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	userID := c.GetInt("userId")
	postResult := s.DB.Model(&post).Where("user_id = ?", userID).Delete(&post)
	if postResult.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "文章不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除成功",
	})
}

type pagePost struct {
	Page   int `form:"page" binding:"required"`
	Size   int `form:"size" binding:"required"`
	UserID int `form:"userId"`
}
type pagePostResult struct {
	Page int          `json:"page"`
	Size int          `json:"size"`
	Data []model.Post `json:"data"`
}

func (s *postService) GetPosts(c *gin.Context) {
	var pagePost pagePost
	if err := c.ShouldBindQuery(&pagePost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "参数错误",
		})
		return
	}
	db := s.DB.Model(&model.Post{})
	if pagePost.UserID > 0 {
		db = db.Where("user_id = ?", pagePost.UserID)
	}
	var result = pagePostResult{
		Page: pagePost.Page,
		Size: pagePost.Size,
	}
	db.Order("id desc").Offset((pagePost.Page - 1) * pagePost.Size).Limit(pagePost.Size).Find(&result.Data)
	c.JSON(200, result)
}
