package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// home Route handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// about Route handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")

}

// kigo Route handler
func Ki_go(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "kiGo.html")
}

// Function to Parse the html files
func renderTemplate(w http.ResponseWriter, temp string) {
	// Go Std. library template used to parse the html file
	parsedTemp, err2 := template.ParseFiles("./src/templates/"+temp, "./src/templates/layout/base.layout.tmpl")
	fmt.Print(parsedTemp, err2)
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parshing the template", err)
		return
	}

}
