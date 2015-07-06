package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("MESSAGE")
	w.Write([]byte(fmt.Sprintf(msg)))
}

func Bye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Bye World!")))
}

func React(w http.ResponseWriter, r *http.Request) {
	type Page struct {
		Title     string
		StaticDir string
		//Body  []byte
	}

	p := &Page{Title: "Hello World", StaticDir: os.Getenv("STATIC_DIR")}

	t, _ := template.ParseFiles("resources/views/index.html")

	t.Execute(w, p)
}

func main() {
	// Set up dotenv
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
