package middlewares

import (
	"fmt"
	"microservice_release/pkg/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Auth verifica se o cabeçalho de autorização contém um token JWT válido para autenticar um usuário.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := header[len(bearer_schema):]

		err := security.ValidateToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := header[len(bearer_schema):]

		err := security.ValidateToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// pega permissões do token
		permissions, err := security.GetPermissions(c)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// Pega level nas permissões do token
		level, err := strconv.Atoi(fmt.Sprint(permissions["level"]))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Verifica se o user é um admin (level acima de 1)
		if level <= 1 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
