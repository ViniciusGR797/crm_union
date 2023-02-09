package controller

import (
	"microservice_group/pkg/service"
	"strconv"

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

func GetGroupByID(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	group, err := service.GetGroupByID(newid)

	if group == nil {
		c.JSON(404, gin.H{
			"error": "group not found, 404",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, group)
}
