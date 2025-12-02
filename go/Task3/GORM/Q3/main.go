package main

import (
	"fmt"
	"task/Task3/GORM/common"
	"task/Task3/GORM/models"
)

func main() {
	db := common.GetDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})

	//在文章创建时自动更新用户的文章数量统计字段。
	// var user models.User
	// post := models.Post{
	// 	Title:   "向阳花1",
	// 	Content: "那美丽的天 总是一望无边",
	// 	UserId:  3,
	// }
	// db.Create(&post)
	// db.First(&user, 3)
	// fmt.Println(user.Name, "的文章数量：", user.PostCount)

	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	var post models.Post
	db.Last(&post)
	comment := models.Comment{
		PostId:  post.ID,
		Content: "it's good",
		UserId:  3,
	}
	db.Create(&comment)
	db.First(&post, post.ID)
	fmt.Println(post.Title, "评论数量：", post.CommentCount, "，评论状态：", post.CommentStatus)
	db.Delete(&comment)
	db.First(&post, post.ID)
	fmt.Println(post.Title, "评论数量：", post.CommentCount, "，评论状态：", post.CommentStatus)
}
