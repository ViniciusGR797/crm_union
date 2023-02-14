package routes

import (
	"microservice_business/pkg/controller"
	"microservice_business/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.BusinessServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		Business := main.Group("/v1")
		{
			Business.GET("/business", func(c *gin.Context) {
				controller.GetBusiness(c, service)
			})
			Business.GET("/business/:id", func(c *gin.Context) {
				controller.GetBusinessByID(c, service)
			})
			Business.POST("/business", func(c *gin.Context) {
				controller.CreateBusiness(c, service)
			})
			Business.PUT("/business/update/:id", func(c *gin.Context) {
				controller.UpdateBusiness(c, service)
			})
			Business.PUT("/business/update/status/:id", func(c *gin.Context) {
				controller.SoftDeleteBusiness(c, service)
			})
			Business.GET("/business/name/:Business_name", func(c *gin.Context) {
				controller.GetBusinessByName(c, service)
			})
		}
	}

	// retorna rota
	return router
}
