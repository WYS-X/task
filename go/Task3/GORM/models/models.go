package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Posts    []Post
	Comments []Comment
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
	Post    Post
	UserId  uint
	User    User
}
