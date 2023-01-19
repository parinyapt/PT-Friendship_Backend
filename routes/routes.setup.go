package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	
	"github.com/parinyapt/PT-Friendship_Backend/handler"
)

func Setup() {
	router := gin.Default()
	//config
	configCors(router)
	configRateLimit(router)
	s := configApi(router, os.Getenv("PORT"))

	//setup all api route
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//Public API
			handler.SetupHealthAPI(v1)

			//Private API with JWT Auth
			
		}
	}

	s.ListenAndServe()
}
