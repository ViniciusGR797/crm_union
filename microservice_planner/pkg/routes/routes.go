package routes

import (
	"microservice_planner/pkg/controller"
	"microservice_planner/pkg/middlewares"
	"microservice_planner/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.PlannerServiceInterface) *gin.Engine {

	router.Use(middlewares.CORS())
	main := router.Group("union")
	{
		planner := main.Group("/v1")
		{
			// Rota que retorna lista de planners (GET que dispara método GetPlannerByID controller)
			planner.GET("/planners/id/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetPlannerByID(c, service)
			})
			planner.POST("/planners", middlewares.Auth(), func(c *gin.Context) {
				controller.CreatePlanner(c, service)
			})
			planner.GET("/planners/name/:name", middlewares.Auth(), func(c *gin.Context) {
				controller.GetPlannerByName(c, service)
			})
			// Rota que retorna lista de planners (GET que dispara método GetSubmissivePlanners controller)
			planner.GET("planners/submissives", middlewares.Auth(), func(c *gin.Context) {
				controller.GetSubmissivePlanners(c, service)
			})

			planner.GET("/planners/business/:business_name", middlewares.Auth(), func(c *gin.Context) {
				controller.GetPlannerByBusiness(c, service)
			})
			planner.GET("/planners/guest/client/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.GetGuestClientPlanners(c, service)
			})
			planner.PUT("/planners/update/:id", middlewares.Auth(), func(c *gin.Context) {
				controller.UpdatePlanner(c, service)
			})
		}
	}
	// retorna rota
	return router
}
