package routes

import (
	"microservice_spreadsheet/pkg/controller"
	"microservice_spreadsheet/pkg/middlewares"
	"microservice_spreadsheet/pkg/service"

	"github.com/gin-gonic/gin"
)

// ConfigRoutes recebe uma instância do gin.Engine e uma instância do service.RemarkServiceInterface e configura as rotas para as requisições HTTP no servidor.
func ConfigRoutes(router *gin.Engine, service service.RemarkServiceInterface) *gin.Engine {

	router.Use(middlewares.CORS())

	main := router.Group("union")
	{
		remarks := main.Group("/v1")
		{
			// GET /remarks/submissives/:user_ID: retorna uma lista de todas as avaliações submetidas pelo usuário com o ID especificado na URL, disparando o método controller.GetSubmissiveRemarks.
			remarks.POST("/spreadsheet/generate", middlewares.Auth(), func(c *gin.Context) {
				controller.GenerateSpreadSheet(c, service)
			})

		}

	}

	// retorna rota
	return router

}
