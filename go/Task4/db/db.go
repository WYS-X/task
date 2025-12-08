package db

import (
	"task/Task4/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1)/xblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("初始化数据库错误")
	}

	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	return db
}
