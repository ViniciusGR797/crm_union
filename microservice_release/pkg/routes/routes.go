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
			// Rota que retorna release pelo ID (GET que dispara método GetReleaseTrainByID controller)
			release.GET("/releasetrains/id/:releasetrain_id", func(c *gin.Context) {
				controller.GetReleaseTrainByID(c, service)
			})
			release.PUT("/releasetrains/update/:releasetrain_id", func(c *gin.Context) {
				controller.UpdateReleaseTrain(c, service)
			})
			release.GET("/releasetrains/tag/:releasetrain_id", func(c *gin.Context) {
				controller.GetTagsReleaseTrain(c, service)
			})
			// Rota que altera status ativo/inativo (PUT que dispara método UpdateStatusReleaseTrain controller)
			release.PUT("/releasetrains/update/status/:releasetrain_id", func(c *gin.Context) {
				controller.UpdateStatusReleaseTrain(c, service)
			})
			// Rota que retorna release pelo business ID (GET que dispara método GetReleaseTrainByBusiness controller)
			release.GET("/releasetrains/business/:business_id", func(c *gin.Context) {
				controller.GetReleaseTrainByBusiness(c, service)
			})
			// Rota que cadastra release (POST que dispara método CreateReleaseTrain controller)
			release.POST("/releasetrains", func(c *gin.Context) {
				controller.CreateReleaseTrain(c, service)
			})

		}
	}

	return router
}
