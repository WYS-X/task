package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Posts     []Post
	Comments  []Comment
	PostCount int
}
type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserId        uint
	Comments      []Comment
	CommentCount  int
	CommentStatus string `gorm:"default:'无评论'"`
}
type Comment struct {
	gorm.Model
	Content string
	PostId  uint
	Post    Post
	UserId  uint
	User    User
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	tx.First(&user, p.UserId)
	user.PostCount += 1
	tx.Model(&user).Update("post_count", user.PostCount)
	return
}
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var post Post
	tx.First(&post, c.PostId)
	post.CommentCount += 1
	post.CommentStatus = "有评论"
	tx.Model(&post).Select("CommentCount", "CommentStatus").Updates(&post)
	return
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	tx.First(&post, c.PostId)
	post.CommentCount -= 1
	if post.CommentCount == 0 {
		post.CommentStatus = "无评论"
	}
	tx.Model(&post).Select("CommentCount", "CommentStatus").Updates(&post)
	return
}
