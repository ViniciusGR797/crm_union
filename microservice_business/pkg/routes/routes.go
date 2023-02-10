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

		}
	}

	// retorna rota
	return router
}
