package routers

import (
	"kiGo/src/api/handlers"
	"net/http"
)

func Main_route_register(mux *http.ServeMux) {
	register_API_routes(mux)
}

// Register API routes here
func register_API_routes(mux *http.ServeMux) {

	mux.HandleFunc("/api/", handlers.Handle_empty)
	mux.HandleFunc("/api/kigo", handlers.Handle_kigo)

}
