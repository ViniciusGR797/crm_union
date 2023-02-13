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
			"error": err.Error(),
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

// rota 100%
func UpdateStatusGroup(c *gin.Context, service service.GroupServiceInterface) {

	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	service.UpdateStatusGroup(newid)

	c.JSON(200, gin.H{
		"message": "status changed",
	})

}

func GetUsersGroup(c *gin.Context, service service.GroupServiceInterface) {
	id := c.Param("id")

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	UserGroup, err := service.GetUsersGroup(newid)

	if UserGroup == nil {

		c.JSON(404, gin.H{
			"error": "user group not found, 404",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, UserGroup)

}

func CreateGroup(c *gin.Context, service service.GroupServiceInterface) {

	var group entity.CreateGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	service.CreateGroup(&group)

	c.JSON(200, gin.H{
		"message": "group created",
	})

}

// insert user_list in group
func InsertUserGroup(c *gin.Context, service service.GroupServiceInterface) {

	id := c.Param("id")

	group_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users entity.GroupIDList

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	idReturn := service.InsertUserGroup(&users, group_id)

	group, err := service.GetGroupByID(uint64(idReturn))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, group)
}
