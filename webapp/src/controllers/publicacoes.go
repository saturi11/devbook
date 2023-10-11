package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarPublicacao cria uma publicacoes no banco
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/publicacoes", config.ApiUrl)
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
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

// CurtirPublicacao chama a API para curtir uma publicação
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	//extrai o id da publicacao pelo parametro da req
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.ApiUrl, publicacaoID)
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

// DescurtirPublicacao chama a API para descurtir uma publicação
func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	//extrai o id da publicacao pelo parametro da req
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/descurtir", config.ApiUrl, publicacaoID)
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

// AtualizarPublicacao chama a API para atualizar uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	//extrai o id da publicacao pelo parametro da req
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//extrai o corpo da req pelo parse form fazendo um map nos valores passados pelo formulario
	r.ParseForm()
	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//constroi a url da api que vai realmente realizar a acao, passando a url base pela variavel de ambiente e o id da publicacao jogado na variavel publicacaoId
	url := fmt.Sprintf("%s/publicacoes/%d", config.ApiUrl, publicacaoID)
	response, erro := requisicoes.FazerReqComAuth(r, http.MethodPut, url, bytes.NewBuffer(publicacao))
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

// DeletarPublicacao chama a API para deletar uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	//extrai o id da publicacao pelo parametro da req
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//constroi a url da api que vai realmente realizar a acao, passando a url base pela variavel de ambiente e o id da publicacao jogado na variavel publicacaoId
	url := fmt.Sprintf("%s/publicacoes/%d", config.ApiUrl, publicacaoID)

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
