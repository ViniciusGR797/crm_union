package routes

import (
	"microservice_client/pkg/controller"
	"microservice_client/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.ClientServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		clients := main.Group("/v1")
		{
			clients.GET("/clients/mygroups/:user_id", func(c *gin.Context) {
				controller.GetClientsMyGroups(c, service)
			})
			clients.GET("/clients/id/:client_id", func(c *gin.Context) {
				controller.GetClientByID(c, service)
			})
			clients.GET("/clients/tag/:client_id", func(c *gin.Context) {
				controller.GetTagsClient(c, service)
			})
			clients.POST("/clients/", func(c *gin.Context) {
				controller.CreateClient(c, service)
			})
			clients.PUT("/clients/update/:client_id", func(c *gin.Context) {
				controller.UpdateClient(c, service)
			})
			clients.PUT("/clients/update/status/:client_id", func(c *gin.Context) {
				controller.UpdateStatusClient(c, service)
			})
		}
	}

	return router
}
