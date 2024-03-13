package crud

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
