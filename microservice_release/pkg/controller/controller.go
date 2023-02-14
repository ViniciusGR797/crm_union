package controller

import (
	"microservice_release/pkg/entity"
	"microservice_release/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetReleases do service e retorna json com lista de release
func GetReleasesTrain(c *gin.Context, service service.ReleaseServiceInterface) {

	list := service.GetReleasesTrain()
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}

	c.JSON(200, list)
}

// Função que chama método GetReleseTrainByID do service e retorna json com lista de users
func GetReleaseTrainByID(c *gin.Context, service service.ReleaseServiceInterface) {

	ID := c.Param("releasetrain_id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	// Chama método GetUsers e retorna release
	release := service.GetReleaseTrainByID(newId)
	// Verifica se a release está vazia
	if release == nil {
		c.JSON(404, gin.H{
			"error": "release not found, 404",
		})
		return
	}
	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(200, release)
}

func UpdateReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
	ID := c.Param("releasetrain_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	var release *entity.Release_Update

	err = c.ShouldBind(&release)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON produto, 400" + err.Error(),
		})
		return
	}

	idResult := service.UpdateReleaseTrain(newID, release)
	if idResult == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON" + err.Error(),
		})
		return
	}

	_, err = service.InsertTagsReleaseTrain(newID, release.Tags)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot update tags" + err.Error(),
		})
		return
	}

	releaseUpdated := service.GetReleaseTrainByID(idResult)
	c.JSON(200, releaseUpdated)
}

// Função que chama método GetTagsReleaseTrain do service e retorna json com uma lista de tags do client
func GetTagsReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
	ID := c.Param("releasetrain_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	tags := service.GetTagsReleaseTrain(&newID)
	if len(tags) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}

	c.JSON(200, tags)
}

// Função que chama método UpdateStatusReleaseTrain do service e retorna json com mensagem de sucesso
func UpdateStatusReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	id := c.Param("releasetrain_id")

	// Converter ":id" string para int id (newid)
	newID, err := strconv.ParseUint(id, 10, 64)
	// Verifica se teve erro na conversão
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	// Chama método UpdateStatusUser passando id como parâmetro
	result, err := service.UpdateStatusReleaseTrain(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot update JSON",
		})
		return
	}
	// Verifica se o id é zero (caso for deu erro ao editar o user no banco)
	if result == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	// Retorno json com mensagem de sucesso
	c.JSON(http.StatusOK, gin.H{
		"response": "Release Train Status Updated",
	})
}

// Função que chama método GetReleaseTrainByBusiness do service e retorna json com release
func GetReleaseTrainByBusiness(c *gin.Context, service service.ReleaseServiceInterface) {
	ID := c.Param("business_id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}

	// Chama método GetReleaseTrainByBusiness e retorna release
	list, err := service.GetReleaseTrainByBusiness(&newId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
		return
	}

	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(200, list)
}

// Função que chama método CreateReleaseTrain do service e retorna json com mensagem de sucesso
func CreateReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
	// Cria variável do tipo release (inicialmente vazia)
	var release *entity.Release_Update

	// Converte json em release
	err := c.ShouldBind(&release)
	// Verifica se tem erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON user" + err.Error(),
		})
		return
	}

	err = service.CreateReleaseTrain(release)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	/*_, err = service.InsertTagsReleaseTrain(newID, release.Tags)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot update tags" + err.Error(),
		})
		return
	}

	releaseUpdated := service.GetReleaseTrainByID(idResult)*/

	// Retorno json com o user
	c.Status(http.StatusNoContent)
}
