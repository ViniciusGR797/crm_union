package security

import (
	"errors"
	"fmt"
	"microservice_user/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewToken(userID uint64, level uint) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["level"] = level
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID

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
		return nil, errors.New(fmt.Sprintf("invalid method: %v", t.Header["alg"]))
	}

	return []byte(config.Secret), nil
}
