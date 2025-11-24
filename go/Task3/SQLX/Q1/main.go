package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float32
	CreateTime time.Time `db:"createTime"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect database fail")
		return
	}

	var employees []Employee
	err = db.Select(&employees, "select * from employees where department=?", "技术部")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(employees)

	var maxSalaryEmployee Employee
	err = db.Get(&maxSalaryEmployee, "select * from employees order by salary desc limit 1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(maxSalaryEmployee.ID)
}
