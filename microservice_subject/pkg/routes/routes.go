package routes

import (
	"microservice_subject/pkg/controller"
	"microservice_subject/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.SubjectServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		subject := main.Group("/v1")
		{
			subject.GET("/subjects/user/:id", func(c *gin.Context) {
				controller.GetSubjectList(c, service)
			})

			subject.GET("/subjects/:id", func(c *gin.Context) {
				controller.GetSubject(c, service)
			})

			subject.PUT("/subjects/update/finished/:id", func(c *gin.Context) {
				controller.UpdateStatusSubjectFinished(c, service)
			})

			subject.PUT("/subjects/update/canceled/:id", func(c *gin.Context) {
				controller.UpdateStatusSubjectCanceled(c, service)
			})

			subject.POST("/subjects/create/user/:id", func(c *gin.Context) {
				controller.CreateSubject(c, service)
			})
		}
	}

	return router
}
