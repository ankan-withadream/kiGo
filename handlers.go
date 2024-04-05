package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
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
