package controller

import (
	"fmt"
	"microservice_remark/pkg/entity"
	"microservice_remark/pkg/security"
	"microservice_remark/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetSubmissiveRemark do service e retorna json com lista
func GetSubmissiveRemarks(c *gin.Context, service service.RemarkServiceInterface) {
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

	/*ID := c.Param("user_ID")
	NewID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}*/

	remarks, err := service.GetSubmissiveRemarks(&id)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusOK, remarks)

}

// Função que chama método GetRemarkByID do service e retorna json com um client
func GetRemarkByID(c *gin.Context, service service.RemarkServiceInterface) {
	ID := c.Param("remark_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	remark, err := service.GetRemarkByID(&newID)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusOK, remark)

}

// Função que cria um Remark
func CreateRemark(c *gin.Context, service service.RemarkServiceInterface) {

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

	var remark *entity.RemarkUpdate

	err = c.ShouldBindJSON(&remark)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = service.CreateRemark(remark, &logID)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"result": "remark created successfully",
	})
}

func GetBarChartRemark(c *gin.Context, service service.RemarkServiceInterface) {
	ID := c.Param("user_ID")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
	}

	remark := service.GetBarChartRemark(&newID)
	if remark == nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}
	c.JSON(200, remark)
}

func GetPieChartRemark(c *gin.Context, service service.RemarkServiceInterface) {
	ID := c.Param("user_ID")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
	}

	remark := service.GetPieChartRemark(&newID)
	if remark == nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}
	c.JSON(200, remark)
}

// Função que chama método UpdateStatusRemark do service e realiza o softdelete
func UpdateStatusRemark(c *gin.Context, service service.RemarkServiceInterface) {
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

	ID := c.Param("remark_id")

	var remark *entity.Remark

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = c.ShouldBind(&remark)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = service.UpdateStatusRemark(&newID, remark, &logID)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"result": "remark_status updated successfully",
	})

}

func UpdateRemark(c *gin.Context, service service.RemarkServiceInterface) {
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

	id := c.Param("remark_id")

	var remark *entity.RemarkUpdate

	newid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = c.ShouldBind(&remark)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	err = service.UpdateRemark(&newid, remark, &logID)
	if err != nil {
		JSONMessenger(c, http.StatusInternalServerError, c.Request.URL.Path, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"result": "remark updated successfully",
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
