package response

import (
	"crud/crud"
	"encoding/json"
	"net/http"
)

func HandleCriarUsuario(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Mensagem string `json:"mensagem"`
		Error    bool   `json:"error"`
	}
	var resposta response

	status_code, msg, err := crud.CriarUsuario(w, r)

	if err != nil {
		w.WriteHeader(status_code)
		resposta = response{msg + err.Error(), true}
	} else {
		w.WriteHeader(status_code)
		resposta = response{msg, false}
	}

	json.NewEncoder(w).Encode(resposta)
}

func HandleBuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Mensagem string      `json:"mensagem"`
		Usuários []crud.User `json:"usuarios"`
		Error    bool        `json:"error"`
	}
	var resposta response

	status_code, msg, users, err := crud.BuscarUsuarios(w, r)

	if err != nil {
		w.WriteHeader(status_code)
		resposta = response{msg + err.Error(), users, true}
	} else {
		w.WriteHeader(status_code)
		resposta = response{msg, users, false}
	}

	json.NewEncoder(w).Encode(resposta)
}

func HandleBuscarUsuario(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Mensagem string `json:"mensagem"`
		Error    bool   `json:"error"`
	}
	var resposta response

	status_code, msg, err := crud.CriarUsuario(w, r)

	if err != nil {
		w.WriteHeader(status_code)
		resposta = response{msg + err.Error(), true}
	} else {
		w.WriteHeader(status_code)
		resposta = response{msg, false}
	}

	json.NewEncoder(w).Encode(resposta)
}
