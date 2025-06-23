package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"corsaf1/internal/service"
)

func GetConstructors(c *gin.Context) {
	data, err := service.FetchConstructors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch constructors"})
		return
	}
	c.JSON(http.StatusOK, data)
}
