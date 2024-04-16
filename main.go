package main

import (
	// Std library to handel http request on Go
	"net/http"

	"kiGo/src/handlers"
)

const PORT = ":8080"

func main() {

	// A base url to define the path
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/kigo", handlers.Ki_go)

	// http request listener
	http.ListenAndServe(PORT, nil)
}
