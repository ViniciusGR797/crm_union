package controller

import (
	"microservice_group/pkg/service"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context, service service.GroupServiceInterface) {

	list := service.GetGroups()

	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}

	c.JSON(200, list)
}
