package main

import (
	// Std library to handel http request on Go
	"net/http"

	"simplehttp/src/handlers"
)

const PORT = ":8080"

func main() {

	// A base url to define the path
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// http request listener
	http.ListenAndServe(PORT, nil)
}
