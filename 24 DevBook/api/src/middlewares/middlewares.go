package middlewares

import (
	"api/src/autenticacao"
	"api/src/resposta"
	"log"
	"net/http"
)

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s", r.Method, r.Host, r.RequestURI)
		proximaFuncao(w, r)
	}
}

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			resposta.Erro(w, http.StatusUnauthorized, erro)
		} else {
			proximaFuncao(w, r)
		}
	}
}
