package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo"
	canal <- "Programando em Go!"

	msg1 := <-canal
	msg2 := <-canal

	fmt.Println(msg1)
	fmt.Println(msg2)
}
