package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/rvrtex/backend_test_mattj/api/server"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	server.Routes(router)
	router.Run(":8080")
}
