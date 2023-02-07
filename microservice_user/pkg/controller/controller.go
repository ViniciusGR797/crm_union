package controller

import (
	"microservice_user/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetAll do service e retorna json com lista de produtos
func GetUsers(c *gin.Context, service service.UserServiceInterface) {
	// Chama método GetAll e retorna list de products
	users := service.GetUsers()
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(*users) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
	//retorna sucesso 200 e retorna json da lista de products
	c.JSON(200, users)
}
