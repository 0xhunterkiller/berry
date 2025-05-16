package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoleRoutes(rg *gin.RouterGroup) {
	roles := rg.Group("/roles")

	roles.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, "role created")
	})
	roles.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, "role read")
	})
	roles.PUT("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "role updated")
	})
	roles.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "role updated")
	})
}
