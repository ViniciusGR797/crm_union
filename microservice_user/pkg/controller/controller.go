package controller

import (
	"errors"
	"fmt"
	"microservice_user/pkg/entity"
	"microservice_user/pkg/security"
	"microservice_user/pkg/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

// Função que chama método GetUsers do service e retorna json com lista de users
func GetUsers(c *gin.Context, service service.UserServiceInterface) {
	ctx := c.Request.Context()

	// Chama método GetUsers e retorna list de users
	list, err := service.GetUsers(ctx)
	// Verifica se teve ao buscar user no banco
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se a lista está vazia (tem tamanho zero - não tem users no banco)
	if len(list.List) == 0 {
		sendError(c, http.StatusNotFound, errors.New("no users found"))
		return
	}

	//retorna sucesso 200 e retorna json da lista de users
	send(c, http.StatusOK, list)
}

// Função que chama método GetUserByID do service e retorna json com user
func GetUserByID(c *gin.Context, service service.UserServiceInterface) {
	// Pega id passada como parâmetro na URL da rota
	id := c.Param("user_id")

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		sendError(c, http.StatusBadRequest, errors.New("ID must be an integer"))
		return
	}

	ctx := c.Request.Context()

	// Chama método GetUserByID passando id como parâmetro
	user, err := service.GetUserByID(&newId, ctx)
	// Verifica se teve ao buscar user no banco
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se o id é zero (caso for deu erro ao buscar o user no banco)
	if user.ID == 0 {
		sendError(c, http.StatusNotFound, errors.New("user not found"))
		return
	}

	// Retorno json com user
	send(c, http.StatusOK, user)
}

// Função que chama método GetUserByName do service e retorna json com user
func GetUserByName(c *gin.Context, service service.UserServiceInterface) {
	// Pega name passada como parâmetro na URL da rota
	name := c.Param("user_name")
	// Chama método GetUserByName passando name como parâmetro

	ctx := c.Request.Context()

	list, err := service.GetUserByName(&name, ctx)
	// Verifica se teve ao buscar users no banco
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se a lista de users tem tamanho zero (caso for não tem user com esse name)
	if len(list.List) == 0 {
		sendError(c, http.StatusNotFound, errors.New("no user found"))
		return
	}

	// Retorno json com user
	send(c, http.StatusOK, list)
}

// Função que chama método GetUsersNotInGroup do service e retorna json com user
func GetUsersNotInGroup(c *gin.Context, service service.UserServiceInterface) {
	ctx := c.Request.Context()

	// Chama método GetUsersNotInGroup
	list, err := service.GetUsersNotInGroup(ctx)
	// Verifica se teve ao buscar users no banco
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se a lista de users tem tamanho zero (caso for, não tem user sem grupo)
	if len(list.List) == 0 {
		sendError(c, http.StatusNotFound, errors.New("no user found"))
		return
	}

	// Retorno json com user
	send(c, http.StatusOK, list)
}

// Função que chama método GetSubmissiveUsers do service e retorna json com user
func GetSubmissiveUsers(c *gin.Context, service service.UserServiceInterface) {
	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Pega id e nivel passada como token na rota
	id, err := strconv.Atoi(fmt.Sprint(permissions["userID"]))
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	level, err := strconv.Atoi(fmt.Sprint(permissions["level"]))
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	// Verifica se o user é level 1, logo não tem user submissive
	if level <= 1 {
		sendNoContent(c)
		return
	}

	ctx := c.Request.Context()

	// Chama método GetSubmissiveUsers passando id como parâmetro
	list, err := service.GetSubmissiveUsers(&id, level, ctx)
	// Verifica se teve ao buscar users no banco
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se a lista de users tem tamanho zero (caso for user não tem users submissive)
	if len(list.List) == 0 {
		sendNoContent(c)
		return
	}

	// Retorno json com user
	send(c, http.StatusOK, list)
}

// Função que chama método CreateUser do service e retorna json com mensagem de sucesso
func CreateUser(c *gin.Context, service service.UserServiceInterface) {
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

	// Cria variável do tipo user (inicialmente vazia)
	var user *entity.User

	// Converte json em user
	err = c.ShouldBind(&user)
	// Verifica se tem erro
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	// user.Password = security.RandStringRunes(12)

	// Prepara e valida dados
	if err = user.Prepare(); err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	user.Password = security.RandStringRunes(12)

	// Faz hash com a senha
	user.Hash, err = security.HashPassword(user.Password)
	// Verifica se teve erro ao fazer hash
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	ctx := c.Request.Context()

	// Chama método Create passando user como parâmetro, cadastra no banco user
	_, err = service.CreateUser(user, &logID, ctx)
	// Verifica se teve erro na criação de user
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	subject := "Usuario criado no CRM."
	body := "Olá!\n\nFoi criado para você um usuário para o site CRM. \nPara acessar o site entre com seu email corporativo e a seguinte senha: " + user.Password

	sendMail(user.Email, subject, body)

	// Retorno json com o user
	send(c, http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"password": user.Password,
	})
}

