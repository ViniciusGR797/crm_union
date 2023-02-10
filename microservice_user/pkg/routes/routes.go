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
		users := main.Group("/v1")
		{
			// Rota que retorna lista de users (GET que dispara método GetUsers controller)
			users.GET("/users", func(c *gin.Context) {
				controller.GetUsers(c, service)
			})
			// Rota que retorna user pelo ID (GET que dispara método GetUserByID controller)
			users.GET("/users/id/:user_id", func(c *gin.Context) {
				controller.GetUserByID(c, service)
			})
			// Rota que retorna users pelo nome (GET que dispara método GetUserByName controller)
			users.GET("/users/name/:user_name", func(c *gin.Context) {
				controller.GetUserByName(c, service)
			})
			// Rota que retorna lista de users submissos do seus grupos (GET que dispara método GetSubmissiveUsers controller)
			users.GET("/users/submissives/:user_id", func(c *gin.Context) {
				controller.GetSubmissiveUsers(c, service)
			})
			// Rota que cadastra user (POST que dispara método CreateUser controller)
			users.POST("/users", func(c *gin.Context) {
				controller.CreateUser(c, service)
			})
			// Rota que altera status ativo/inativo (PUT que dispara método UpdateStatusUser controller)
			users.PUT("/users/update/status/:user_id", func(c *gin.Context) {
				controller.UpdateStatusUser(c, service)
			})
			// Rota que edita user (PUT que dispara método UpdateUser controller)
			users.PUT("/users/update/:user_id", func(c *gin.Context) {
				controller.UpdateUser(c, service)
			})
			// Rota de login de user (POST que dispara método Login controller)
			users.POST("/users/login", func(c *gin.Context) {
				controller.Login(c, service)
			})
		}
	}
	// retorna rota
	return router
}
