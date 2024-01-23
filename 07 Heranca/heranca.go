package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("--------------------- Herança 'Só que não' ---------------------")

	p1 := pessoa{"João", "Pedro", 20, 1.75}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	e2 := estudante{pessoa{"Maria", "Madalena", 30, 1.62}, "Medicina", "Unicamp"}
	fmt.Println(e2)
	fmt.Println(e2.sobrenome)
	teste := strings.Split(e2.sobrenome, "l")
	fmt.Println(teste)
	fmt.Println(teste[0])

}
