package api

import (
	"github.com/gin-gonic/gin"
	"your-module-path/jobs"
	"net/http"
)

func StandingsHandler(c *gin.Context) {
	data, err := jobs.FetchStandings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", data)
}
