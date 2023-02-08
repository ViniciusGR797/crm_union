package security

import "golang.org/x/crypto/bcrypt"

// HashPassword - recebe uma senha e retorna um hash dela
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword - valida se o hash condiz com a senha
func ValidatePassword(senhaComHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senha))
}
