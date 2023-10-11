package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// rota representa todas as rotas da aplicacao web
type Rota struct {
	URI        string
	Metodo     string
	Funcao     func(http.ResponseWriter, *http.Request)
	RequerAuth bool
}

// configurar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotaLogout)
	rotas = append(rotas, RotasUsuario...)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		//verifica se a rota precisa de autenticacao e se sim, passa a funcao de autenticacao dentro da funcao logger
		if rota.RequerAuth {
			router.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}
	}

	FileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", FileServer))
	return router

}
