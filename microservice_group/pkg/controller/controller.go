package controller

import (
	"microservice_group/pkg/entity"
	"microservice_group/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// rota 100% funcional
func GetGroups(c *gin.Context, service service.GroupServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	list, err := service.GetGroups(newid)
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

func GetGroupByID(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	group, err := service.GetGroupByID(newid)

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

// rota 100%
func UpdateStatusGroup(c *gin.Context, service service.GroupServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	result, err := service.UpdateStatusGroup(newid)
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
			"message": "group Active",
		})
		return
	}

	if result == 2 {
		c.JSON(200, gin.H{
			"message": "group Inactive",
		})
		return
	}

}

func GetUsersGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	UserGroup, err := service.GetUsersGroup(newid)

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

func CreateGroup(c *gin.Context, service service.GroupServiceInterface) {

	var group entity.CreateGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	service.CreateGroup(&group)

	c.JSON(201, gin.H{
		"message": "group created",
	})

}

// insert user_list in group
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

	idReturn, err := service.AttachUserGroup(&users, group_id)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	group, err := service.GetGroupByID(uint64(idReturn))
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
	}

	c.JSON(200, group)
}

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

	idReturn, err := service.DetachUserGroup(&users, group_id)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
		return
	}

	group, err := service.GetGroupByID(uint64(idReturn))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, group)
}

func CountUsersGroup(c *gin.Context, service service.GroupServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, 400, c.Request.URL.Path, err)
		return
	}

	CountUser, err := service.CountUsersGroup(newid)
	if err != nil {
		JSONMessenger(c, 500, c.Request.URL.Path, err)
	}

	c.JSON(200, CountUser)

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
