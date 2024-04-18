package handlers

import (
	"fmt"
	"kiGo/src/utills"
	"net/http"
)

func Api_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Got api route: ", r.RequestURI, "\n")
	switch utills.Rem_prefix(r.RequestURI, "/api") {
	case "/":
		handle_empty(w, r)
	case "/kigo":
		handle_kigo(w, r)
	}
}

func handle_empty(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Get Response"))
	case http.MethodPost:
		fmt.Fprint(w, "POST request")
	default:
		fmt.Fprintf(w, "Unsupported method: %s", r.Method)
	}
}

func handle_kigo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, http.MethodGet)
	w.Write([]byte("Ei cholche Go :)"))
}
