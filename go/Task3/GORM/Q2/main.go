package main

import (
	"context"
	"fmt"
	"task/Task3/GORM/common"
	"task/Task3/GORM/models"

	"gorm.io/gorm"
)

func main() {
	db := common.GetDB()
	if db == nil {
		fmt.Println("链接数据库失败")
		return
	}
	ctx := context.Background()

	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	user, err := gorm.G[models.User](db).Preload("Posts", nil).Preload("Posts.Comments", nil).Where("id = ?", 1).First(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("用户", user.Name, "有", len(user.Posts), "篇博客")
	for _, p := range user.Posts {
		fmt.Println(p.Title, "有", len(p.Comments), "个评论")
	}

	//后去评论最多的文章信息
	gorm.G[models.Post](db).Joins("")
}
