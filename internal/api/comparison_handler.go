package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"corsaf1/internal/service"
)

func CompareDrivers(c *gin.Context) {
	season := c.DefaultQuery("season", "2024")
	driverA := c.Query("a")
	driverB := c.Query("b")

	if driverA == "" || driverB == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing query params 'a' and 'b'"})
		return
	}

	result := service.CompareDrivers(season, driverA, driverB)
	c.JSON(http.StatusOK, result)
}
