package resposta

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, dados interface{}) {
	w.WriteHeader(status)

	if dados != nil {
		w.Header().Set("Content-Type", "application/json")
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(w http.ResponseWriter, status int, paramErro error) {
	type respErro struct {
		Erro string `json:"erro"`
	}
	err := respErro{Erro: paramErro.Error()}

	JSON(w, status, err)
}
