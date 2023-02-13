package controller

import (
	"microservice_release/pkg/entity"
	"microservice_release/pkg/service"
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

	_, err = service.InsertTagsReleaseTrain(newID,release.Tags)
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
/*
func InsertTagsReleaseTrain(c *gin.Context, service service.ReleaseServiceInterface) {
	ID := c.Param("releasetrain_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	var release *entity.Release_Insert_Tag

	
	err = c.ShouldBindJSON(&release)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON release, 400" + err.Error(),
		})
		return
	}
	
	release.ID = newID
	
	_, err = service.InsertTagsReleaseTrain(&newID, release)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}


	c.JSON(200, gin.H{
		"result": "tags inserted successfully",
	})

}
*/
