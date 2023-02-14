package routes

import (
	"microservice_customer/pkg/controller"
	"microservice_customer/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.CustomerServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		customer := main.Group("/v1")
		{
			// Rota que retorna lista de cuystomers (GET que dispara método GetAllCustomer controller)
			customer.GET("/customer/GetAll", func(c *gin.Context) {
				controller.GetAllCustomer(c, service)
			})
			customer.GET("/customer/:id", func(c *gin.Context) {
				controller.GetCustomerByID(c, service)
			})
			customer.POST("/Customer/create", func(c *gin.Context) {
				controller.CreateCustomer(c, service)
			})
			customer.PUT("/customer/update/:id", func(c *gin.Context) {
				controller.UpdateCustomer(c, service)
			})
			customer.PUT("/customer/SoftDelete/:id", func(c *gin.Context) {
				controller.SoftDeleteCustomer(c, service)
			})
		}
	}

	// retorna rota
	return router
}
