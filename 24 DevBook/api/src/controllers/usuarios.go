package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/resposta"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user modelos.Usuario
	if erro = json.Unmarshal(body, &user); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Preparar("cadastro"); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	user.ID, erro = repositorio.Criar(user)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	user.Senha = ""
	resposta.JSON(w, http.StatusCreated, user)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	users, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSON(w, http.StatusOK, users)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	user, erro := repositorio.BuscarPorID(id)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if user.ID == 0 {
		erroNotUser := errors.New("usuário não encontrado")
		resposta.Erro(w, http.StatusNotFound, erroNotUser)
		return
	}

	resposta.JSON(w, http.StatusOK, user)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if !compararId(id, w, r) {
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user modelos.Usuario
	if erro = json.Unmarshal(body, &user); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Preparar("edicao"); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(id, user); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if !compararId(id, w, r) {
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(id); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSON(w, http.StatusNoContent, nil)
}

func compararId(id uint64, w http.ResponseWriter, r *http.Request) bool {
	usuarioIDNotoken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return false
	}

	if id != usuarioIDNotoken {
		fmt.Println("usuarioIDNotoken: ", usuarioIDNotoken)
		fmt.Println("ID requisição: ", id)
		resposta.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return false
	}

	return true
}
