package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPermissionRoutes(rg *gin.RouterGroup) {
	perms := rg.Group("/perms")

	perms.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, "perm created")
	})
	perms.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, "perm read")
	})
	perms.PUT("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "perm updated")
	})
	perms.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "perm updated")
	})
}
