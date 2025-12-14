package main

import (
	"fmt"
	"task/Task4/db"
	logger "task/Task4/log"
	"task/Task4/router"
)

func main() {
	logger.Init()
	defer logger.Log.Sync()

	db := db.InitDB()
	fmt.Println("start !!!")
	router.Init(db)
	fmt.Println("end !!!")
}
