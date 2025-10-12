// internal/router/router.go
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saku-730/web-specimen/backend/internal/handler"
	"github.com/saku-730/web-specimen/backend/internal/middleware"
)

func SetupRouter(
	occHandler handler.OccrrenceHandler,
)*gin.Engine {
	router := gin.Default()

	//CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	router.User(cores.New(config))

	//API Version
	apiV0_0_2 := router.Group("/api/v0_0_2")//router.Group() make gin.RouterGroup
	{
		apiV0_0_2.POST("/login", authHandler.Login)

		secure := apiV0_0_2.Group("")
		secure.Use(middlrware.AuthMiddleware())
		{	// /create page
			secure.GET("/create", occHandler.GetCreatePage)


		}

	}

	return router
}
