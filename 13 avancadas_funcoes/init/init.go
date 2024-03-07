package main

var n int

func init() {
	println("Inicializando...")
	n = 10
}

func main() {
	// A função init é chamada antes da função main
	// Ela é chamada apenas uma vez
	// Cada arquivo pode ter uma função init

	println("Main...")
	println(n)
}
