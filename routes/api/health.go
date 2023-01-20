package routes

import (
	"github.com/gin-gonic/gin"
	
	"github.com/parinyapt/PT-Friendship_Backend/handler/controller"
)

func SetupHealthEndpoint(router *gin.RouterGroup) {
	health := router.Group("/health")
	{
		health.GET("/status", handler.CheckApiStatus)
	}
}
