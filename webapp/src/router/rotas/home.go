package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotaPaginaPrincipal = Rota{
	URI:        "/home",
	Metodo:     http.MethodGet,
	Funcao:     controllers.CarregarPaginaHome,
	RequerAuth: true,
}
