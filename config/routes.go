package config

import (
	"github.com/gorilla/mux"
	"go-starter/app/controllers"
	"net/http"
)

func Routes() {
	// Set up the Router
	r := mux.NewRouter()
	r.HandleFunc("/bye", controllers.Bye)

	// Static files
	r.HandleFunc("/static/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// Catch-all route for ReactJS to handle the routing
	r.HandleFunc("/{path:.*}", controllers.React)

	http.ListenAndServe(":3001", r)
}
