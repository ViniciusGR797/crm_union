package controller

import (
	"microservice_release/pkg/service"

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
