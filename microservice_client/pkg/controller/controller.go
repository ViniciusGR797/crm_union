package controller

import (
	"microservice_client/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetUsers do service e retorna json com lista de users
func GetClientsMyGroups(c *gin.Context, service service.ClientServiceInterface) {

	ID := c.Param("user_id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	// Chama método GetUsers e retorna list de users
	list := service.GetClientsMyGroups(&newId)
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(200, list)
}

func UpdateStatusClient(c *gin.Context, service service.ClientServiceInterface) {
	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	result := service.UpdateStatusClient(&newID)
	if result == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"response": "Client Status Updated",
	})
}

func GetClientByID(c *gin.Context, service service.ClientServiceInterface) {

	ID := c.Param("client_id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	// Chama método GetUsers e retorna list de users
	client := service.GetClientByID(&newId)
	// Verifica se a lista está vazia (tem tamanho zero)
	if client.ID == 0 {
		c.JSON(404, gin.H{
			"error": "client not found, 404",
		})
		return
	}

	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(200, client)
}
