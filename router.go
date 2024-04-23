package main

import (
	"kiGo/src/handlers"
	"net/http"
)

func main_route_register(mux *http.ServeMux) {
	Register_API_routes(mux)
	Register_Template_routes(mux)
}

// Register API routes here
func Register_API_routes(mux *http.ServeMux) {

	mux.HandleFunc("/api/", handlers.Handle_empty)
	mux.HandleFunc("/api/kigo", handlers.Handle_kigo)

}

// Register general routes here
func Register_Template_routes(mux *http.ServeMux) {

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/kigo", handlers.Ki_go)

}
