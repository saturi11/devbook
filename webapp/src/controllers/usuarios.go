package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarUsuario chama a api para cadastrar o usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
	}

	url := fmt.Sprintf("%s/usuarios", config.ApiUrl)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// PararDeSeguirUsuario chama a api para parar de seguir o ususario
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	//extrai o id da publicacao pelo parametro da req
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//constroi a url da api que vai realmente realizar a acao, passando a url base pela variavel de ambiente e o id da publicacao jogado na variavel publicacaoId
	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.ApiUrl, usuarioID)

	//chama a funcao que realiza requisicoes com autenticacoes, passando a requisicao, o metodo, a url, e os dados
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// SeguirUsuario chama a api para seguir o ususario
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	//extrai o id da publicacao pelo parametro da req
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//constroi a url da api que vai realmente realizar a acao, passando a url base pela variavel de ambiente e o id da publicacao jogado na variavel publicacaoId
	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.ApiUrl, usuarioID)

	//chama a funcao que realiza requisicoes com autenticacoes, passando a requisicao, o metodo, a url, e os dados
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	// le os valors do form e monta um user com essas informacoes
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//ler o cookie e extrai o id atravez dele
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//constroi a url da api que vai realmente realizar a acao, passando a url base pela variavel de ambiente e o id da publicacao jogado na variavel publicacaoId
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)
	//chama a funcao que realiza requisicoes com autenticacoes, passando a requisicao, o metodo, a url, e os dados
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodPut, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func EditarSenha(w http.ResponseWriter, r *http.Request) {
	// le os valors do form e monta um user com essas informacoes
	r.ParseForm()
	senhas, erro := json.Marshal(map[string]string{
		"nova":  r.FormValue("nova"),
		"atual": r.FormValue("atual"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//ler o cookie e extrai o id atravez dele
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//constroi a url da api que vai realmente realizar a acao, passando a url base pela variavel de ambiente e o id da publicacao jogado na variavel publicacaoId
	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.ApiUrl, usuarioID)

	//chama a funcao que realiza requisicoes com autenticacoes, passando a requisicao, o metodo, a url, e os dados
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodPost, url, bytes.NewBuffer(senhas))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// DeletarUsuario Exclui o usuario do sistema
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	//ler o cookie e extrai o id atravez dele
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//constroi a url da api que vai realmente realizar a acao, passando a url base pela variavel de ambiente e o id da publicacao jogado na variavel publicacaoId
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)

	//chama a funcao que realiza requisicoes com autenticacoes, passando a requisicao, o metodo, a url, e os dados
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
