package main

import (
	"fmt"
	// Std library to handel http request on Go
	"net/http"
)

const PORT = ":8080"

// Home Route handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page.")
}

// About Route handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page.")
}

func main() {

	// A base url to define the path
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// http request listener
	http.ListenAndServe(PORT, nil)
}
