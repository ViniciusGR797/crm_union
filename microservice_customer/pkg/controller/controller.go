package controller

import (
	"microservice_customer/pkg/entity"
	"microservice_customer/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetCustomer do service e retorna json com lista
func GetAllCustomer(c *gin.Context, service service.CustomerServiceInterface) {

	lista, err := service.GetAllCustomer()
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}
	c.JSON(http.StatusOK, lista)
}

// buscar customer por ID
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

func CreateCustomer(c *gin.Context, service service.CustomerServiceInterface) {

	var customer *entity.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = service.CreateCustomer(customer)
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

	err = service.UpdateCustomer(&newID, customer)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)

}

func SoftDeleteCustomer(c *gin.Context, service service.CustomerServiceInterface) {
	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return

	}

	err = service.SoftDeleteCustomer(&newID)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Customer Status Updated",
	})
}
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
