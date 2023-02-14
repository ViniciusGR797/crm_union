package controller

import (
	"microservice_subject/pkg/entity"
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

func UpdateStatusSubjectFinished(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid subject id",
			"code":    "400",
			"path":    "/subjects/update/finished/:id",
		})
	}

	_, err = service.UpdateStatusSubjectFinished(newid)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Subject status updated successfully",
		"code":    "200",
		"path":    "/subjects/update/finished/:id",
	})

}

func UpdateStatusSubjectCanceled(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid subject id",
			"code":    "400",
			"path":    "/subjects/update/canceled/:id",
		})
	}

	_, err = service.UpdateStatusSubjectCanceled(newid)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Subject status updated successfully",
		"code":    "200",
		"path":    "/subjects/update/canceled/:id",
	})

}

// create a new subject
func CreateSubject(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid subject id",
			"code":    "400",
			"path":    "/subjects/update/canceled/:id",
		})
	}

	var subject entity.CreateSubject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
			"code":    "400",
			"path":    "/subjects/create/user/:id",
		})
	}

	subjectCreated, err := service.CreateSubject(&subject, newid)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(201, subjectCreated)

}
