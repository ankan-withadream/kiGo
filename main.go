package main

import (
	"fmt"
	"text/template"

	// Std library to handel http request on Go
	"net/http"
)

const PORT = ":8080"

// Home Route handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// About Route handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")

}

// Function to Parse the html files
func renderTemplate(w http.ResponseWriter, temp string) {
	// Go Std. library template used to parse the html file
	parsedTemp, _ := template.ParseFiles("./templates/" + temp)
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parshing the template", err)
		return
	}

}

func main() {

	// A base url to define the path
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// http request listener
	http.ListenAndServe(PORT, nil)
}
