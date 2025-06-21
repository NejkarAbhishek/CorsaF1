package api

import (
	"f1insight/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDrivers(c *gin.Context) {
	drivers, err := service.FetchDriverStandings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch driver standings"})
		return
	}
	c.JSON(http.StatusOK, drivers)
}
