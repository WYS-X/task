package main

import (
	"context"
	"fmt"
	"math/rand"
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

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})

	ctx := context.Background()

	// user1 := models.User{
	// 	Name: "wang",
	// 	Posts: []models.Post{
	// 		{Title: "不一样的卡梅拉1", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 		{Title: "不一样的卡梅拉2", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 		{Title: "不一样的卡梅拉3", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 		{Title: "不一样的卡梅拉4", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 		{Title: "不一样的卡梅拉5", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 	},
	// }
	// user2 := models.User{
	// 	Name: "lao wang",
	// 	Posts: []models.Post{
	// 		{Title: "不一样的卡梅拉6", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 		{Title: "不一样的卡梅拉7", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 		{Title: "不一样的卡梅拉8", Content: "是由法国作家克利斯提昂·约里波瓦创作的儿童绘本系列，讲述了母鸡卡梅拉及其家人的奇幻冒险故事"},
	// 	},
	// }
	// users := []models.User{user1, user2}
	// err := gorm.G[models.User](db).CreateInBatches(ctx, &users, 10)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	users, err := gorm.G[models.User](db).Preload("Posts", nil).Find(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, u := range users {
		for _, p := range u.Posts {
			c := rand.Intn(10)
			for i := 0; i < c; i++ {
				gorm.G[models.Comment](db).Create(ctx, &models.Comment{PostId: p.ID, UserId: users[0].ID, Content: fmt.Sprintf("用户1的评价%d", i)})
			}
			c = rand.Intn(10)
			for i := 0; i < c; i++ {
				gorm.G[models.Comment](db).Create(ctx, &models.Comment{PostId: p.ID, UserId: users[1].ID, Content: fmt.Sprintf("用户2的评价%d", i)})
			}
		}
	}
}
