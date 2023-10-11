package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//ApiUrl representa a url de comunicacao para url
	ApiUrl = ""
	//Porta onde a api esta rodando
	Porta = 0
	//HashKey e utilizada para autenticar o cookie
	HashKey []byte
	//BlockKey e utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	ApiUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
