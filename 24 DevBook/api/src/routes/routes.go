package routes

import (
	"api/src/routes/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	return rotas.Configurar(mux.NewRouter())
}
