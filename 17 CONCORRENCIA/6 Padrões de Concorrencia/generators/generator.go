package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Second)
		}
		close(c)
	}()
	return c
}
