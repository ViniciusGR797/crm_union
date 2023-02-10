package controller

import (
	"microservice_business/pkg/entity"
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
