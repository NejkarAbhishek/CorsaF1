package main

import (
	"CorsaF1/internal/api"
	"CorsaF1/internal/cache"
	"CorsaF1/internal/config"
	"CorsaF1/internal/repository"
	"CorsaF1/internal/scheduler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	config.LoadEnv()

	if os.Getenv("SKIP_DB") != "true" {
		repository.InitDB()
	} else {
		log.Println("SKIP_DB is true — skipping database initialization.")
	}

	cache.InitRedis()

	scheduler.StartDataSync()

	r := gin.Default()
	api.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
