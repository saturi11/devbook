package modelos

import "time"

//Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`       //omitempty exclui o campo em caso de id vazio
	Nome     string    `json:"nome,omitempty"`     //omitempty exclui o campo em caso de nome vazio
	Nick     string    `json:"nick,omitempty"`     //omitempty exclui o campo em caso de nick vazio
	Email    string    `json:"email,omitempty"`    //omitempty exclui o campo em caso de email vazio
	Senha    string    `json:"senha,omitempty"`    //omitempty exclui o campo em caso de senha vazia
	CriadoEm time.Time `json:"CriadoEm,omitempty"` //omitempty exclui o campo em caso de CriadoEm vazio
}
