package main

func calculos(a, b int) (soma int, subtracao int) {
	soma = a + b
	subtracao = a - b
	return
}

func main() {
	soma, subtracao := calculos(10, 20)
	println(soma, subtracao)
}
