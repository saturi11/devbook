package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout remove os dados de autenticacao salvos no navegador
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	//chama a funcao que deleta os dados dos cookies
	cookies.Deletar(w)
	//redireciona para a tela de login
	http.Redirect(w, r, "/login", 302)
}
