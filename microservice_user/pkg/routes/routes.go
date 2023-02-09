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
		produtos := main.Group("/v1")
		{
			// Rota que retorna lista de users (GET que dispara método GetUsers controller)
			produtos.GET("/users", func(c *gin.Context) {
				controller.GetUsers(c, service)
			})
			// Rota que retorna user pelo ID (GET que dispara método GetUserByID controller)
			produtos.GET("/users/id/:user_id", func(c *gin.Context) {
				controller.GetUserByID(c, service)
			})
			// Rota que retorna users pelo nome (GET que dispara método GetUserByName controller)
			produtos.GET("/users/name/:user_name", func(c *gin.Context) {
				controller.GetUserByName(c, service)
			})
			// Rota que retorna lista de users submissos do seus grupos (GET que dispara método GetUsers controller)
			produtos.GET("/users/submissives/:user_id", func(c *gin.Context) {
				controller.GetSubmissiveUsers(c, service)
			})
		}
	}

	// retorna rota
	return router
}
