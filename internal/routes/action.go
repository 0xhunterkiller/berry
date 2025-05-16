package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addActionRoutes(rg *gin.RouterGroup) {
	actions := rg.Group("/actions")

	actions.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, "action created")
	})
	actions.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, "action read")
	})
	actions.PUT("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "action updated")
	})
	actions.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "action updated")
	})
}
