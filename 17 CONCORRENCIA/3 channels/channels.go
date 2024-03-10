package main

import (
	"fmt"
	"time"
)

func escrever(texto string, milli uint16, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Millisecond * time.Duration(milli))
	}
	close(canal)
}

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo", 100, canal)

	// for {
	// 	msg, isOpen := <-canal
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// ou

	for msg := range canal {
		fmt.Println(msg)
	}

	fmt.Println("----------------- FIM ----------------")
}
