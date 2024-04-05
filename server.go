package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/kigo", handle_kigo)

	err := http.ListenAndServe(":4277", nil)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
}
