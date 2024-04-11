package main

import (
	"fmt"
	// Std library to handel http request on Go
	"net/http"
)

func main() {

	// A base url to define the path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("No of Bytes written:%d", n))

	})

	// http request listener
	http.ListenAndServe(":8080", nil)
}
