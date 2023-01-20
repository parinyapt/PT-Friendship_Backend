package routes

import (
	"github.com/gin-gonic/gin"

	ApiRoutes "github.com/parinyapt/PT-Friendship_Backend/routes/api"
)

func configApiRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//Public API
			ApiRoutes.SetupHealthEndpoint(v1)

			//Private API with JWT Auth
			
		}
	}
}