package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addResourceTypeRoutes(rg *gin.RouterGroup) {
	restypes := rg.Group("/restypes")

	restypes.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, "restype created")
	})
	restypes.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, "restype read")
	})
	restypes.PUT("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "restype updated")
	})
	restypes.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "restype updated")
	})
}
