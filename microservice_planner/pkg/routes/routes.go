package routes

import (
	"microservice_planner/pkg/controller"
	"microservice_planner/pkg/service"

	"github.com/gin-gonic/gin"
)

// Função que configura todas as rotas da api
func ConfigRoutes(router *gin.Engine, service service.PlannerServiceInterface) *gin.Engine {
	main := router.Group("union")
	{
		planners := main.Group("/v1")
		{
			// Rota que retorna lista de planners (GET que dispara método GetPlannerByID controller)
			planners.GET("/planners/id/:id", func(c *gin.Context) {
				controller.GetPlannerByID(c, service)
			})
			planners.POST("/planners", func(c *gin.Context) {
				controller.CreatePlanner(c, service)
			})

		}
	}
	// retorna rota
	return router
}
