package controller

import (
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
	release := service.GetReleaseTrainByID(&newId)
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



