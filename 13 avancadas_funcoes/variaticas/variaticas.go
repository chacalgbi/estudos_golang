package main

// Funções variáticas são funções que aceitam um número variável de argumentos.
func soma(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func escrever(texto string, numeros ...int) {
	// OBS: O parâmetro variático deve ser o último parâmetro da função. E cada função só pode ter um parâmetro variático.
	for _, value := range numeros {
		println(texto, value)
	}
}

func main() {
	println("Total: ", soma(1, 2, 3, 4, 5, 7))
	escrever("Olá", 1, 2, 3, 4, 5)
}
