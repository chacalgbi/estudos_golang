package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=devbook password=devbook dbname=devbook sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o banco de dados testada com sucesso!")

	linhas, err := db.Query("SELECT * FROM usuarios")

	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()

	fmt.Println(linhas.Columns())
}
