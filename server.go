package main

import (
	"fmt"
	"kiGo/internal/config"
	"kiGo/src/routers"
	"net/http"
)

func main() {
	// http.HandleFunc("/", handlers.Main_handler)
	// http.HandleFunc("/kigo", handlers.handle_kigo)
	// http.HandleFunc("/", handlers.Main_handler)
	mux := http.NewServeMux()
	routers.Main_route_register(mux)

	go func() {

		err := http.ListenAndServe(":"+fmt.Sprint(config.APP_CONFIG.API_PORT), mux)
		if err != nil {
			fmt.Println("err")
			fmt.Println(err)
		}
	}()

	select {}
}

// for testing only
// use this when calling hhtp.ListenAndServe()
// http.ListenAndServe(":"+fmt.Sprint(API_PORT), &Main_handler{})

// type Main_handler struct{}

// func (mh *Main_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hitting main handler ServeHTTP")
// }
