package routes

import (
	"github.com/gin-gonic/gin"
	
	"github.com/parinyapt/PT-Friendship_Backend/handler/controller"
)

func SetupTestEndpoint(router *gin.RouterGroup) {
	test := router.Group("/test")
	{
		test.GET("/demo", handler.DemoHandler)
	}
}
