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
			ApiRoutes.SetupTestEndpoint(v1)
			ApiRoutes.SetupHealthEndpoint(v1)
			// ApiRoutes.SetupProfileEndpoint(v1)

			//Private API with JWT Auth
			
		}
	}
}