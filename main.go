package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"mta-hosting-optimizer/base"
	"mta-hosting-optimizer/routes"
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
