package routes

import (
	"microservice_subject/pkg/controller"
	"microservice_subject/pkg/middlewares"
	"microservice_subject/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.SubjectServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		subject := main.Group("/v1")
		{
			subject.GET("/subjects/submissives", middlewares.Auth(), func(c *gin.Context) {
				controller.GetSubmissiveSubjects(c, service)
			})
			subject.GET("/subjects/id/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetSubjectByID(c, service)
			})

			subject.PUT("/subjects/update/finished/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateStatusSubjectFinished(c, service)
			})

			subject.PUT("/subjects/update/canceled/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateStatusSubjectCanceled(c, service)
			})

			subject.PUT("/subjects/update/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdateSubject(c, service)
			})

			subject.POST("/subjects/create/user/:id", func(c *gin.Context) {
				controller.CreateSubject(c, service)
			})
		}
	}

	return router
}
