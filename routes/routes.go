package routes

import (
	"brevo/controller"

	"github.com/gin-gonic/gin"
)

func Handler(r *gin.Engine) {
	r.GET("/admin/healthCheckup", controller.HealthCheckup)
	r.GET("/hostnames", controller.GetHostnamesHandler)
}
