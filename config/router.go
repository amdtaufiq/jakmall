package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(ginMode string) *gin.Engine {
	if ginMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	// CORS
	router.Use(cors.Default())

	return router
}
