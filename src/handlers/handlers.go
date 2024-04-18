package handlers

import (
	"fmt"
	"kiGo/src/utills"
	"net/http"
)

func Main_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(w)
	// fmt.Println(r)
	if utills.Check_prefix(r.URL.Path, "/api") {
		fmt.Println("got /api...")
		Api_handler(w, r)
	} else {
		fmt.Println("got /...")
		Template_handler(w, r)
	}
}
