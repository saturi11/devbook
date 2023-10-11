package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"net/http"
)

// Autenticar verifica s eo ususario fazendo a req esta autendicado
func Autenticar(proximaFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFunc(w, r)
	}
}
