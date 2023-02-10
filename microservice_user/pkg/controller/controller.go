package controller

import (
	"microservice_user/pkg/service"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
// Função que chama método GetAll do service e retorna json com lista de produtos
func GetUsers(c *gin.Context, service service.UserServiceInterface) {
	// Chama método GetAll e retorna list de products
	users := service.GetUsers()
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(*users) == 0 {
=======
// Função que chama método GetUsers do service e retorna json com lista de users
func GetUsers(c *gin.Context, service service.UserServiceInterface) {
	// Chama método GetUsers e retorna list de users
	list := service.GetUsers()
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(list.List) == 0 {
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
<<<<<<< HEAD
	//retorna sucesso 200 e retorna json da lista de products
	c.JSON(200, users)
=======
	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(200, list)
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
}
