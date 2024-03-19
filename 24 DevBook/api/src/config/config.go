package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar o arquivo .env" + erro.Error())
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	// StringConexaoBanco = "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=" + os.Getenv("DB_SSLMODE")
	StringConexaoBanco = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
}
