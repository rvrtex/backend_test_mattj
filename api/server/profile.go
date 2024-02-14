package server

import (
	"github.com/rvrtex/backend_test_mattj/pkg/backendTestMattj"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	profile := route.Group("/profile")
	{
		profile.GET("/:id", getProfile)
	}

}

func getProfile(ctx *gin.Context) {

	singleProfile := &backendTestMattj.Profile{}
	singleProfile, err := singleProfile.GetProfile(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"payload": singleProfile,
	})
}
