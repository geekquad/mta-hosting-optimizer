package main

import (
	"os"

	"brevo/base"
	"brevo/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	base.InitKeyDB()
	// base.InitRedis()
}

func startServer() {
	router := gin.Default()
	routes.Handler(router)
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("localhost:" + port)
}

func main() {
	startServer()
}
