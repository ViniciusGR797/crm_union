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
		c.JSON(400, gin.H{
			"message": "Invalid user id",
			"code":    "400",
			"path":    "/groups/user/:id",
		})
		return
	}

	list, err := service.GetGroups(newid)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"message": "group not found",
			"code":    "404",
			"path":    "/groups/user/:id",
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
			"message": "Invalid group id",
			"code":    "400",
			"path":    "/groups/:id",
		})
		return
	}

	group, err := service.GetGroupByID(newid)

	if group == nil {
		c.JSON(404, gin.H{
			"message": "group not found",
			"code":    "404",
			"path":    "/groups/:id",
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

// rota 100%
func UpdateStatusGroup(c *gin.Context, service service.GroupServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid group id",
			"code":    "400",
			"path":    "/groups/update/status/:id",
		})
		return
	}

	result := service.UpdateStatusGroup(newid)

	if result == 0 {
		c.JSON(404, gin.H{
			"message": "group not found",
			"code":    "404",
			"path":    "/groups/update/status/:id",
		})
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
		c.JSON(400, gin.H{
			"message": "Invalid group id",
			"code":    "400",
			"path":    "groups/usersGroup/:id",
		})
		return
	}

	UserGroup, err := service.GetUsersGroup(newid)

	if UserGroup == nil {

		c.JSON(404, gin.H{
			"message": "group not found",
			"code":    "404",
			"path":    "groups/usersGroup/:id",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return

	}

	if len(UserGroup.List) == 0 {
		c.JSON(200, gin.H{
			"message": "group without users",
		})
		return
	}

	c.JSON(200, UserGroup)
}

func CreateGroup(c *gin.Context, service service.GroupServiceInterface) {

	var group entity.CreateGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid group data",
			"code":    "400",
			"path":    "/groups",
		})
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
		c.JSON(400, gin.H{
			"message": "Invalid group id",
			"code":    "400",
			"path":    "/groups/attach/:id",
		})
		return
	}

	var users entity.GroupIDList

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user list",
			"code":    "400",
			"path":    "/groups/attach/:id",
		})

		return
	}

	idReturn := service.AttachUserGroup(&users, group_id)

	group, err := service.GetGroupByID(uint64(idReturn))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, group)
}

func DetachUserGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	group_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid group id",
			"code":    "400",
			"path":    "/groups/detach/:id",
		})
		return
	}

	var users entity.GroupIDList

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user list",
			"code":    "400",
			"path":    "/groups/detach/:id",
		})

		return
	}

	idReturn := service.DetachUserGroup(&users, group_id)

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
		c.JSON(400, gin.H{
			"message": "Invalid group id",
			"code":    "400",
			"path":    "/groups/count/:id",
		})
		return
	}

	CountUser, err := service.CountUsersGroup(newid)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, CountUser)

}
