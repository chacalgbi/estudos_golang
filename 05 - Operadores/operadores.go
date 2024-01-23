package main

import "fmt"

func main() {
	fmt.Println("--------------------- Aritmético ---------------------")
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	fmt.Println("--------------------- Atribuição ---------------------")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	fmt.Println("--------------------- Relacionais ---------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	fmt.Println("--------------------- Lógicos ---------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("--------------------- Unários ---------------------")
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)
	numero *= 3 // numero = numero * 3
	fmt.Println(numero)
	numero /= 2 // numero = numero / 2
	fmt.Println(numero)
	numero %= 3 // numero = numero % 3
	fmt.Println(numero)

	// Ternário não existe em Go
}
