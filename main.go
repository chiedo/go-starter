package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello World!")))
}

func Bye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Bye World!")))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/bye", Bye)

	http.ListenAndServe(":8080", r)
}
