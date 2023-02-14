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

	list, err := service.GetReleasesTrain()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains",
		})
		return
	}

	c.JSON(http.StatusOK, list)
}

// Função que chama método GetReleseTrainByID do service e retorna json com lista de users
func GetReleaseTrainByID(c *gin.Context, service service.ReleaseServiceInterface) {

	ID := c.Param("releasetrain_id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains/id/:releasetrain_id",
		})
		return
	}

	// Chama método GetUsers e retorna release
	release, err := service.GetReleaseTrainByID(newId)
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
	c.JSON(http.StatusOK, release)
}

func UpdateReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
	ID := c.Param("releasetrain_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains/update/:releasetrain_id",
		})
		return
	}

	var release *entity.Release_Update

	err = c.ShouldBind(&release)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains/update/:releasetrain_id",
		})
		return
	}

	idResult, err := service.UpdateReleaseTrain(newID, release)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/update/:releasetrain_id",
		})
		return
	}

	_, err = service.InsertTagsReleaseTrain(newID, release.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/update/:releasetrain_id",
		})
		return
	}

	releaseUpdated, err := service.GetReleaseTrainByID(idResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/id/:releasetrain_id",
		})
		return
	}
	c.JSON(http.StatusOK, releaseUpdated)
}

// Função que chama método GetTagsReleaseTrain do service e retorna json com uma lista de tags do client
func GetTagsReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
	ID := c.Param("releasetrain_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains/tag/:releasetrain_id",
		})
		return
	}

	tags, err := service.GetTagsReleaseTrain(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/tag/:releasetrain_id",
		})
		return
	}

	c.JSON(http.StatusOK, tags)
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
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains/update/status/:releasetrain_id",
		})
		return
	}

	// Chama método UpdateStatusUser passando id como parâmetro
	_, err = service.UpdateStatusReleaseTrain(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/update/status/:releasetrain_id",
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains/business/:business_id",
		})
		return
	}

	// Chama método GetReleaseTrainByBusiness e retorna release
	list, err := service.GetReleaseTrainByBusiness(&newId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/business/:business_id",
		})
		return
	}

	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(http.StatusOK, list)
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
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains",
		})
		return
	}

	err = service.CreateReleaseTrain(release)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains",
		})
		return
	}
	// Retorno json com o user
	c.Status(http.StatusNoContent)
}
