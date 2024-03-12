package main

import (
	"fmt"
	"time"
)

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	// tarfas recebe dados e resultados envia dados
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	start := time.Now()
	println("Iniciando TESTE Worker Pools")

	var calcular_fibonacci int = 45

	tarefas := make(chan int, calcular_fibonacci)
	resultados := make(chan int, calcular_fibonacci)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < calcular_fibonacci; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < calcular_fibonacci; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	elapsed := time.Since(start)
	fmt.Printf("A função main levou %s para executar.\n", elapsed)
}
