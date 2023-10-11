package modelos

//DadosAutenticacao contem o id e o token do user
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"Token"`
}
