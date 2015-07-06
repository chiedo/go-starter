package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title     string
	StaticDir string
	//Body  []byte
}

func Hello(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("MESSAGE")
	w.Write([]byte(fmt.Sprintf(msg)))
}

func Bye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Bye World!")))
}

func React(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Hello World", StaticDir: os.Getenv("STATIC_DIR")}

	t, _ := template.ParseFiles("resources/views/index.html")

	t.Execute(w, p)
}
