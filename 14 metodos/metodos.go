package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Println("Salvando os dados do usuário: ", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Fulano", 17}
	usuario2 := usuario{"Ciclano", 18}

	usuario1.salvar()
	usuario2.salvar()

	fmt.Println(usuario1.maiorDeIdade())
	fmt.Println(usuario2.maiorDeIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.nome, "fez", usuario2.idade, "anos")
}
