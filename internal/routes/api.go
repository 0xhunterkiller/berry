package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {
	getRoutes()

	router.Run(":9091")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addUtilRoutes(v1)
	addRoleRoutes(v1)
	addResourceRoutes(v1)
}
