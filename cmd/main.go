package main

import (
	"CORSAF1/internal/api"
	"CORSAF1/internal/cache"
	"CORSAF1/internal/config"
	"CORSAF1/internal/repository"
	"CORSAF1/internal/scheduler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config.LoadEnv()

	repository.InitDB()

	cache.InitRedis()

	scheduler.StartDataSync()

	r := gin.Default()
	api.RegisterRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
