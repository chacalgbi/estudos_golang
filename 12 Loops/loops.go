package main

import (
	"fmt"
	"time"
)

func for1() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Valor de i: ", i)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("Valor final de i: ", i)
}

func for2() {
	for j := 0; j < 10; j++ {
		fmt.Println("Valor de J: ", j)
		time.Sleep(50 * time.Millisecond)
	}
}

// Interar sobre um array
func for3() {
	nomes := [3]string{"João", "Maria", "José"}

	// Usando chave e valor
	for index, value := range nomes {
		fmt.Println(index, value)
	}

	// Usando apenas o valor
	for _, value := range nomes {
		fmt.Println(value)
	}
}

// Interar sobre um map
func for4() {
	telefones := map[string]string{
		"João":  "1234-5678",
		"Maria": "1234-5679",
		"José":  "1234-5680",
	}

	// Usando chave e valor
	for nome, telefone := range telefones {
		fmt.Println(nome, telefone)
	}

	// Usando apenas o valor
	for _, telefone := range telefones {
		fmt.Println(telefone)
	}
}

// Interar sobre uma string
func for5() {
	for index, letra := range "PALAVRA" {
		fmt.Println(index, string(letra))
	}

	// Interar sobre uma string ASCII
	for index, letra := range "PALAVRA" {
		fmt.Println(index, letra)
	}
}

func main() {
	for1()
	for2()
	for3()
	for4()
	for5()
}
