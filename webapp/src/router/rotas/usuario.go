package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotasUsuario = []Rota{
	{
		URI:        "/criar-usuario",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/buscar-usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPerfilDoUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.PararDeSeguirUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioId}/seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/perfil",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPerfilDoUsuarioLogado,
		RequerAuth: true,
	},
	{
		URI:        "/editar-usuario",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPerfilDeEdicaoDeUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/editar-usuario",
		Metodo:     http.MethodPut,
		Funcao:     controllers.EditarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/atualizar-senha",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeAtualizacaoDeSenha,
		RequerAuth: true,
	},
	{
		URI:        "/atualizar-senha",
		Metodo:     http.MethodPost,
		Funcao:     controllers.EditarSenha,
		RequerAuth: true,
	},
	{
		URI:        "/deletar-usuario",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		RequerAuth: true,
	},
}
