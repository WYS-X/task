package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Posts []Post
}
type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserId   uint
	Comments []Comment
}
type Comment struct {
	gorm.Model
	Content string
	PostId  uint
	UserId  uint
	User    User
}
