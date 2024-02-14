package server

import (
	"github.com/rvrtex/backend_test_mattj/pkg/backendTestMattj"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/profile", getProfile)
	route.GET("/person/:id", getPersonByID)
	route.POST("/person", postPerson)
}

func getProfile(ctx *gin.Context, params backendTestMattj.Profile) {
	ctx.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
