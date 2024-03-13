package banco

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if erro := godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func Conectar() (*sql.DB, error) {
	connStr := "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=" + os.Getenv("DB_SSLMODE")
	db, erro := sql.Open("postgres", connStr)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
