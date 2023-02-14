package controller

import (
	"microservice_client/pkg/entity"
	"microservice_client/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetClientsMyGroups do service e retorna json com lista de clients
func GetClientsMyGroups(c *gin.Context, service service.ClientServiceInterface) {

	ID := c.Param("user_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	// Chama método GetUsers e retorna list de users
	list := service.GetClientsMyGroups(&newID)
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

// Função que chama método GetTagsClient do service e retorna json com uma lista de tags do client
func GetTagsClient(c *gin.Context, service service.ClientServiceInterface) {
	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	tags := service.GetTagsClient(&newID)
	if len(tags) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}

	c.JSON(200, tags)
}

// Função que chama método GetClientByID do service e retorna json com um client
func GetClientByID(c *gin.Context, service service.ClientServiceInterface) {

	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	// Chama método GetUsers e retorna list de users
	client := service.GetClientByID(&newID)
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

func CreateClient(c *gin.Context, service service.ClientServiceInterface) {
	var client entity.ClientUpdate

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := service.CreateClient(&client)

	c.JSON(200, id)
}

func UpdateClient(c *gin.Context, service service.ClientServiceInterface) {
	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	var client entity.ClientUpdate

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := service.UpdateClient(&newID, &client)
	if *result == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "Client updated successfully",
	})
}

// Função que chama método UpdateStatusClient do service e realiza o softdelete
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
