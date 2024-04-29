package routers

import (
	"github.com/ankan-withadream/kiGo/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	mainRouter := gin.New()

	mainRouter.Use(gin.Logger())
	mainRouter.Use(gin.Recovery())

	apiV1Router := mainRouter.Group("/api")

	{

		apiV1Router.GET("/", handlers.Handle_empty)
		apiV1Router.POST("/", handlers.Handle_empty)

		apiV1Router.GET("/kigo", handlers.Handle_kigo)
		apiV1Router.POST("/kigo", handlers.Handle_kigo)

	}

	return mainRouter
}
