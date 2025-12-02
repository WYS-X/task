package main

import (
	"context"
	"fmt"
	"task/Task3/GORM/common"
	"task/Task3/GORM/models"

	"gorm.io/gorm"
)

type PostWithCount struct {
	models.Post
	CommentCount int
}

func main() {
	db := common.GetDB()
	if db == nil {
		fmt.Println("链接数据库失败")
		return
	}
	ctx := context.Background()

	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	user, err := gorm.G[models.User](db).Preload("Posts", nil).Preload("Posts.Comments", nil).First(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("用户", user.Name, "有", len(user.Posts), "篇博客")
	for _, p := range user.Posts {
		fmt.Println(p.Title, "有", len(p.Comments), "个评论")
	}
	var pc PostWithCount
	//获取评论最多的文章信息
	db.Debug().Table("posts as p").
		Select("p.*, count(c.id) as comment_count").
		Joins("left join comments c on p.id = c.post_id").
		Group("p.id").
		Order("comment_count desc").
		First(&pc)

	fmt.Printf("%s 有最多的评论，共%d条评论", pc.Title, pc.CommentCount)
}
