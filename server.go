package main

import (
	"fmt"
	"kiGo/src/handlers"
	"net/http"
)

func main() {
	// http.HandleFunc("/", handlers.Main_handler)
	// http.HandleFunc("/kigo", handlers.handle_kigo)
	http.HandleFunc("/", handlers.Main_handler)

	go func() {

		err := http.ListenAndServe(":"+fmt.Sprint(API_PORT), nil)
		if err != nil {
			fmt.Println("err")
			fmt.Println(err)
		}
	}()

	select {}
}
