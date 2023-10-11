package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configurar utiliza as variaveis de ambiente para configurar o securecookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar salva as infos de login
func Salvar(w http.ResponseWriter, ID, token string) error {
	//cria os dados do cookie atravez do map
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}
	//codifica os dados
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}
	//escreve os dados no navegador
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

// Ler retorna os valores armazenados nos cookies
func Ler(r *http.Request) (map[string]string, error) {
	//verifica a existencia de coockies com o nome "dados"
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	//descriptograva os valores contidos no cookie lido
	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}
	return valores, nil

}

// Deletar limpa os valores dos cookies escritos no navegador
func Deletar(w http.ResponseWriter) {
	//Sobreescreve os dados no navegador com dados vazios
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
