package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CreateRoutes() {
	// Set up the Router
	r := mux.NewRouter()
	r.HandleFunc("/bye", Bye)

	// Static files
	r.HandleFunc("/static/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// Catch-all route for ReactJS to handle the routing
	r.HandleFunc("/{path:.*}", React)

	http.ListenAndServe(":8080", r)
}
