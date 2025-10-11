// internal/router/router.go
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	occHandler handler.OccrrenceHandler,
)*gin.Engine {
	router := gin.Default()

	//CORS midleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	router.User(cores.New(config))

	//API Version
	apiV0_0_2 := router.Group("/api/v0_0_2")//router.Group() make gin.RouterGroup
	{
		occHandler.RegisterOccurrenceRoutes(apiV0_0_2)

	}

	return router
}
