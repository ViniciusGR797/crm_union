package routes

import (
	"microservice_customer/pkg/controller"
	"microservice_customer/pkg/service"
	"microservice_user/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.CustomerServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		customer := main.Group("/v1")
		{
			// Rota que retorna lista de cuystomers (GET que dispara método GetCustomers controller)
			customer.GET("/customers", middlewares.Auth(), func(c *gin.Context) {
				controller.GetCustomers(c, service)
			})
			customer.GET("/customers/id/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetCustomerByID(c, service)
			})
			customer.POST("/customers", middlewares.Auth(), func(c *gin.Context) {
				controller.CreateCustomer(c, service)
			})
			customer.PUT("/customers/update/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateCustomer(c, service)
			})
			customer.PUT("/customers/update/status/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateStatusCustomer(c, service)
			})

		}
	}

	// retorna rota
	return router
}
