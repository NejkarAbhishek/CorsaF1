package main

import (
	"github.com/gin-gonic/gin"
	"your-module-path/api"
	"your-module-path/db"
	"log"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/standings", api.StandingsHandler)

	r.Run() // localhost:8080
}
