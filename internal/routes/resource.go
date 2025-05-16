package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addResourceRoutes(rg *gin.RouterGroup) {
	resources := rg.Group("/resources")

	resources.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, "resource created")
	})
	resources.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, "resource read")
	})
	resources.PUT("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "resource updated")
	})
	resources.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "resource updated")
	})
}
