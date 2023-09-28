package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`       //omitempty exclui o campo em caso de id vazio
	Nome     string    `json:"nome,omitempty"`     //omitempty exclui o campo em caso de nome vazio
	Nick     string    `json:"nick,omitempty"`     //omitempty exclui o campo em caso de nick vazio
	Email    string    `json:"email,omitempty"`    //omitempty exclui o campo em caso de email vazio
	Senha    string    `json:"senha,omitempty"`    //omitempty exclui o campo em caso de senha vazia
	CriadoEm time.Time `json:"CriadoEm,omitempty"` //omitempty exclui o campo em caso de CriadoEm vazio
}

// Preparar chama as funcoes de validamento e formatacao dos dados do usuario
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.Validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) Validar(etapa string) error {

	if usuario.Nome == "" {
		return errors.New("O nome e obrigatorio e nao pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick e obrigatorio e nao pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O Email e obrigatorio e nao pode estar em branco")
	}
	//valida o formato do email com o pacote checkmail
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o e-mail inserido e invalido")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha e obrigatorio e nao pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
