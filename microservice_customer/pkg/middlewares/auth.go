package middlewares

import (
	"microservice_customer/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth esta função ela verifica se o cabeçalho de autorização contém um token JWT válido para autenticar um usuário.
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
