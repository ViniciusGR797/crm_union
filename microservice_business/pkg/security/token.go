package security

import (
	"errors"
	"fmt"
	"microservice_business/config"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret []byte

func SecretConfig(config *config.Config) error {
	secret = []byte(config.Secret)
	if len(secret) == 0 {
		return errors.New("env token secret not set")
	}
	return nil
}

// ValidateToken recebe uma string token como argumento e verifica se o token é válido ou não.
func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if !isValid {

			return nil, errors.New("invalid token: " + token)
		}

		return secret, nil
	})
	if err != nil {
		return err
	}

	err = IsActive(token)

	return err
}

// ExtractToken tem como objetivo extrair as informações do token JWT e retorná-las como um mapa de reivindicações (jwt.MapClaims).
func ExtractToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get permissions")
	}

	return permissions, nil
}

// keyFunc é utilizada para verificar se o método de assinatura do token JWT é do tipo HMAC e se a chave secreta utilizada para assinar e verificar a integridade do token é a mesma definida na configuração da aplicação.
func keyFunc(t *jwt.Token) (interface{}, error) {
	_, ok := t.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("invalid method: %v", t.Header["alg"])
	}

	return secret, nil
}

// GetToken recebe um objeto gin.Context, que é usado para acessar os headers da requisição.
func GetToken(c *gin.Context) (string, error) {
	const bearer_schema = "Bearer "
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("empty header")
	}

	token := header[len(bearer_schema):]

	err := ValidateToken(token)
	if err != nil {
		return "", errors.New("invalid token")
	}

	return token, nil
}

// GetPermissions é responsável por extrair as permissões do token presente no header de autorização de uma requisição.
func GetPermissions(c *gin.Context) (jwt.MapClaims, error) {
	token, err := GetToken(c)
	if err != nil {
		return nil, err
	}

	permissions, err := ExtractToken(token)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// IsActive verifica se o user é ATIVO, se for retorna nil, senão retorna erro
func IsActive(token string) error {
	// pega permissões do token
	permissions, err := ExtractToken(token)
	if err != nil {
		return errors.New("error getting permissions")
	}
	// Pega status nas permissões do token
	status := fmt.Sprint(permissions["status"])

	// Verifica se o user é está ativo
	if status == "ACTIVE" {
		return nil
	} else {
		return errors.New("inactive user")
	}
}

// Função que verifica se o user é um Adm, se for retorna nil, senão retorna erro
func IsUser(c *gin.Context) error {
	// pega permissões do token
	permissions, err := GetPermissions(c)
	if err != nil {
		return errors.New("error getting permissions")
	}
	// Pega level nas permissões do token
	level, err := strconv.Atoi(fmt.Sprint(permissions["level"]))
	if err != nil {
		return errors.New("conversation error")
	}

	// Verifica se o user é um admin (level acima de 1)
	if level == 1 {
		return nil
	} else {
		return errors.New("admin exclusive route")
	}
}
