package controller

import (
	"microservice_planner/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetPlannerByID do service e retorna json com lista de users
func GetPlannerByID(c *gin.Context, service service.PlannerServiceInterface) {
	c.Status(http.StatusOK)
}
