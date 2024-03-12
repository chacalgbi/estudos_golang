package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	// Convertendo struct para JSON
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(cachorroEmJSON)
	fmt.Println(string(cachorroEmJSON))

	// Convertendo de JSON para struct
	cachorroJSON := `{"nome":"Ted","raca":"Poodle","idade":2}`

	cachorroStruct := cachorro{}

	if erro := json.Unmarshal([]byte(cachorroJSON), &cachorroStruct); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroStruct)
	fmt.Println(cachorroStruct.Nome)

	// Convertendo de JSON para map
	cachorroJSON2 := `{"nome":"Ted","raca":"Poodle"}`
	cachorroMap := map[string]string{}

	if erro := json.Unmarshal([]byte(cachorroJSON2), &cachorroMap); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroMap)

}
