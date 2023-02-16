package controller

import (
	"microservice_remark/pkg/entity"
	"microservice_remark/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função que chama método GetRemark do service e retorna json com lista de users
// func GetRemark(c *gin.Context, service service.RemarkServiceInterface) {
// 	// Chama método GetRemark e retorna list de users
// 	list := service.GetRemark()
// 	// Verifica se a lista está vazia (tem tamanho zero)
// 	if len(list.List) == 0 {
// 		c.JSON(404, gin.H{
// 			"error": "lista not found, 404",
// 		})
// 		return
// 	}
// 	//retorna sucesso 200 e retorna json da lista de users
// 	c.JSON(200, list)

// }

// Função que chama método GetSubmissiveRemark do service e retorna json com lista
func GetSubmissiveRemarks(c *gin.Context, service service.RemarkServiceInterface) {
	ID := c.Param("user_ID")
	NewID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
	}

	remarks := service.GetSubmissiveRemarks(&NewID)
	if len(remarks.List) == 0 {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)

		return
	}
	c.JSON(200, remarks)

}

// Função que chama método GetRemarkByID do service e retorna json com um client
func GetRemarkByID(c *gin.Context, service service.RemarkServiceInterface) {
	ID := c.Param("remark_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
	}

	remark := service.GetRemarkByID(&newID)
	if remark == nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}
	c.JSON(200, remark)

}

// Função que cria um Remark
func CreateRemark(c *gin.Context, service service.RemarkServiceInterface) {

	var remark *entity.RemarkUpdate

	err := c.ShouldBindJSON(&remark)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	id := service.CreateRemark(remark)
	if id == 0 {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)

	}

	var remarkCreated *entity.Remark
	remarkCreated = service.GetRemarkByID(&id)
	c.JSON(200, remarkCreated)
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

	err = service.UpdateStatusRemark(&newID, remark)
	if err != nil {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	remarkResult := service.GetRemarkByID(&newID)
	c.JSON(200, remarkResult)

}

func UpdateRemark(c *gin.Context, service service.RemarkServiceInterface) {

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

	idResult := service.UpdateRemark(&newid, remark)
	if idResult == 0 {
		JSONMessenger(c, http.StatusBadRequest, c.Request.URL.Path, err)
		return
	}

	remarkResult := service.GetRemarkByID(&newid)
	c.JSON(200, remarkResult)

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
