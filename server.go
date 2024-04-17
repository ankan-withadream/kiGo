package main

import (
	"fmt"
	"kiGo/src/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Handle_empty)
	http.HandleFunc("/kigo", handlers.Handle_kigo)

	err := http.ListenAndServe(":"+fmt.Sprint(API_PORT), nil)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
}
