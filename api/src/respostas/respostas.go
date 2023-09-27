package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON recebe o status code da resposta, adiciona esse codigo ao WriteHeader, recebe tambem os dados passados pelos parametros e convertem ao JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
