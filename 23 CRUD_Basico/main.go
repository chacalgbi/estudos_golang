package main

import (
	"crud/response"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", response.HandleCriarUsuario).Methods("POST")
	router.HandleFunc("/usuarios", response.HandleBuscarUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", response.HandleBuscarUsuario).Methods("GET")

	fmt.Println("Servidor iniciado na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
