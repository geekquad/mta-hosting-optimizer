package controller

import (
	"mta-hosting-optimizer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckup(c *gin.Context) {
	c.JSON(200, "Okay!")
}

func GetHostnamesHandler(c *gin.Context) {

	result, _ := service.GetActiveMTAsAboveThreshold()

	c.JSON(http.StatusOK, result)
}
