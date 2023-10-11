package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:        "/publicacoes",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}/curtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CurtirPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}/descurtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.DescurtirPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}/atualizar",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeAtualizacaoDePublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarPublicacao,
		RequerAuth: true,
	},
}
