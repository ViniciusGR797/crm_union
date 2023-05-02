package controller

import (
	"fmt"
	"microservice_business/pkg/entity"
	"microservice_business/pkg/security"
	"microservice_business/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBusiness função que chama o metodo GetBusiness do service e traz todos os dados de Business do banco em formato de lista
func GetBusiness(c *gin.Context, service service.BusinessServiceInterface) {

	list := service.GetBusiness()

	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}

	c.JSON(200, list)
}

// GetBusinessById função que chama o metodo GetBusinessById do service e traz todos os dados de um Business do banco
func GetBusinessById(c *gin.Context, service service.BusinessServiceInterface) {
	ID := c.Param("id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/business/:id",
		})
		return
	}

	// Chama método GetUsers e retorna business
	business, err := service.GetBusinessById(newId)
	// Verifica se a business está vazia
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business/:id",
		})
		return
	}

	//retorna sucesso 200 e retorna json da lista de business
	c.JSON(http.StatusOK, business)
}

// CreateBusiness interage com o service de CreateBusiness e cria um Business no banco e tem o retorno do mesmo criado
func CreateBusiness(c *gin.Context, service service.BusinessServiceInterface) {
	// Cria variável do tipo business (inicialmente vazia)
	var business *entity.Business_Update

	// Converte json em business
	err := c.ShouldBind(&business)
	// Verifica se tem erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/business",
		})
		return
	}

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

	ctx := c.Request.Context()

	err = service.CreateBusiness(business, &logID, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business",
		})
		return
	}

	// Retorno json com o business
	c.Status(http.StatusNoContent)
}

// UpdateBusiness interage com o service de UpdateBusiness e altera a informações de um Business no banco
func UpdateBusiness(c *gin.Context, service service.BusinessServiceInterface) {
	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/business/update/:id",
		})
		return
	}

	var business *entity.Business_Update

	err = c.ShouldBind(&business)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/business/update/:id",
		})
		return
	}

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

	ctx := c.Request.Context()

	_, err = service.UpdateBusiness(newID, business, &logID, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business/update/:id",
		})
		return
	}

	businessUpdated, err := service.GetBusinessById(newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business/:id",
		})
		return
	}
	c.JSON(http.StatusOK, businessUpdated)

}

// UpdateStatusBusiness interage com o service de UpdateStatusBusiness e altera o status de Business no banco
func UpdateStatusBusiness(c *gin.Context, service service.BusinessServiceInterface) {
	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

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

	ctx := c.Request.Context()

	result := service.UpdateStatusBusiness(&newID, &logID, ctx)
	if result == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"response": "Business Status Updated",
	})
}

// GetBusinessByName interage com o service de GetbusinessByname e traz os dados de business pelo nome pesquisado
func GetBusinessByName(c *gin.Context, service service.BusinessServiceInterface) {
	// Pega name passada como parâmetro na URL da rota
	name := c.Param("Business_name")
	// Chama método GetBusinessByName passando name como parâmetro
	list, err := service.GetBusinessByName(&name)
	// Verifica se teve ao buscar Businesss no banco
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch Businesss",
		})
		return
	}
	// Verifica se a lista de Businesss tem tamanho zero (caso for não tem Business com esse name)
	if len(list.List) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no Businesss found",
		})
		return
	}

	// Retorno json com Business
	c.JSON(http.StatusOK, list)
}

// GetTagsBusiness interage com o service de Business e taz as tags de um business
func GetTagsBusiness(c *gin.Context, service service.BusinessServiceInterface) {
	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/business/tag/:id",
		})
		return
	}

	tags, err := service.GetTagsBusiness(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business/tag/:id",
		})
		return
	}

	c.JSON(http.StatusOK, tags)
}
