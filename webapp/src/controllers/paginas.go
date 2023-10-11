package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaDeLogin carrega a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	//chama a funcao de leitura dos cookies passando a req
	cookie, _ := cookies.Ler(r)
	//verifica a existencia de um token, se existir, o usuario sera redirecionado para tela home, caso contrario, a tela de login sera chamada
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario carrega a tela de cadastro
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaHome carrega a tela principal
func CarregarPaginaHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.ApiUrl)
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// CarregarPaginaDeAtualizacaoDePublicacao carrega a página de edição de publicação
func CarregarPaginaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	//obtem os parametros da req atravez do mux, e atribui para uma variavel
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//monta a variavel url, passando a url padrao da api, e o id obtido pelos parametros da req
	url := fmt.Sprintf("%s/publicacoes/%d", config.ApiUrl, publicacaoID)
	//chama o metodo que verifica a autenticacao do token, passando a req, o metodo, a url, e os dados que serao mandados para api
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	//cria uma instancia de publicacoes atravez do modelo de publis, e povoa a variavel com os dados obtidos pela api
	var publicacao modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//chama o template passando o writer, o arquivo html, e os dados obtidos na api
	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)
}

// CarregarPaginaDeUsuario carrega a pagina com os uruarios buscados
func CarregarPaginaDeUsuario(w http.ResponseWriter, r *http.Request) {
	//pega o parametro passado pela url e atribui para a variavel
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	//monta a url passando a url pafrao da api e o parametro na query
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.ApiUrl, nomeOuNick)

	//chama o metodo que verifica a autenticacao do token, passando a req, o metodo, a url, e os dados que serao mandados para api
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	//cria uma instancia de usuarios atravez do modelo de usuarios, e povoa a variavel com os dados obtidos pela api
	var usuarios []modelos.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chama o template passando o writer, o arquivo html, e os dados obtidos na api
	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario carrega a pagina de perfil dos usuarios
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	//obtem os parametros da req atravez do mux, e atribui para uma variavel
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//verifica se o usuarioId e o mesmo do usuario logado, se for, redireciona para tela de perfil do usuario logado
	if usuarioID == usuarioLogadoID {
		http.Redirect(w, r, "/perfil", 302)
		return
	}
	//Chama a funcao do Modelo que monta o usuario com todas as infos
	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         modelos.Usuario
		UsuarioLogadoID uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoID: usuarioLogadoID,
	})

}

// CarregarPerfilDoUsuario carrega a pagina de perfil dos usuarios Logado
func CarregarPerfilDoUsuarioLogado(w http.ResponseWriter, r *http.Request) {
	//ler o cookie e extrai o id atravez dele
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//Chama a funcao do Modelo que monta o usuario com todas as infos
	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)

}

// CarregarPerfilDeEdicaoDeUsuario carrega  a pagina para a edicao do usuario
func CarregarPerfilDeEdicaoDeUsuario(w http.ResponseWriter, r *http.Request) {
	//ler o cookie e extrai o id atravez dele
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	canal := make(chan modelos.Usuario)
	go modelos.BuscarDadosDoUsuario(canal, usuarioID, r)
	usuario := <-canal

	if usuario.ID == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao buscar usuario"})
		return
	}
	utils.ExecutarTemplate(w, "editar-usuario.html", usuario)
}

// CarregarPaginaDeAtualizacaoDeSenha carrega  a pagina para a edicao da senha do usuario
func CarregarPaginaDeAtualizacaoDeSenha(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "atualizar-senha.html", nil)

}
