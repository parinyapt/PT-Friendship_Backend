package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Setup() {
	router := gin.Default()
	//config
	configCors(router)
	configRateLimit(router)
	s := configApi(router, os.Getenv("PORT"))

	//setup all api route
	

	s.ListenAndServe()
}
