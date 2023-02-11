package controller

import (
	"microservice_remark/pkg/entity"
	"microservice_remark/pkg/service"
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

func GetSubmissiveRemarks(c *gin.Context, service service.RemarkServiceInterface) {
	ID := c.Param("user_ID")
	NewID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID Has to be interger, 400",
		})
	}

	remarks := service.GetSubmissiveRemarks(&NewID)
	if len(remarks.List) == 0 {
		c.JSON(404, gin.H{
			"error": "ID Remark not found, 404",
		})

		return
	}
	c.JSON(200, remarks)

}
func GetRemarkByID(c *gin.Context, service service.RemarkServiceInterface) {
	ID := c.Param("remark_id")

	newID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID Has to be interger, 400",
		})
	}

	remark := service.GetRemarkByID(&newID)
	if remark == nil {
		c.JSON(404, gin.H{
			"error": "ID Remark not found, 404",
		})

		return
	}
	c.JSON(200, remark)

}

func CreateRemark(c *gin.Context, service service.RemarkServiceInterface) {

	var remark *entity.Remark

	err := c.ShouldBind(&remark)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON remark" + err.Error(),
		})
		return
	}

	id := service.CreateRemark(remark)
	if id == 0 {
		c.JSON(400, gin.H{
			"error": "cannot create JSON: " + err.Error(),
		})

	}

	remark = service.GetRemarkByID(&id)
	c.JSON(200, remark)

}
