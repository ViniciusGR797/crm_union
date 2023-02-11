package security

import (
	"errors"
	"fmt"
	"microservice_user/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func NewToken(userID uint64, level uint) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	permissions["level"] = level

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.Secret))
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, errors.New("invalid token: " + token)
		}

		return []byte(config.Secret), nil
	})
	return err
}

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

func keyFunc(t *jwt.Token) (interface{}, error) {
	_, ok := t.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("invalid method: %v", t.Header["alg"])
	}

	return []byte(config.Secret), nil
}

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

// Função que verifica se o user é um Adm, se for retorna nil, senão retorna erro
func IsAdm(c *gin.Context) error {
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
	if level > 1 {
		return nil
	} else {
		return errors.New("admin exclusive route")
	}
}
