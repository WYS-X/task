package main

import (
	"fmt"
	"task/Task3/GORM/common"
	"task/Task3/GORM/models"
)

func main() {
	db := common.GetDB()
	if db == nil {
		fmt.Println("链接数据库失败")
		return
	}

	user1 := models.User{
		Name: "wang",
		Posts: []models.Post{
			{Title: "富强民主文明和谐1", Content: "社会主义核心价值观：", Comments: []models.Comment{
				{Content: "非常好1"},
			}},
			{Title: "富强民主文明和谐2", Content: "社会主义核心价值观：", Comments: []models.Comment{
				{Content: "非常好1"}, {Content: "非常好3"}, {Content: "非常好2"},
			}},
			{Title: "富强民主文明和谐3", Content: "社会主义核心价值观：", Comments: []models.Comment{
				{Content: "非常好1"}, {Content: "非常好3"},
			}},
			{Title: "富强民主文明和谐4", Content: "社会主义核心价值观：", Comments: []models.Comment{}},
		},
	}
	user2 := models.User{
		Name: "wang",
		Posts: []models.Post{
			{Title: "富强民主文明和谐5", Content: "社会主义核心价值观：", Comments: []models.Comment{
				{Content: "非常好1"},
			}},
			{Title: "富强民主文明和谐6", Content: "社会主义核心价值观：", Comments: []models.Comment{
				{Content: "非常好1"}, {Content: "非常好3"}, {Content: "非常好2"},
			}},
			{Title: "富强民主文明和谐3", Content: "社会主义核心价值观：", Comments: []models.Comment{
				{Content: "非常好1"}, {Content: "非常好3"},
			}},
			{Title: "富强民主文明和谐4", Content: "社会主义核心价值观：", Comments: []models.Comment{}},
		},
	}
}
