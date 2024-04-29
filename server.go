package main

import (
	"fmt"
	"net/http"

	"github.com/ankan-withadream/kiGo/internal/config"
	"github.com/ankan-withadream/kiGo/internal/db"
	"github.com/ankan-withadream/kiGo/src/routers"
)

func main() {
	// http.HandleFunc("/", handlers.Main_handler)
	// http.HandleFunc("/kigo", handlers.handle_kigo)
	// http.HandleFunc("/", handlers.Main_handler)
	// mux := http.NewServeMux()
	// routers.Main_route_register(mux)
	config.Init_env()

	db_pgx := db.Init_DB()
	fmt.Println(db_pgx)

	router := routers.InitRouter()

	server := &http.Server{
		Addr:         ":" + fmt.Sprint(config.APP_CONFIG.API_PORT),
		Handler:      router,
		ReadTimeout:  config.APP_CONFIG.ReadTimeout,
		WriteTimeout: config.APP_CONFIG.WriteTimeout,
	}

	go func() {

		err := server.ListenAndServe()
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
