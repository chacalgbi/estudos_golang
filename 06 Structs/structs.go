package main

import "fmt"

type user struct {
	nome     string
	idade    uint8
	endereco street
}

type street struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	endereco := street{"Rua dos Bobos", 0}

	//Modo 1
	var usuario1 user
	usuario1.nome = "Rafael"
	usuario1.idade = 32
	fmt.Println(usuario1)

	// Modo 2
	usuario2 := user{"Rafael", 32, endereco}
	fmt.Println(usuario2)

	// Modo 3
	usuario3 := user{nome: "Rafael"}
	fmt.Println(usuario3)
}
