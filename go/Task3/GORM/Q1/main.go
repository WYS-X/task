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

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})

}
