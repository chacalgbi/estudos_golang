package seguranca

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerificarSenha(senhaHash, textPlainSenha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(textPlainSenha))
}
