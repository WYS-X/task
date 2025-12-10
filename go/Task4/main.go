package main

import (
	"fmt"
	"task/Task4/db"
	"task/Task4/router"
)

func main() {
	db := db.InitDB()
	fmt.Println("start !!!")
	router.Init(db)
	fmt.Println("end !!!")
}
