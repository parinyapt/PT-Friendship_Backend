package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

func CheckApiStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
		"message": "API is running",
		"timestamp": time.Now().Unix(),
	})
}