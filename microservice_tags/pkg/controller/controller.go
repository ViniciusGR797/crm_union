package controller

import (
	"microservice_tags/pkg/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTags função que chama o metodo GetTags do service e traz todos os dados de Tags do banco em formato de lista
func GetTags(c *gin.Context, service service.TagsServiceInterface) {

	list := service.GetTags()

	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "list not found, 404",
		})
		return
	}

	c.JSON(200, list)
}

// GetTagsById função que chama o metodo GetTagsById do service e traz todos os dados de um Tags do banco
func GetTagsById(c *gin.Context, service service.TagsServiceInterface) {
	ID := c.Param("id")

	newId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
			"path":    "/Tags/:id",
		})
		return
	}

	// Chama método GetUsers e retorna Tags
	Tags, err := service.GetTagsById(newId)
	// Verifica se a Tags está vazia
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"path":    "/Tags/:id",
		})
		return
	}

	//retorna sucesso 200 e retorna json da lista de Tags
	c.JSON(http.StatusOK, Tags)
}
