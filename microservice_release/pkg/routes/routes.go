package routes

import (
	"microservice_release/pkg/controller"
	"microservice_release/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.ReleaseServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		release := main.Group("/v1")
		{
			release.GET("/releasetrains", func(c *gin.Context) {
				controller.GetReleasesTrain(c, service)
			})

		}
	}

	return router
}
