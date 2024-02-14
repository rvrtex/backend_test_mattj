package server

import (
	"github.com/rvrtex/backend_test_mattj/pkg/backendTestMattj"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	profile := route.Group("/profile")
	{
		profile.GET("/", getProfile)
	}

}

func getProfile(ctx *gin.Context) {

	profile := backendTestMattj.Profile{}
	profile.GetProfile(ctx)
	ctx.JSON(200, gin.H{
		"message": profile,
	})
}
