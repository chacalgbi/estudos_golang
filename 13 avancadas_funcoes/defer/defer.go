package main

func funcao1() {
	println("funcao1")
}

func funcao2() {
	println("funcao2")
}

func main() {
	// defer adia a execução da função até o final do bloco
	defer funcao1()
	funcao2()
}