// Função que chama método UpdateStatusUser do service e retorna json com mensagem de sucesso
func UpdateStatusUser(c *gin.Context, service service.UserServiceInterface) {
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

	// Pega id passada como parâmetro na URL da rota
	id := c.Param("user_id")

	// Converter ":id" string para int id (newid)
	newID, err := strconv.ParseUint(id, 10, 64)
	// Verifica se teve erro na conversão
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	// Chama método UpdateStatusUser passando id como parâmetro
	result, err := service.UpdateStatusUser(&newID, &logID, ctx)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se o id é zero (caso for deu erro ao editar o user no banco)
	if result == 0 {
		sendError(c, http.StatusNotFound, errors.New("user not found"))
		return
	}

	// Retorno json com mensagem de sucesso
	sendNoContent(c)
}

// Função que chama método UpdateUser do service e retorna json com mensagem de sucesso
func UpdateUser(c *gin.Context, service service.UserServiceInterface) {
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

	// Pega id passada como parâmetro na URL da rota
	id := c.Param("user_id")
	// Cria variável do tipo user (inicialmente vazia)
	var user *entity.User

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	// Converte json em user
	err = c.ShouldBind(&user)
	// Verifica se tem erro
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	// Verifica se senha tem o tamanho mínimo de caracteres
	user.Hash, err = security.HashPassword(user.Password)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	ctx := c.Request.Context()

	// Chama método UpdateUser passando user e id como parâmetro
	err = service.UpdateUser(&newId, user, &logID, ctx)
	// Verifica se teve erro na edição de user
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	// Retorna json com o status 200
	sendNoContent(c)
}

// Função que chama método Login do service e retorna json com token
func Login(c *gin.Context, service service.UserServiceInterface) {
	// Cria variável do tipo user (inicialmente vazia)
	var user *entity.User

	// Converte json em user
	err := c.ShouldBind(&user)
	// Verifica se tem erro
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	// Verifica se email formato válido
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	// Verifica se senha tem o tamanho mínimo de caracteres
	if len(user.Password) < 8 {
		sendError(c, http.StatusBadRequest, errors.New("password too short"))
		return
	}

	ctx := c.Request.Context()

	// Chama método Login passando user como parâmetro
	hash, err := service.Login(user, ctx)
	// Verifica se teve erro ao buscar user no banco
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se a senha com hash está vazia
	if hash == "" {
		sendError(c, http.StatusUnauthorized, errors.New("incorrect credentials"))
		return
	}
	// Verifica se o user é inativo
	if user.Status != "ACTIVE" {
		sendError(c, http.StatusUnauthorized, errors.New("inactive user"))
		return
	}

	// Chama método que compara o hash com a senha, para verificar se são iguais
	err = security.ValidatePassword(hash, user.Password)
	// Caso coloque a senha errada, cai nesse erro
	if err != nil {
		sendError(c, http.StatusUnauthorized, errors.New("incorrect credentials"))
		return
	}

	// Gera token com base no ID do user logado com sucesso
	token, err := security.NewToken(user.ID, user.Level, user.Status)
	// Verifica se teve erro ao gerar o token
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	firstAccess := user.FirstAccess

	// Retorna JSON com o token
	send(c, http.StatusOK, gin.H{
		"token":      token,
		"firstAcess": firstAccess,
	})
}

// Função que chama método GetUserMe do service e retorna json com user
func GetUserMe(c *gin.Context, service service.UserServiceInterface) {
	// pegar informamções do usuário
	permissions, err := security.GetPermissions(c)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	id := fmt.Sprint(permissions["userID"])

	// Converter ":id" string para int id (newid)
	newId, err := strconv.Atoi(strings.Replace(id, ":", "", 1))
	// Verifica se teve erro na conversão
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	// Chama método GetUserByID passando id como parâmetro
	user, err := service.GetUserByID(&newId, ctx)
	// Verifica se teve ao buscar user no banco
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}
	// Verifica se o id é zero (caso for deu erro ao buscar o user no banco)
	if user.ID == 0 {
		sendError(c, http.StatusNotFound, errors.New("user not found"))
		return
	}

	// Retorno json com user
	send(c, http.StatusOK, user)
}

// Função que chama método forgotPwd do service e retorna msg de email enviado com sucesso
func ForgotPwd(c *gin.Context, service service.UserServiceInterface) {
	// Cria variável do tipo user (inicialmente vazia)
	var user *entity.User

	ctx := c.Request.Context()

	// Converte json em user
	err := c.ShouldBind(&user)
	// Verifica se tem erro
	if err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	// Verifica se email formato válido
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	userReturn, err := service.GetUserByEmail(user, ctx)

	if err != nil || userReturn == nil {
		sendError(c, http.StatusBadRequest, err)
		return
	}

	userReturn.Password = security.RandStringRunes(12)

	userReturn.FirstAccess = true

	subject := "Recuperação de senha no CRM."
	body := "Olá!\n\nSua senha de recuperação para o CRM: " + userReturn.Password

	// Faz hash com a senha
	user.Hash, err = security.HashPassword(user.Password)
	// Verifica se teve erro ao fazer hash
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	logID := int(userReturn.ID)
	ID := int(userReturn.ID)

	// Chama método UpdateUser passando user e id como parâmetro
	err = service.UpdateUser(&ID, userReturn, &logID, ctx)
	// Verifica se teve erro na edição de user
	if err != nil {
		sendError(c, http.StatusInternalServerError, err)
		return
	}

	sendMail(user.Email, subject, body)

	// Retorna JSON com o token
	send(c, http.StatusOK, gin.H{
		"Message": "Senha de recuperação enviada para o email do usuário",
	})
}
