package api

import (
	"CorsaF1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CompareDrivers(c *gin.Context) {
	season := c.DefaultQuery("season", "2024")
	driverA := c.Query("a")
	driverB := c.Query("b")

	if driverA == "" || driverB == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing query params 'a' and 'b'"})
		return
	}

	result, err := service.CompareDrivers(season, driverA, driverB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
