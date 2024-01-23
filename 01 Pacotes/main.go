package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("chacalgbi@hotmail.com")
	fmt.Println(erro)
}
