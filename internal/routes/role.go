package routes

import (
	"encoding/json"
	"net/http"

	"github.com/0xhunterkiller/berry/pkgs/models"
	"github.com/gin-gonic/gin"
)

func addRoleRoutes(rg *gin.RouterGroup) {
	roles := rg.Group("/roles")

	roles.POST("/create", func(c *gin.Context) {
		var dat models.Role
		if err := json.NewDecoder(c.Request.Body).Decode(&dat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		if err := dat.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
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
