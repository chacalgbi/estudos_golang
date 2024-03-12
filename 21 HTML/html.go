package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ROTA GET handler")
	w.Write([]byte("Hello, World!"))
}

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ROTA GET inicial")
	u := usuario{"Fulano", "fulano@gmail.com"}
	templates.ExecuteTemplate(w, "inicial.html", u)
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/hello", handler)
	http.HandleFunc("/inicial", inicial)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
