package main

func recuperarPanic() {
	if r := recover(); r != nil {
		println("Recuperado:", r)
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarPanic()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6!")
}

func main() {
	println(alunoAprovado(7, 8))
	println("-------------Pós execução 1")
	println(alunoAprovado(6, 6))
	println("-------------Pós execução 2")
	println(alunoAprovado(8, 9))
	println("-------------Pós execução 3")
}
