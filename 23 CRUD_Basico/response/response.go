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
		Error    bool        `json:"error"`
		Usuários []crud.User `json:"usuarios"`
	}
	var resposta response

	status_code, msg, users, err := crud.BuscarUsuarios(w, r)

	if err != nil {
		w.WriteHeader(status_code)
		resposta = response{msg + err.Error(), true, users}
	} else {
		w.WriteHeader(status_code)
		resposta = response{msg, false, users}
	}

	json.NewEncoder(w).Encode(resposta)
}

func HandleBuscarUsuario(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Mensagem string    `json:"mensagem"`
		Error    bool      `json:"error"`
		Usuário  crud.User `json:"usuario"`
	}
	var resposta response

	status_code, msg, user, err := crud.BuscarUsuario(w, r)

	if err != nil {
		w.WriteHeader(status_code)
		resposta = response{msg + " " + err.Error(), true, user}
	} else {
		w.WriteHeader(status_code)
		resposta = response{msg, false, user}
	}

	json.NewEncoder(w).Encode(resposta)
}

func HandleAtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Mensagem string    `json:"mensagem"`
		Error    bool      `json:"error"`
		Usuário  crud.User `json:"usuario"`
	}
	var resposta response

	status_code, msg, user, err := crud.AtualizarUsuario(w, r)

	if err != nil {
		w.WriteHeader(status_code)
		msg_error := msg + " " + err.Error()
		resposta = response{Mensagem: msg_error, Error: true}
	} else {
		w.WriteHeader(status_code)
		resposta = response{msg, false, user}
	}

	json.NewEncoder(w).Encode(resposta)
}

func HandleDeletarUsuario(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Mensagem string `json:"mensagem"`
		Error    bool   `json:"error"`
	}
	var resposta response

	status_code, msg, err := crud.DeletarUsuario(w, r)
	w.WriteHeader(status_code)

	if err != nil {
		resposta = response{msg + " " + err.Error(), true}
	} else {
		resposta = response{msg, false}
	}

	json.NewEncoder(w).Encode(resposta)
}
