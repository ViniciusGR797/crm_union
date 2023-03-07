package routes

import (
	"microservice_customer/pkg/controller"
	"microservice_customer/pkg/middlewares"
	"microservice_customer/pkg/service"

	"github.com/gin-gonic/gin"
)

// ConfigRoutes Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.CustomerServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		customer := main.Group("/v1")
		{
			// GET /customers: Rota que retorna lista de cuystomers (GET que dispara método GetCustomers controller)
			customer.GET("/customers", middlewares.Auth(), func(c *gin.Context) {
				controller.GetCustomers(c, service)
			})
			// GET /customers/id/:id: Retorna um cliente específico a partir do seu ID. A rota dispara o método GetCustomerByID do controlador controller.
			customer.GET("/customers/id/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.GetCustomerByID(c, service)
			})
			// POST /customers: Cria um novo cliente na base de dados. A rota dispara o método CreateCustomer do controlador controller.
			customer.POST("/customers", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.CreateCustomer(c, service)
			})
			// PUT /customers/update/:id: Atualiza as informações de um cliente existente na base de dados. A rota dispara o método UpdateCustomer do controlador controller.
			customer.PUT("/customers/update/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.UpdateCustomer(c, service)
			})
			// PUT /customers/update/status/:id: Atualiza o status de um cliente existente na base de dados. A rota dispara o método UpdateStatusCustomer do controlador controller.
			customer.PUT("/customers/update/status/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.UpdateStatusCustomer(c, service)
			})

		}
	}

	// retorna rota
	return router
}
