package main

import (
	// Std library to handel http request on Go
	"net/http"
)

const PORT = ":8080"

func main() {

	// A base url to define the path
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// http request listener
	http.ListenAndServe(PORT, nil)
}
