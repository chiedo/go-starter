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

func StaticPagesHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("hello")))
}

func React(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Hello World", StaticDir: os.Getenv("STATIC_DIR")}

	t, _ := template.ParseFiles("resources/views/index.html")

	t.Execute(w, p)
}
