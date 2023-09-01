package controller

import (
	"log"
	"net/http"

	"brevo/service"

	"github.com/gin-gonic/gin"
)

func HealthCheckup(c *gin.Context) {
	c.JSON(200, "Okay!")
}

func GetHostnamesHandler(c *gin.Context) {
	result, err := service.GetActiveMTAsAboveThreshold()
	if err != nil {
		log.Println("Unable to get Active MTAs Above Threshold")
	}
	c.JSON(http.StatusOK, result)
}
