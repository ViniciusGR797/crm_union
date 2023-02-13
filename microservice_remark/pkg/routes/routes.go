package routes

import (
	"microservice_remark/pkg/controller"
	"microservice_remark/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.RemarkServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		remarks := main.Group("/v1")
		{
			// Rota que retorna lista de users (GET que dispara método GetUsers controller)
			remarks.GET("/remarks/submissives/:user_ID", func(c *gin.Context) {
				controller.GetSubmissiveRemarks(c, service)
			})
			remarks.GET("/remarks/id/:remark_id", func(c *gin.Context) {
				controller.GetRemarkByID(c, service)
			})
			remarks.POST("/remarks/ ", func(c *gin.Context) {
				controller.CreateRemark(c, service)
			})

		}

	}

	// retorna rota
	return router

}
