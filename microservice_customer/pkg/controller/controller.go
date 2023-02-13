package controller

import (
	"fmt"
	"microservice_customer/pkg/entity"
	"microservice_customer/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetCustomer do service e retorna json com lista
func GetAllCustomer(c *gin.Context, service service.CustomerServiceInterface) {

	lista := service.GetAllCustomer()
	if len(lista.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
	fmt.Printf("tudo certo")
	c.JSON(200, lista)
}

// buscar customer por ID
func GetCustomerByID(c *gin.Context, service service.CustomerServiceInterface) {

	id := c.Param("id")

	newID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	produto := service.GetCustomerByID(&newID)
	if produto.ID == 0 {
		c.JSON(404, gin.H{
			"error": "produto not found, 404",
		})
		return
	}

	c.JSON(200, produto)

}

func CreateCustomer(c *gin.Context, service service.CustomerServiceInterface) {

	var customer *entity.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON customer" + err.Error(),
		})
		return
	}

	id := service.CreateCustomer(customer)
	if id == 0 {
		c.JSON(400, gin.H{
			"error": "cannot create JSON: " + err.Error(),
		})
	}

	customer = service.GetCustomerByID(&id)
	c.JSON(200, customer)
}

func UpdateCustomer(c *gin.Context, service service.CustomerServiceInterface) {
	id := c.Param("id")

	var customer *entity.Customer

	newID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer, 400" + err.Error(),
		})
		return
	}

	err = c.ShouldBind(&customer)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bin JSON customer, 400" + err.Error(),
		})
		return
	}

	idResult := service.UpdateCustomer(&newID, customer)
	if idResult == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	customer = service.GetCustomerByID(&newID)
	c.JSON(200, customer)

}
