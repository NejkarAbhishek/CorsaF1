package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)

func GetSchedule(c *gin.Context) {
	resp, err := http.Get("https://api.jolpi.ca/ergast/f1/current.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedule"})
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	c.Data(http.StatusOK, "application/json", body)
}
