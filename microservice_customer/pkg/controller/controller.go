package controller

import (
	"fmt"
	"microservice_customer/pkg/entity"
	"microservice_customer/pkg/security"
	"microservice_customer/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCustomers Função que chama método GetCustomer do service e retorna json com lista
func GetCustomers(c *gin.Context, service service.CustomerServiceInterface) {

	lista, err := service.GetCustomers()
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}
	c.JSON(http.StatusOK, lista)
}

// GetCustomerByID buscar customer por ID
func GetCustomerByID(c *gin.Context, service service.CustomerServiceInterface) {
	id := c.Param("id")

	newID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	customer, err := service.GetCustomerByID(&newID)
	if err != nil {
		JSONMessenger(c, http.StatusNotFound, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusOK, customer)

}

// CreateCustomer verifica se a rota e a função são exclusivas do administrador.
func CreateCustomer(c *gin.Context, service service.CustomerServiceInterface) {
	var customer *entity.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	err = service.CreateCustomer(customer, &logID, ctx)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func UpdateCustomer(c *gin.Context, service service.CustomerServiceInterface) {
	id := c.Param("id")

	var customer *entity.Customer

	newID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = c.ShouldBindJSON(&customer)
	if err != nil {
		JSONMessenger(c, http.StatusUnprocessableEntity, c.Request.URL.Path, err)
		return
	}

	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	err = service.UpdateCustomer(&newID, customer, &logID, ctx)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)

}

// UpdateCustomer é uma rota para atualizar um cliente existente. Primeiro, ele verifica se o usuário que faz a solicitação é um administrador usando a função security.IsAdm
func UpdateStatusCustomer(c *gin.Context, service service.CustomerServiceInterface) {
	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	err = service.UpdateStatusCustomer(&newID, &logID, ctx)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Customer Status Updated",
	})
}

// JSONMessenger é utilizada para formatar as respostas em JSON enviadas para o cliente.
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
