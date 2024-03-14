package crud

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) (status_code int, msg string, erro error) {
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	fmt.Println("POST usuário: ", string(corpoDaRequisicao))

	if erro != nil {
		return 400, "Erro ao ler o corpo da requisição.", fmt.Errorf(" %w", erro)
	}

	var new_usuario User

	if erro = json.Unmarshal(corpoDaRequisicao, &new_usuario); erro != nil {
		return 400, "Erro ao decodificar JSON.", fmt.Errorf(" %w", erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		return 500, "Erro ao conectar ao banco de dados.", fmt.Errorf(" %w", erro)
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values ($1, $2) RETURNING id")
	if erro != nil {
		return 500, "Erro ao preparar a instrução SQL.", fmt.Errorf(" %w", erro)
	}
	defer statement.Close()

	var idInserido int32
	erro = statement.QueryRow(new_usuario.Nome, new_usuario.Email).Scan(&idInserido)
	if erro != nil {
		return 500, "Erro ao inserir usuário.", fmt.Errorf(" %w", erro)
	}

	return 201, fmt.Sprintf("Usuário criado com sucesso! ID: %d", idInserido), nil
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) (status_code int, msg string, list_users []User, erro error) {
	fmt.Println("GET usuários")

	db, erro := banco.Conectar()
	if erro != nil {
		return 500, "Erro ao conectar ao banco de dados.", nil, fmt.Errorf(" %w", erro)
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		return 500, "Erro ao buscar os usuários.", nil, fmt.Errorf(" %w", erro)
	}
	defer linhas.Close()

	var usuarios []User
	for linhas.Next() {
		var one_user User
		if erro = linhas.Scan(&one_user.ID, &one_user.Nome, &one_user.Email); erro != nil {
			return 500, "Erro ao buscar os usuários.", nil, fmt.Errorf(" %w", erro)
		}
		usuarios = append(usuarios, one_user)
	}

	return 200, "Usuários listados com sucesso!", usuarios, nil
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) (status_code int, msg string, user User, erro error) {
	var one_user User
	params := mux.Vars(r)
	id, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		return 500, "Erro ao converter o parâmetro para inteiro.", one_user, fmt.Errorf(" %w", erro)
	}

	fmt.Println("GET usuário", params["id"])

	db, erro := banco.Conectar()
	if erro != nil {
		return 500, "Erro ao conectar ao banco de dados.", one_user, fmt.Errorf(" %w", erro)
	}
	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = $1", id)
	if erro != nil {
		return 500, fmt.Sprintf("Erro ao buscar o usuário. ID %d", id), one_user, fmt.Errorf(" %w", erro)
	}
	defer linha.Close()

	if linha.Next() {
		if erro = linha.Scan(&one_user.ID, &one_user.Nome, &one_user.Email); erro != nil {
			return 500, "Erro ao scanear o usuário.", one_user, fmt.Errorf(" %w", erro)
		}
	}

	if one_user.ID == 0 {
		return 404, fmt.Sprintf("ID: %d inválido.", id), one_user, fmt.Errorf("Usuário não encontrado.")
	}

	return 200, "Usuário listado com sucesso!", one_user, nil
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) (status_code int, msg string, user User, erro error) {
	var one_user User
	params := mux.Vars(r)
	id, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		return 500, "Erro ao converter o parâmetro para inteiro.", one_user, fmt.Errorf(" %w", erro)
	}

	corpoDaRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		return 400, "Erro ao ler o corpo da requisição.", one_user, fmt.Errorf(" %w", erro)
	}

	if erro = json.Unmarshal(corpoDaRequisicao, &one_user); erro != nil {
		return 400, "Erro ao decodificar JSON.", one_user, fmt.Errorf(" %w", erro)
	}

	fmt.Println("PUT usuário ", params["id"], string(corpoDaRequisicao))

	db, erro := banco.Conectar()
	if erro != nil {
		return 500, "Erro ao conectar ao banco de dados.", one_user, fmt.Errorf(" %w", erro)
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios set nome = $1, email = $2 WHERE id = $3")
	if erro != nil {
		return 500, "Erro ao preparar a instrução SQL.", one_user, fmt.Errorf(" %w", erro)
	}
	defer statement.Close()

	if _, erro = statement.Exec(one_user.Nome, one_user.Email, id); erro != nil {
		return 500, "Erro ao atualizar o usuário.", one_user, fmt.Errorf(" %w", erro)
	}

	return 200, "Usuário atualizado com sucesso!", one_user, nil
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) (status_code int, msg string, erro error) {
	params := mux.Vars(r)
	id, erro := strconv.ParseUint(params["id"], 10, 32)

	if erro != nil {
		return 500, "Erro ao converter o parâmetro para inteiro.", fmt.Errorf(" %w", erro)
	}

	fmt.Println("DELETE usuário ", params["id"])

	db, erro := banco.Conectar()
	if erro != nil {
		return 500, "Erro ao conectar ao banco de dados.", fmt.Errorf(" %w", erro)
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = $1")
	if erro != nil {
		return 500, "Erro ao preparar a instrução SQL.", fmt.Errorf(" %w", erro)
	}
	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return 500, "Erro ao deletar o usuário.", fmt.Errorf(" %w", erro)
	}

	return 204, "Usuário atualizado com sucesso!", nil
}
