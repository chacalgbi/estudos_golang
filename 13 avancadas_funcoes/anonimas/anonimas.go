package main

func main() {
	// Função anônima
	func() {
		println("Olá, mundo!")
	}()

	// Função anônima com parâmetros
	func(texto string) {
		println(texto)
	}("Passando parâmetro para a função anônima")

	// Função anônima com retorno
	retorno := func(texto string) string {
		return texto
	}("Passando parâmetro para a função anônima com retorno")
	println(retorno)
}
