package routes

import (
	"github.com/gin-gonic/gin"

	"mta-hosting-optimizer/controller"
)

func Handler(r *gin.Engine) {
	r.GET("/admin/healthCheckup", controller.HealthCheckup)
	r.GET("/hostnames", controller.GetHostnamesHandler)
}
