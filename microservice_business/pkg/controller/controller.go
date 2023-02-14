package controller

import (
	"microservice_business/pkg/entity"
	"microservice_business/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	// Chama método GetUsers e retorna release
	business, err := service.GetBusinessById(newId)
	// Verifica se a release está vazia
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/id/:releasetrain_id",
		})
		return
	}
	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(http.StatusOK, business)
}

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

	err = service.CreateBusiness(business)
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

	idResult, err := service.UpdateBusiness(newID, business)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business/update/:id",
		})
		return
	}

	_, err = service.InsertTagsBusiness(newID, business.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/business/update/:id",
		})
		return
	}

	businessUpdated, err := service.GetBusinessById(idResult)
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

func SoftDeleteBusiness(c *gin.Context, service service.BusinessServiceInterface) {
	ID := c.Param("id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	result := service.SoftDeleteBusiness(&newID)
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
