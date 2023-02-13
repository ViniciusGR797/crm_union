package security

import "golang.org/x/crypto/bcrypt"

// HashPassword - recebe uma senha e retorna um hash dela
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ValidatePassword - valida se o hash condiz com a senha
func ValidatePassword(senhaComHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senha))
}
