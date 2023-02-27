package routes

import (
	"microservice_business/pkg/controller"
	"microservice_business/pkg/middlewares"
	"microservice_business/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.BusinessServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		Business := main.Group("/v1")
		{
			Business.GET("/business", middlewares.Auth(), func(c *gin.Context) {
				controller.GetBusiness(c, service)
			})
			Business.GET("/business/id/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.GetBusinessById(c, service)
			})
			Business.POST("/business", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.CreateBusiness(c, service)
			})
			Business.PUT("/business/update/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.UpdateBusiness(c, service)
			})
			Business.PUT("/business/update/status/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.UpdateStatusBusiness(c, service)
			})
			Business.GET("/business/name/:Business_name", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.GetBusinessByName(c, service)
			})
			Business.GET("/business/tag/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.GetTagsBusiness(c, service)
			})
		}
	}

	// retorna rota
	return router
}
