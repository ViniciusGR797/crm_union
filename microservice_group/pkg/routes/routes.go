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
			Group.GET("/groups/user/:id", func(c *gin.Context) {
				controller.GetGroups(c, service)
			})

			Group.GET("/groups/:id", func(c *gin.Context) {
				controller.GetGroupByID(c, service)
			})

			Group.PUT("/groups/update/status/:id", func(c *gin.Context) {
				controller.SoftDelete(c, service)
			})
		}
	}

	// retorna rota
	return router
}
