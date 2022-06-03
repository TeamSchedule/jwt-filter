package server

import (
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.New()

	// Change default logger to custom
	r = setLogger(r)

	// Set paths to handler functions
	r = setRoutes(r)

	return r
}
