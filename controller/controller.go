package controller

import (
	"mta-hosting-optimizer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HealthCheckup(c *gin.Context) {
	c.JSON(200, "Okay!")
}

func GetHostnamesHandler(c *gin.Context) {

	result, err := service.GetActiveMTAsAboveThreshold()
	if err != nil {
		c.JSON(http.StatusOK, Response{Message: err.Error(), Data: result})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "Success", Data: result})
}
