package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CriarUsuario insere um user no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoReq, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	//chama a funcao Criar do repositorio e passa para ela o usuario que veio da req
	repositorio := repositorios.NovoRepDeUsuarios(db)
	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioId)))
}

// BuscarUsuarios busca todos os users
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

// BuscarUsuario busca um user
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// AtualizarUsuario atualiza um user no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuário"))
}

// DeletarUsuario deleta um user no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}
