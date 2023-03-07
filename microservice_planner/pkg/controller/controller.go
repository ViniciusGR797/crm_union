package controller

import (
	"fmt"
	"microservice_planner/pkg/entity"
	"microservice_planner/pkg/security"
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

func CreatePlanner(c *gin.Context, service service.PlannerServiceInterface) {
	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Cria variável do tipo Planner (inicialmente vazia)
	var planner *entity.PlannerUpdate

	// Converte json em Planner
	err = c.ShouldBind(&planner)
	// Verifica se tem erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/planner",
		})
		return
	}

	err = service.CreatePlanner(planner, &logID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/planner",
		})
		return
	}

	// Retorno json com o Planner
	c.Status(http.StatusNoContent)
}

func GetPlannerByName(c *gin.Context, service service.PlannerServiceInterface) {
	// Pega name passada como parâmetro na URL da rota
	name := c.Param("name")

	// Pega permissões do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id passada como token na rota
	id, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Pega level passada como token na rota
	level, err := strconv.Atoi(fmt.Sprint(permissions["level"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Chama método GetSubmissivePlanners passando id e level como parâmetro
	list, err := service.GetPlannerByName(&id, level, &name)
	// Verifica se teve erro ao buscar planners no banco
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch planners",
		})
		return
	}
	// Verifica se a lista de planners tem tamanho zero (caso for user não tem planners submissive)
	if len(list.List) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	// Retorno json com planner
	c.JSON(http.StatusOK, list)

}

// Função que chama método GetSubmissivePlanners do service e retorna json com planners
func GetSubmissivePlanners(c *gin.Context, service service.PlannerServiceInterface) {
	// Pega permissões do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id passada como token na rota
	id, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Pega level passada como token na rota
	level, err := strconv.Atoi(fmt.Sprint(permissions["level"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Chama método GetSubmissivePlanners passando id e level como parâmetro
	list, err := service.GetSubmissivePlanners(&id, level)
	// Verifica se teve erro ao buscar planners no banco
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch planners",
		})
		return
	}
	// Verifica se a lista de planners tem tamanho zero (caso for user não tem planners submissive)
	if len(list.List) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	// Retorno json com planner
	c.JSON(http.StatusOK, list)
}

func GetPlannerByBusiness(c *gin.Context, service service.PlannerServiceInterface) {

	name := c.Param("name")

	planner, err := service.GetPlannerByBusiness(&name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/planners/business/:business_name",
		})
		return
	}

	c.JSON(http.StatusOK, planner)

}

func GetGuestClientPlanners(c *gin.Context, service service.PlannerServiceInterface) {
	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/business/guest/client/:id",
		})
		return
	}

	tags, err := service.GetGuestClientPlanners(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business/guest/client/:id",
		})
		return
	}

	c.JSON(http.StatusOK, tags)
}

func UpdatePlanner(c *gin.Context, service service.PlannerServiceInterface) {
	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/planners/update/:id",
		})
		return
	}

	// Cria variável do tipo Planner (inicialmente vazia)
	var planner *entity.PlannerUpdate

	// Converte json em Planner
	err = c.ShouldBind(&planner)
	// Verifica se tem erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/planner",
		})
		return
	}

	_, err = service.UpdatePlanner(newID, planner, &logID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/planner",
		})
		return
	}

	plannerUpdated, err := service.GetPlannerByID(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/planners/:id",
		})
		return
	}
	c.JSON(http.StatusOK, plannerUpdated)

	//// Retorno json com o Planner
	//c.Status(http.StatusNoContent)
}
