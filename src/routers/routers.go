package routers

import (
	"kiGo/src/api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main_route_register(mux *http.ServeMux) {
	register_API_routes(mux)
}

// Register API routes here
func register_API_routes(mux *http.ServeMux) {

	mux.HandleFunc("/api/", handlers.Handle_empty)
	mux.HandleFunc("/api/kigo", handlers.Handle_kigo)

}

func InitRouter() {
	mainRouter := gin.New()

	mainRouter.use(gin.Logger)
	mainRouter.use(gin.Recovery)

	apiV1Router := mainRouter.Group("/api")
	apiV1Router.GET("/", handlers.Handle_empty)
	// apiV1Router
}
