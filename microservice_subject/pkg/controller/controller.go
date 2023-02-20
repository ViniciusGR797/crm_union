package controller

import (
	"microservice_subject/pkg/entity"
	"microservice_subject/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSubjectList retorna uma lista de Subjects de um determinado usuario
func GetSubjectList(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	list, err := service.GetSubjectList(newid)

	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, list)

}

// GetSubject retorna um Subject pelo id
func GetSubject(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	subject, err := service.GetSubject(newid)

	if err != nil {
		JSONMessenger(c, 404, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, subject)

}

// UpdateStatusSubjectFinished atualiza o status de um Subject para "finished" pelo id
func UpdateStatusSubjectFinished(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	_, err = service.UpdateStatusSubjectFinished(newid)

	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Subject status updated successfully",
		"code":    "200",
		"path":    "/subjects/update/finished/:id",
	})

}

// UpdateStatusSubjectCanceled atualiza o status de um Subject para "canceled" pelo id
func UpdateStatusSubjectCanceled(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	_, err = service.UpdateStatusSubjectCanceled(newid)

	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Subject status updated successfully",
		"code":    "200",
		"path":    "/subjects/update/canceled/:id",
	})

}

// CreateSubject cria um novo Subject
func CreateSubject(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	var subject entity.CreateSubject

	if err := c.ShouldBindJSON(&subject); err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	subjectCreated, err := service.CreateSubject(&subject, newid)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(201, subjectCreated)

}

// UpdateSubject atualiza um Subject
func UpdateSubject(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	var subject entity.UpdateSubject

	if err := c.ShouldBindJSON(&subject); err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	_, err = service.UpdateSubject(newid, &subject)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Subject updated successfully",
		"code":    "200",
		"path":    "/subjects/update/:id",
	})

}

func JSONMessenger(c *gin.Context, status int, path string, err error) {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	c.JSON(status, gin.H{
		"status":  status,
		"message": errorMessage,
		"error":   err,
		"path":    path,
	})
}
