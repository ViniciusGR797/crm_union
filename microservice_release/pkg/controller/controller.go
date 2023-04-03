package controller

import (
	"fmt"
	"microservice_release/pkg/entity"
	"microservice_release/pkg/security"
	"microservice_release/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetReleasesTrain Função que chama método GetReleasesTrain do service e retorna json com lista de release
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

// GetReleaseTrainByID Função que chama método GetReleseTrainByID do service e retorna json
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

// UpdateReleaseTrain Função que chama método UpdateReleaseTrain do service e retorna json
func UpdateReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
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

	idResult, err := service.UpdateReleaseTrain(newID, release, &logID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/update/:releasetrain_id",
		})
		return
	}

	_, err = service.InsertTagsReleaseTrain(newID, release.Tags, &logID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/update/:releasetrain_id",
		})
		return
	}

	_, err = service.GetReleaseTrainByID(idResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/id/:releasetrain_id",
		})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetTagsReleaseTrain Função que chama método GetTagsReleaseTrain do service e retorna json
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

// UpdateStatusReleaseTrain Função que chama método UpdateStatusReleaseTrain do service e retorna json
func UpdateStatusReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
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
	_, err = service.UpdateStatusReleaseTrain(&newID, &logID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/releasetrains/update/status/:releasetrain_id",
		})
		return
	}

	// Retorno json com mensagem de sucesso
	c.JSON(http.StatusNoContent, nil)
}

// GetReleaseTrainByBusiness Função que chama método GetReleaseTrainByBusiness do service e retorna json com release
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

// CreateReleaseTrain Função que chama método CreateReleaseTrain do service e retorna json com mensagem de sucesso
func CreateReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
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

	// Cria variável do tipo release (inicialmente vazia)
	var release *entity.Release_Update

	// Converte json em release
	err = c.ShouldBind(&release)
	// Verifica se tem erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/releasetrains",
		})
		return
	}

	err = service.CreateReleaseTrain(release, &logID)
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
