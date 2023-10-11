package modelos

//DadosAutenticacao contem o token e o id do user autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`    //omitempty exclui o campo em caso de id vazio
	Token string `json:"Token"` //omitempty exclui o campo em caso de nome vazio
}
