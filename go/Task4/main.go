package main

import (
	"task/Task4/db"
	logger "task/Task4/log"
	"task/Task4/router"
)

func main() {
	logger.Init()
	defer logger.Log.Sync()

	db := db.InitDB()
	logger.Log.Info("start")
	router.Init(db)
	logger.Log.Info("end")
}
