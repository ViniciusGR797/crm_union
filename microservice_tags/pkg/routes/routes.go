package routes

import (
	"microservice_tags/pkg/controller"
	"microservice_tags/pkg/middlewares"
	"microservice_tags/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.TagsServiceInterface) *gin.Engine {

	router.Use(middlewares.CORS())

	main := router.Group("union")
	{
		Tags := main.Group("/v1")
		{
			Tags.GET("/tags", middlewares.Auth(), func(c *gin.Context) {
				controller.GetTags(c, service)
			})
			Tags.GET("/tags/id/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetTagsById(c, service)
			})
			Tags.GET("/domains", middlewares.Auth(), func(c *gin.Context) {
				controller.GetDomains(c, service)
			})
			Tags.GET("/domain/id/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetDomainById(c, service)
			})
		}
	}

	// retorna rota
	return router
}
