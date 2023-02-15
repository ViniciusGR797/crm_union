package controller

import (
	"microservice_planner/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetPlannerByID do service e retorna json com lista de users
func GetPlannerByID(c *gin.Context, service service.PlannerServiceInterface) {
	ID := c.Param("id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/planners/:id",
		})
		return
	}

	// Chama método GetUsers e retorna release
	planner, err := service.GetPlannerByID(&newId)
	// Verifica se a release está vazia
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/planners/:id",
		})
		return
	}

	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(http.StatusOK, planner)
}
