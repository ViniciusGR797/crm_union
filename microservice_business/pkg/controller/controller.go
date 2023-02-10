package controller

import (
	"microservice_business/pkg/service"
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

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	Business, err := service.GetBusinessByID(newid)

	if Business == nil {
		c.JSON(404, gin.H{
			"error": "Business not found, 404",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, Business)
}
