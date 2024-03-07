package main

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	println("Função recursiva")

	// posicao := uint(12)
	// ou
	var posicao uint = 30
	println(fibonacci(posicao))
}
