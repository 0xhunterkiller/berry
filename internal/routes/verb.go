package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addVerbRoutes(rg *gin.RouterGroup) {
	verbs := rg.Group("/verbs")

	verbs.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, "verb created")
	})
	verbs.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, "verb read")
	})
	verbs.PUT("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, "verb updated")
	})
	verbs.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "verb updated")
	})
}
