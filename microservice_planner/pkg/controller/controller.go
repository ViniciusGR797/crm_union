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
	// Cria variável do tipo Planner (inicialmente vazia)
	var planner *entity.PlannerUpdate

	// Converte json em Planner
	err := c.ShouldBind(&planner)
	// Verifica se tem erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/planner",
		})
		return
	}

	err = service.CreatePlanner(planner)
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
	/*name := c.Param("name")

	// Chama método GetPlannerByName passando name como parâmetro
	list, err := service.GetPlannerByName(&name)

	// Verifica se teve ao buscar Planner no banco
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch Planner",
		})
		return
	}
	// Verifica se a lista de Planner tem tamanho zero (caso for não tem Planner com esse name)
	if len(list.List) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no Planner found",
		})
		return
	}

	// Retorno json com Planner
	c.JSON(http.StatusOK, list)
	*/

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
