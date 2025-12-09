package main

import (
	"task/Task4/db"
	"task/Task4/router"
)

func main() {
	db := db.InitDB()
	router.Init(db)
}
