package controller

import (
	"fmt"
	"microservice_subject/pkg/entity"
	"microservice_subject/pkg/security"
	"microservice_subject/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSubmissiveSubjects retorna uma lista de Subjects de um determinado usuario
func GetSubmissiveSubjects(c *gin.Context, service service.SubjectServiceInterface) {
	// Pega permissões do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id passada como token na rota
	id, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	list, err := service.GetSubmissiveSubjects(id)

	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, list)

}

// GetSubjectByID retorna um Subject pelo id
func GetSubjectByID(c *gin.Context, service service.SubjectServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	subject, err := service.GetSubjectByID(newid)

	if err != nil {
		JSONMessenger(c, 404, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, subject)

}

// UpdateStatusSubjectFinished atualiza o status de um Subject para "finished" pelo id
func UpdateStatusSubjectFinished(c *gin.Context, service service.SubjectServiceInterface) {

	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	ctx := c.Request.Context()

	_, err = service.UpdateStatusSubjectFinished(newid, &logID, ctx)

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

	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	ctx := c.Request.Context()

	_, err = service.UpdateStatusSubjectCanceled(newid, &logID, ctx)

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

	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	ctx := c.Request.Context()

	subjectCreated, err := service.CreateSubject(&subject, newid, &logID, ctx)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(201, subjectCreated)

}

// UpdateSubject atualiza um Subject
func UpdateSubject(c *gin.Context, service service.SubjectServiceInterface) {

	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Pega id e nivel passada como token na rota
	logID, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	ctx := c.Request.Context()

	_, err = service.UpdateSubject(newid, &subject, &logID, ctx)
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
