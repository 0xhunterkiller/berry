package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func addUtilRoutes(rg *gin.RouterGroup) {
	utils := rg.Group("")

	utils.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	utils.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"healthy":   "true",
			"timestamp": time.Now().UnixMicro(),
		})
	})
}
