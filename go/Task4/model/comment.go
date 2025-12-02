package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`

	UserId int
	User   User
	PostId int
	Post   Post
}
