package main

import "time"

func main() {
	// Select é como switch, mas para canais
	// O select é usado para esperar múltiplas operações em canais.
	// Ele seleciona o primeiro canal que estiver pronto para receber ou enviar.

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			canal2 <- "=============Canal 2"
		}
	}()

	for {
		select {
		case mensagem1 := <-canal1:
			println(mensagem1)
		case mensagem2 := <-canal2:
			println(mensagem2)
		}
	}

}
