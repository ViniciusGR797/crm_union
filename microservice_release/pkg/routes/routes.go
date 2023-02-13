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
			release.GET("/releasetrains/id/:releasetrain_id", func(c *gin.Context) {
				controller.GetReleaseTrainByID(c, service)
			})
			release.PUT("/releasetrains/update/:releasetrain_id", func(c *gin.Context) {
				controller.UpdateReleaseTrain(c, service)
			})
			release.GET("/releasetrains/tag/:releasetrain_id", func(c *gin.Context) {
				controller.GetTagsReleaseTrain(c, service)
			})
			// release.PUT("/releasetrains/update/tag/:releasetrain_id", func(c *gin.Context) {
			// 	controller.InsertTagsReleaseTrain(c, service)
			// })
			// Rota que altera status ativo/inativo (PUT que dispara método UpdateStatusReleaseTrain controller)
			release.PUT("/releasetrains/update/status/:releasetrain_id", func(c *gin.Context) {
				controller.UpdateStatusReleaseTrain(c, service)
			})

		}
	}

	return router
}
