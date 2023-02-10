package middlewares

import (
	"microservice_user/pkg/security"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(bearer_schema):]

		err := security.ValidateToken(token)
		if err != nil {
			c.AbortWithStatus(401)
		}
	}
}
