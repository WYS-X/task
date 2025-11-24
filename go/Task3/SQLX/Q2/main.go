package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

// 类型安全映射
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect database fail")
		return
	}

	var books []Book
	err = db.Select(&books, "select * from books where price>=100 order by id desc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(books)
}
