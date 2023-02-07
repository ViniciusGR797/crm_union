package routes

import (
	"microservice_user/pkg/controller"
	"microservice_user/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.UserServiceInterface) *gin.Engine {
	main := router.Group("api")
	{
		user := main.Group("/v1")
		{
			// Rota que retorna lista de log (GET que dispara método GetLog controller)
			user.GET("/users", func(c *gin.Context) {
				controller.GetUsers(c, service)
			})

		}
	}

	// retorna rota
	return router
}
