package controller

import (
	"errors"
	"fmt"
	"microservice_client/pkg/entity"
	"microservice_client/pkg/security"
	"microservice_client/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetClientsMyGroups: Retorna json com lista de clients
func GetClientsMyGroups(c *gin.Context, service service.ClientServiceInterface) {
	// Pega permissões do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id passada como token na rota
	id, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Chama método GetUsers e retorna list de users
	list, err := service.GetClientsMyGroups(&id)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(list.List) == 0 {
		JSONMessenger(c, http.StatusNotFound, c.Request.URL.Path, errors.New("clients not found for group"))
		return
	}
	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(http.StatusOK, list)
}

// GetTagsClient: Retorna json com uma lista de tags do client
func GetTagsClient(c *gin.Context, service service.ClientServiceInterface) {
	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	tags, err := service.GetTagsClient(&newID)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	if len(tags) == 0 {
		JSONMessenger(c, http.StatusNotFound, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GetClientByID: Retorna json com um client
func GetClientByID(c *gin.Context, service service.ClientServiceInterface) {

	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	// Chama método GetUsers e retorna list de users
	client, err := service.GetClientByID(&newID)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}
	// Verifica se a lista está vazia (tem tamanho zero)
	if client.ID == 0 {
		JSONMessenger(c, http.StatusNotFound, c.Request.URL.Path, err)
		return
	}

	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(http.StatusOK, client)
}

// CreateClient: Chama o serviço para criar um client
func CreateClient(c *gin.Context, service service.ClientServiceInterface) {
	var client entity.ClientUpdate

	if err := c.ShouldBindJSON(&client); err != nil {
		JSONMessenger(c, http.StatusUnprocessableEntity, c.Request.URL.Path, err)
		return
	}

	err := service.CreateClient(&client)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// UpdateClient: Chama o serviço para atualizar um client
func UpdateClient(c *gin.Context, service service.ClientServiceInterface) {
	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	var client entity.ClientUpdate

	if err := c.ShouldBindJSON(&client); err != nil {
		JSONMessenger(c, http.StatusUnprocessableEntity, c.Request.URL.Path, err)
		return
	}

	err = service.UpdateClient(&newID, &client)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, gin.H{
		"result": "Client updated successfully",
	})
}

// UpdateStatusClient: Realiza o softdelete em um client
func UpdateStatusClient(c *gin.Context, service service.ClientServiceInterface) {
	ID := c.Param("client_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = service.UpdateStatusClient(&newID)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Client Status Updated",
	})
}

// JSONMessenger: Estrutura o erro recebido
func JSONMessenger(c *gin.Context, status int, path string, err error) {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	c.JSON(status, gin.H{
		"status":  status,
		"message": errorMessage,
		"error":   err,
		"path":    path,
	})
}
