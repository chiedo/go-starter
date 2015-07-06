package app

import (
	"fmt"
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

	CreateRoutes()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
