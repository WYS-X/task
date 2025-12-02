package model

import (
	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string `gorm:"not null"`
	Nickname string `gorm:"not null"`

	Posts   []Post
	Comment []Comment
}
