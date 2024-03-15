package repositorios

import (
	"api/src/modelos"
	"api/src/queries"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db_user struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Db_user {
	return &Db_user{db}
}

func (banco Db_user) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := banco.db.Prepare(queries.Q.CriarUsuario)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	var idInserido uint64
	erro = statement.QueryRow(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha).Scan(&idInserido)
	if erro != nil {
		return 0, erro
	}

	return idInserido, nil
}

func (banco Db_user) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := banco.db.Query(queries.Q.BuscarUsuarios, nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var insert modelos.Usuario
		if erro = linhas.Scan(&insert.Nome, &insert.Nick, &insert.Email, &insert.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, insert)
	}

	return usuarios, nil
}

func (banco Db_user) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linha, erro := banco.db.Query(queries.Q.BuscarUsuarioPorID, id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (banco Db_user) Atualizar(id uint64, usuario modelos.Usuario) error {
	statement, erro := banco.db.Prepare(queries.Q.AtualizarUsuario)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (banco Db_user) Deletar(id uint64) error {
	statement, erro := banco.db.Prepare(queries.Q.DeletarUsuario)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}
