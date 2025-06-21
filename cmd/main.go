package main

import (
	"CorsaF1/internal/api"
	"CorsaF1/internal/cache"
	"CorsaF1/internal/config"
	"CorsaF1/internal/repository"
	"CorsaF1/internal/scheduler"
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
