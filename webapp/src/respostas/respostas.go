package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI representa a resposta de erro da api
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

// TratarStatusCodeDeErro trata as req com status de erro > 400
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
