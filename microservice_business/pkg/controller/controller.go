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

func GetBusinessByID(c *gin.Context, service service.BusinessServiceInterface) {
	id := c.Param("id")

	newID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/business/id/:id",
		})
		return
	}

	business := service.GetBusinessByID(&newID)
	if business.Business_id == 0 {
		c.JSON(404, gin.H{
			"error": "produto not found, 404",
		})
		return
	}

	c.JSON(http.StatusOK, business)
}

func CreateBusiness(c *gin.Context, service service.BusinessServiceInterface) {

	var business entity.CreateBusiness

	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	service.CreateBusiness(&business)

	c.JSON(200, gin.H{
		"business_code":       business.Busines_code,
		"business_name":       business.Business_name,
		"business_Segment_id": business.Business_Segment_id,
		"status_id":           business.Business_Status_id,
	})

}

func UpdateBusiness(c *gin.Context, service service.BusinessServiceInterface) {
	ID := c.Param("id")

	var business *entity.Business

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer, 400" + err.Error(),
		})
		return
	}

	err = c.ShouldBind(&business)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bin JSON business, 400" + err.Error(),
		})
		return
	}

	idResult := service.UpdateBusiness(&newID, business)
	if idResult == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	business = service.GetBusinessByID(&newID)
	c.JSON(200, business)

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
