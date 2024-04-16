package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

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
	parsedTemp, _ := template.ParseFiles("./templates/"+temp, "./templates/layout/base.layout.tmpl")
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parshing the template", err)
		return
	}

}
