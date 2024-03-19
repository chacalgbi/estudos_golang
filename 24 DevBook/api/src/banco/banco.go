package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// _ "github.com/lib/pq"

func Conectar() (*sql.DB, error) {
	tipoBanco := "mysql"
	// tipoBanco := "postgres"
	db, erro := sql.Open(tipoBanco, config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
