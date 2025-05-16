package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, "user created")
	})
	users.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, "user read")
	})
	users.PUT("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users updated")
	})
	users.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users updated")
	})
}
