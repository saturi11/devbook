package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexao é a string que realiza a conexão com o banco MySQL
	StringConexao = ""
	//porta onde a api está rodando
	Porta = 0
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT")) //converte string obtida atravez do arquivo .env para int e atribui a variavel Porta
	if erro != nil {
		Porta = 9000 // joga uma porta padrão caso o get env de erro
	}

	StringConexao = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", //string de conexão que vai ser formatada com os valores do getenv
		os.Getenv("DB_USUARIO"), //subistitui o primeiro %s
		os.Getenv("DB_SENHA"),   //subistitui o segundo %s
		os.Getenv("DB_NOME"),    //subistitui o terceiro %s
	)

}
