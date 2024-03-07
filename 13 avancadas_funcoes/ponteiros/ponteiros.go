package main

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	println("Endereço de memória:", numero)
	println("Valor nesse endereço de memória:", *numero)
	*numero = *numero * -1
}

func main() {
	numero := 20
	println(numero)
	println("Invertido:", inverterSinal(numero))
	println(numero)

	println("----------------------")

	novoNumero := 40
	println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	println(novoNumero)

}
