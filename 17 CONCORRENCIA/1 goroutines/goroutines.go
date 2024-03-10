package main

import (
	"fmt"
	"time"
)

func main() {
	escrever1("----------------- TESTE ----------------")
	go escrever("Olá Mundo")
	escrever("Programando em Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func escrever1(texto string) {
	fmt.Println(texto)
}
