package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
