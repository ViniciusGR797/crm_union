package routes

import (
	"microservice_group/pkg/controller"
	"microservice_group/pkg/middlewares"
	"microservice_group/pkg/service"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, service service.GroupServiceInterface) *gin.Engine {

	router.Use(middlewares.CORS())
	main := router.Group("union")
	{
		Group := main.Group("/v1")
		{
			Group.GET("/groups/user/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.GetGroups(c, service)
			})

			Group.GET("/groups/id/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.GetGroupByID(c, service)
			})

			Group.PUT("/groups/update/status/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.UpdateStatusGroup(c, service)
			})

			Group.GET("groups/usersGroup/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.GetUsersGroup(c, service)
			})

			Group.POST("/groups", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.CreateGroup(c, service)
			})

			Group.PUT("/groups/update/attach/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.AttachUserGroup(c, service)
			})

			Group.PUT("/groups/update/detach/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.DetachUserGroup(c, service)

			})
			Group.PUT("/groups/update/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.EditGroup(c, service)
			})
			Group.GET("/groups/count/user/:id", middlewares.AuthAdmin(), func(c *gin.Context) {
				controller.CountUsersGroup(c, service)
			})

		}
	}

	// retorna rota
	return router
}
