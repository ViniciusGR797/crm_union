package service

import (
	"fmt"
	"log"

	// Import interno de packages do próprio sistema
	"microservice_user/pkg/database"
	"microservice_user/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD User (tudo que tiver os métodos abaixo do CRUD são serviços de user)
type UserServiceInterface interface {
	// Pega todos os users, logo lista todos os users
	GetUsers() *entity.UserList
	// Pega produto em específico passando o id dele como parâmetro
	GetUserByID(ID *int) *entity.User
	// Pega produto em específico passando o name dele como parâmetro
	GetUserByName(name *string) *entity.UserList
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type User_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewUserService(dabase_pool database.DatabaseInterface) *User_service {
	return &User_service{
		dabase_pool,
	}
}

// Função que retorna lista de users
func (ps *User_service) GetUsers() *entity.UserList {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("SELECT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo UserList (vazia)
	lista_users := &entity.UserList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo User (vazia)
		user := entity.User{}

		// pega dados da query e atribui a variável user, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Level, &user.Created_At, &user.Status); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_users.List = append(lista_users.List, &user)
		}

	}

	// retorna lista de users
	return lista_users
}

// Função que retorna user
func (ps *User_service) GetUserByID(ID *int) *entity.User {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("SELECT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE U.user_id = ?")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}
	// fecha linha da query, quando sair da função
	defer stmt.Close()
	// variável do tipo user (vazia)
	user := entity.User{}

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	err = stmt.QueryRow(ID).Scan(&user.ID, &user.Name, &user.Email, &user.Level, &user.Created_At, &user.Status)
	// verifica se teve erro
	if err != nil {
		log.Println("error: cannot find user", err.Error())
	}

	// retorna user
	return &user
}

// Função que retorna user
func (ps *User_service) GetUserByName(name *string) *entity.UserList {
	*name = fmt.Sprintf("%%%s%%", *name)

	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("SELECT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE U.user_name LIKE ? ")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}
	// fecha linha da query, quando sair da função
	defer stmt.Close()

	// variável do tipo UserList (vazia)
	lista_users := &entity.UserList{}

	// manda uma query para ser executada no database
	rows, err := stmt.Query(*name)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo User (vazia)
		user := entity.User{}

		// pega dados da query e atribui a variável user, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Level, &user.Created_At, &user.Status); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_users.List = append(lista_users.List, &user)
		}

	}

	// retorna lista de users
	return lista_users
}
