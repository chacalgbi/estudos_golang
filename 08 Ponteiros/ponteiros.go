package main

import "fmt"

func main() {
	fmt.Println("---------- Sem ponteiros ----------")
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)

	fmt.Println("---------- Ponteiro é uma referência de memória ----------")
	var variavel3 int
	var ponteiro *int
	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)  // Assim mostra o endereço de memória do ponteiro
	fmt.Println(variavel3, *ponteiro) // Assim mostra o valor dentro do endereço de memória do ponteiro
	variavel3++
	fmt.Println(variavel3, ponteiro)  // Assim mostra o endereço de memória do ponteiro
	fmt.Println(variavel3, *ponteiro) // Assim mostra o valor dentro do endereço de memória do ponteiro
}
