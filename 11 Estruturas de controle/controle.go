package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func main() {
	fmt.Println("---------------Estruturas de controle")
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// if init (outroNumero só existe dentro do if)
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é menor que zero")
	}

	// switch
	switch numero {
	case 10:
		fmt.Println("É 10")
	case 20:
		fmt.Println("É 20")
	default:
		fmt.Println("Não é 10 nem 20")
	}

	fmt.Println(diaDaSemana(4))

}
