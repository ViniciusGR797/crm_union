package security

import (
	"errors"
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
