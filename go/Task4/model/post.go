package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Content      string `gorm:"not null"`
	ViewCount    int
	CommentCount int
	UserID       int
	User         User
}
