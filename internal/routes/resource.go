package routes

import (
	"encoding/json"
	"net/http"

	"github.com/0xhunterkiller/berry/pkgs/models"
	"github.com/gin-gonic/gin"
)

func addResourceRoutes(rg *gin.RouterGroup) {
	resources := rg.Group("/resources")

	resources.POST("/create", func(c *gin.Context) {
		var dat models.Resource
		if err := json.NewDecoder(c.Request.Body).Decode(&dat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		if err := dat.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "resource created",
			"data":    dat,
		})
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
