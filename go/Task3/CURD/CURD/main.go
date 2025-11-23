package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string
	Age   uint8
	Grade string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("数据库链接失败")
	}

	db.AutoMigrate(&Student{})
	ctx := context.Background()
	//题目1：基本CRUD操作
	//添加张三
	err = gorm.G[any](db).Exec(ctx, "insert into students(name, age, grade) values(?,?,?)", "张三", 20, "三年级")
	if err == nil {
		fmt.Println("添加成功")
	} else {
		fmt.Println(err)
	}
	//查询大于18岁的学生
	var students []Student
	students, _ = gorm.G[Student](db).Raw("select * from students where age > ?", 14).Find(ctx)
	fmt.Println("找到", len(students), "个大于18岁的学生")

	//叫张三的更新为四年级
	err = gorm.G[any](db).Exec(ctx, "update students set grade=? where name = ?", "四年级", "张三")
	if err == nil {
		fmt.Println("更新成功")
	}

	//删除小与15的学生
	err = gorm.G[any](db).Exec(ctx, "delete from students where age < ?", 15)
	if err == nil {
		fmt.Println("删除成功")
	}
}
