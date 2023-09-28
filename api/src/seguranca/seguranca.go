package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e coloca um Hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//VerificarSenha compara uma senha com um Hash e verifica se sao iguais
func VerificarSenha(senhaComHash, senhastring string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhastring))
}
