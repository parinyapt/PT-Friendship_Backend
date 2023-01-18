package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupHealthAPI(router *gin.RouterGroup) {
	health := router.Group("/health")
	{
		health.GET("/status", checkApiStatus)
	}
}
