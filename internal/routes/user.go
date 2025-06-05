package routes

import (
	"encoding/json"
	"net/http"

	"github.com/0xhunterkiller/berry/pkgs/models"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/create", func(c *gin.Context) {
		var dat models.User
		if err := json.NewDecoder(c.Request.Body).Decode(&dat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		if err := dat.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
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
