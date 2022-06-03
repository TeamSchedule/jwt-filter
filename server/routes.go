package server

import (
	"github.com/gin-gonic/gin"
	"jwt-filter/server/handlers"
)

func setRoutes(router *gin.Engine) *gin.Engine {
	v1 := router.Group("")
	{
		v1.GET("/", handlers.ValidateAuthorizationToken)
	}

	return router
}
