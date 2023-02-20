package routes

import (
	"microservice_remark/pkg/controller"
	"microservice_remark/pkg/service"
	"microservice_user/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.RemarkServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		remarks := main.Group("/v1")
		{
			// Rota que retorna lista de users (GET que dispara método GetUsers controller)
			remarks.GET("/remarks/submissives", middlewares.Auth(), func(c *gin.Context) {
				controller.GetSubmissiveRemarks(c, service)
			})
			remarks.GET("/remarks/id/:remark_id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetRemarkByID(c, service)
			})
			remarks.POST("/remarks", middlewares.Auth(), func(c *gin.Context) {
				controller.CreateRemark(c, service)
			})
			remarks.GET("/remarks/barchart/:user_ID", middlewares.Auth(), func(c *gin.Context) {
				controller.GetBarChartRemark(c, service)
			})
			remarks.GET("/remarks/piechart/:user_ID", middlewares.Auth(), func(c *gin.Context) {
				controller.GetPieChartRemark(c, service)
			})
			remarks.PUT("/remarks/update/status/:remark_id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateStatusRemark(c, service)
			})
			remarks.PUT("/remarks/update/:remark_id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateRemark(c, service)
			})

		}

	}

	// retorna rota
	return router

}
