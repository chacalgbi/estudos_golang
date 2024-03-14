package main

import (
	"api/src/config"
	"api/src/routes"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config.Carregar()
	fmt.Println("DevBook! Porta: " + strconv.Itoa(config.Porta))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), routes.Gerar()))
}
