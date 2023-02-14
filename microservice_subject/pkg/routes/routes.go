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
		}
	}

	return router
}
