package routes

import (
	"microservice_group/pkg/controller"
	"microservice_group/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.GroupServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		Group := main.Group("/v1")
		{
			Group.GET("/groups", func(c *gin.Context) {
				controller.GetGroups(c, service)
			})

		}
	}

	// retorna rota
	return router
}
