package controller

import (
	"microservice_user/pkg/entity"
	"microservice_user/pkg/security"
	"microservice_user/pkg/service"
	"strconv"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

// Função que chama método GetUsers do service e retorna json com lista de users
func GetUsers(c *gin.Context, service service.UserServiceInterface) {
	// Chama método GetUsers e retorna list de users
	list := service.GetUsers()
	// Verifica se a lista está vazia (tem tamanho zero)
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}
	//retorna sucesso 200 e retorna json da lista de users
	c.JSON(200, list)
}

// Função que chama método GetUserByID do service e retorna json com user
func GetUserByID(c *gin.Context, service service.UserServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	id := c.Param("user_id")

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}
	// Chama método GetUserByID passando id como parâmetro
	user := service.GetUserByID(&newId)
	if user.ID == 0 {
		c.JSON(404, gin.H{
			"error": "user not found, 404",
		})
		return
	}

	// Retorno json com user
	c.JSON(200, user)
}

// Função que chama método GetUserByName do service e retorna json com user
func GetUserByName(c *gin.Context, service service.UserServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	name := c.Param("user_name")
	// Chama método GetUserByName passando id como parâmetro
	list := service.GetUserByName(&name)
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}

	// Retorno json com user
	c.JSON(200, list)
}

// Função que chama método GetSubmissiveUsers do service e retorna json com user
func GetSubmissiveUsers(c *gin.Context, service service.UserServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	id := c.Param("user_id")

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400",
		})
		return
	}
	// Chama método GetUserByName passando id como parâmetro
	list := service.GetSubmissiveUsers(&newId)
	if len(list.List) == 0 {
		c.JSON(404, gin.H{
			"error": "lista not found, 404",
		})
		return
	}

	// Retorno json com user
	c.JSON(200, list)
}

func CreateUser(c *gin.Context, service service.UserServiceInterface) {
	// Cria variável do tipo usuario (inicialmente vazia)
	var user *entity.User

	// Converte json em usuario
	err := c.ShouldBind(&user)
	// Verifica se tem erro
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON user" + err.Error(),
		})
		return
	}

	// valida email
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email",
		})
		return
	}

	// validação de senha
	if len(user.Password) < 8 {
		c.JSON(400, gin.H{
			"error": "password too short",
		})
		return
	}

	// hash da senha
	user.Password, err = security.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not hash password: " + err.Error(),
		})
		return
	}

	// Chama método Create passando produto como parâmetro que retorna id novo
	_, err = service.CreateUser(user)
	// Verifica se o id é zero (caso for deu erro ao criar produto no banco)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot create user",
		})
		return
	}

	// Retorno json com o produto
	c.JSON(200, gin.H{
		"message": "user registered successfully",
	})
}

func UpdateStatusUser(c *gin.Context, service service.UserServiceInterface) {
	ID := c.Param("user_id")

	newID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger, 400" + err.Error(),
		})
		return
	}

	result := service.UpdateStatusUser(&newID)
	if result == 0 {
		c.JSON(400, gin.H{
			"error": "cannot update JSON, 400" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"response": "Client Status Updated",
	})
}
