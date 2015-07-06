package config

import (
	"github.com/gorilla/mux"
	"go-starter/app/controllers"
	"net/http"
)

func Routes() {
	// Set up the Router
	r := mux.NewRouter()

	//
	// The order of the following three route sections is important
	// Golang, Static, then React
	//

	//
	// Golang Routes
	//
	r.HandleFunc("/hello", controllers.StaticPagesHello)

	//
	// Static files
	//
	r.HandleFunc("/static/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	//
	// ReactJS (Catch-all)
	//
	r.HandleFunc("/{path:.*}", controllers.React)

	http.ListenAndServe(":3001", r)
}
