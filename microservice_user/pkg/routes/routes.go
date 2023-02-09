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
			// Rota que retorna lista de users submissos do seus grupos (GET que dispara método GetSubmissiveUsers controller)
			produtos.GET("/users/submissives/:user_id", func(c *gin.Context) {
				controller.GetSubmissiveUsers(c, service)
			})
			// Rota que cadastra usuários (POST que dispara método CreateUser controller)
			produtos.POST("/users", func(c *gin.Context) {
				controller.CreateUser(c, service)
			})
			// Rota que altera status (PUT que dispara método UpdateStatusUser controller)
			produtos.PUT("/users/update/status/:user_id", func(c *gin.Context) {
				controller.UpdateStatusUser(c, service)
			})
			// Rota que retorna usuário editado (PUT que dispara método Update controller)
			produtos.PUT("/users/update/:user_id", func(c *gin.Context) {
				controller.UpdateUser(c, service)
			})
		}
	}
	// retorna rota
	return router
}
