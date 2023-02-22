package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendError organiza um JSON para responder ao client em caso de erro
func sendError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"status":     status,
		"statusText": http.StatusText(status),
		"error":      err.Error(),
		"path":       c.FullPath(),
	})
}

// Send envia o objeto ao client em caso de sucesso na requisição
func send(c *gin.Context, code int, obj any) {
	c.JSON(code, obj)
}

// Send envia nada ao client em caso de sucesso na requisição
func sendNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
