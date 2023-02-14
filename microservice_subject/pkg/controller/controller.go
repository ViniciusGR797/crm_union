package controller

import (
	"microservice_subject/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSubjectList(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user id",
			"code":    "400",
			"path":    "/subjects/user/:id",
		})
	}

	list, err := service.GetSubjectList(newid)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, list)

}

func GetSubject(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid subject id",
			"code":    "400",
			"path":    "/subjects/:id",
		})
	}

	subject, err := service.GetSubject(newid)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, subject)

}
