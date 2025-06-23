package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/drivers", GetDrivers)
	r.GET("/schedule", GetSchedule)
	r.GET("/constructors", GetConstructors)
	r.GET("/compare", CompareDrivers)
}
