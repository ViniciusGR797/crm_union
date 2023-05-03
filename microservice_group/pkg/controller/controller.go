package controller

import (
	"fmt"
	"microservice_group/pkg/entity"
	"microservice_group/pkg/security"
	"microservice_group/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetGroups lista todos os grupos do usuario
func GetGroups(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	ctx := c.Request.Context()

	list, err := service.GetGroups(newid, ctx)
	if err != nil {
		JSONMessenger(c, 404, c.Request.URL.Path, err)
		return
	}

	if len(list.List) == 0 {
		JSONMessenger(c, 404, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, list)
}

// GetGroupByID lista um grupo pelo id
func GetGroupByID(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	ctx := c.Request.Context()
	group, err := service.GetGroupByID(newid, ctx)

	if group == nil {
		JSONMessenger(c, 404, c.Request.URL.Path, err)
		return
	}

	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	c.JSON(200, group)
}

// UpdateStatusGroup altera o status do grupo para ativo ou inativo
func UpdateStatusGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

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

	ctx := c.Request.Context()

	result, err := service.UpdateStatusGroup(newid, &logID, ctx)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	if result == 0 {
		JSONMessenger(c, 404, c.Request.URL.Path, err)
		return
	}

	if result == 1 {
		c.JSON(200, gin.H{
			"message": "successfully changed status",
		})
		return
	}

}

// GetUsersGroup lista todos os usuarios do grupo
func GetUsersGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	ctx := c.Request.Context()

	UserGroup, err := service.GetUsersGroup(newid, ctx)

	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return

	}
	if len(UserGroup.List) == 0 {
		c.JSON(404, gin.H{
			"message": "group without users",
		})
		return

	}

	c.JSON(200, UserGroup)
}

// CreateGroup cria um novo grupo
func CreateGroup(c *gin.Context, service service.GroupServiceInterface) {
	var group entity.CreateGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

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

	ctx := c.Request.Context()

	service.CreateGroup(&group, &logID, ctx)

	c.JSON(201, gin.H{
		"message": "group created",
	})

}

// AttachUserGroup adiciona um ou varios  usuario ao grupo
func AttachUserGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	group_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	var users entity.GroupIDList

	if err := c.ShouldBindJSON(&users); err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)

		return
	}

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

	ctx := c.Request.Context()

	idReturn, err := service.AttachUserGroup(&users, group_id, &logID, ctx)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	group, err := service.GetGroupByID(uint64(idReturn), ctx)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
	}

	c.JSON(200, group)
}

// DetachUserGroup remove um ou varios  usuario do grupo
func DetachUserGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	group_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	var users entity.GroupIDList

	if err := c.ShouldBindJSON(&users); err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)

		return
	}

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

	ctx := c.Request.Context()

	idReturn, err := service.DetachUserGroup(&users, group_id, &logID, ctx)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	group, err := service.GetGroupByID(uint64(idReturn), ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, group)
}

// CountUsersGroup conta a quantidade de usuarios do grupo
func CountUsersGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	ctx := c.Request.Context()
	CountUser, err := service.CountUsersGroup(newid, ctx)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
	}

	c.JSON(200, CountUser)

}

func EditGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	var group entity.EditGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

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

	ctx := c.Request.Context()

	idReturn, err := service.EditGroup(&group, newid, &logID, ctx)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	groupReturn, err := service.GetGroupByID(uint64(idReturn), ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, groupReturn)
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
