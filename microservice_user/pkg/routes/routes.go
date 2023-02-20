package routes

import (
	"microservice_user/pkg/controller"
	"microservice_user/pkg/middlewares"
	"microservice_user/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.UserServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		user := main.Group("/v1")
		{
			// Rota que retorna lista de users (GET que dispara método GetUsers controller)
			user.GET("/users", middlewares.Auth(), func(c *gin.Context) {
				controller.GetUsers(c, service)
			})
			// Rota que retorna user pelo ID (GET que dispara método GetUserByID controller)
			user.GET("/users/id/:user_id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetUserByID(c, service)
			})
			// Rota que retorna users pelo nome (GET que dispara método GetUserByName controller)
			user.GET("/users/name/:user_name", middlewares.Auth(), func(c *gin.Context) {
				controller.GetUserByName(c, service)
			})
			// Rota que retorna lista de users submissos do seus grupos (GET que dispara método GetSubmissiveUsers controller)
			user.GET("/users/submissives", middlewares.Auth(), func(c *gin.Context) {
				controller.GetSubmissiveUsers(c, service)
			})
			// Rota que cadastra user (POST que dispara método CreateUser controller)
			user.POST("/users", middlewares.Auth(), func(c *gin.Context) {
				controller.CreateUser(c, service)
			})
			// Rota que altera status ativo/inativo (PUT que dispara método UpdateStatusUser controller)
			user.PUT("/users/update/status/:user_id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateStatusUser(c, service)
			})
			// Rota que edita user (PUT que dispara método UpdateUser controller)
			user.PUT("/users/update/:user_id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateUser(c, service)
			})
			// Rota de login de user (POST que dispara método Login controller)
			user.POST("/users/login", func(c *gin.Context) {
				controller.Login(c, service)
			})
			// Rota que retorna user pelo ID do meeu token (GET que dispara método GetUserMe controller)
			user.GET("/users/me", middlewares.Auth(), func(c *gin.Context) {
				controller.GetUserMe(c, service)
			})
		}
	}
	// retorna rota
	return router
}
