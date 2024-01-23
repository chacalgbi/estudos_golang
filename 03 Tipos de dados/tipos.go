package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1 int8 = 100
	var numero2 int16 = 4100
	var numero3 int32 = 554544
	var numero4 int64 = 4105454544654566540
	fmt.Println(numero1, numero2, numero3, numero4)

	var numero5 float32 = 123.45
	var numero6 float64 = 123045.4545
	fmt.Println(numero5, numero6)
	fmt.Println(numero5 + float32(numero6))

	var str string = "Texto"
	fmt.Println(str)

	// quando se declara uma variável numérica com inferencia, ela é do tipo int apenas
	// esse int é definido pelo processador, 32 ou 64 bits, a mesma coisa para float, float32 ou float64

	// error é um tipo de variável

	var erro_teste error = errors.New("Erro interno")
	fmt.Println(erro_teste)
}
