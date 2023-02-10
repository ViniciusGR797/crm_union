package routes

import (
	"microservice_customer/pkg/controller"
	"microservice_customer/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.CustomerServiceInterface) *gin.Engine {
	main := router.Group("api")
	{
		customer := main.Group("/v1")
		{
			// Rota que retorna lista de cuystomers (GET que dispara método GetAllCustomer controller)
			customer.GET("/Allcustomer", func(c *gin.Context) {
				controller.GetAllCustomer(c, service)
			})
			customer.GET("/customer/:id", func(c *gin.Context) {
				controller.GetCustomerByID(c, service)
			})
			customer.POST("/CreateCustomer/", func(c *gin.Context) {
				controller.CreateCustomer(c, service)
			})
		}
	}

	// retorna rota
	return router
}
