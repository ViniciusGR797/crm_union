package routes

import (
	"microservice_remark/pkg/controller"
	"microservice_remark/pkg/middlewares"
	"microservice_remark/pkg/service"

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
			remarks.GET("/remarks/submissives", middlewares.Auth(), func(c *gin.Context) {
				controller.GetSubmissiveRemarks(c, service)
			})
			// GET /remarks/user/id/:remark_id retorna uma lista de todos os Remarks do User especificado na URL, disparando o método controller.GetAllRemarkUser.
			remarks.GET("/remarks/user/id/:remark_id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetAllRemarkUser(c, service)
			})
			// GET /remarks/id/:remark_id: retorna uma avaliação com o ID especificado na URL, disparando o método controller.GetRemarkByID.
			remarks.GET("/remarks/id/:remark_id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetRemarkByID(c, service)
			})
			// POST /remarks: cria uma nova avaliação, disparando o método controller.CreateRemark.
			remarks.POST("/remarks", middlewares.Auth(), func(c *gin.Context) {
				controller.CreateRemark(c, service)
			})
			// GET /remarks/barchart/:user_ID: retorna um gráfico de barras mostrando a contagem de avaliações em relação ao tempo (atrasado, próximo, no prazo) para o usuário com o ID especificado na URL, disparando o método controller.GetBarChartRemark.
			remarks.GET("/remarks/barchart/:user_ID", middlewares.Auth(), func(c *gin.Context) {
				controller.GetBarChartRemark(c, service)
			})
			// GET /remarks/piechart/:user_ID: retorna um gráfico de pizza mostrando a contagem de avaliações em relação ao status (pendente, aprovado, rejeitado) para o usuário com o ID especificado na URL, disparando o método controller.GetPieChartRemark.
			remarks.GET("/remarks/piechart/:user_ID", middlewares.Auth(), func(c *gin.Context) {
				controller.GetPieChartRemark(c, service)
			})
			// PUT /remarks/update/status/:remark_id: atualiza o status de uma avaliação com o ID especificado na URL, disparando o método controller.UpdateStatusRemark.
			remarks.PUT("/remarks/update/status/:remark_id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateStatusRemark(c, service)
			})
			// PUT /remarks/update/:remark_id: atualiza uma avaliação com o ID especificado na URL, disparando o método controller.UpdateRemark.
			remarks.PUT("/remarks/update/:remark_id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateRemark(c, service)
			})

		}

	}

	// retorna rota
	return router

}
