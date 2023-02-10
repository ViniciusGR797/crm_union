package service

import (
<<<<<<< HEAD

	// Import interno de packages do próprio sistema
	"fmt"
=======
	"fmt"

	// Import interno de packages do próprio sistema
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
	"microservice_user/pkg/database"
	"microservice_user/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD User (tudo que tiver os métodos abaixo do CRUD são serviços de user)
type UserServiceInterface interface {
<<<<<<< HEAD
	GetUsers() *[]entity.User
=======
	// Pega todos os users, logo lista todos os users
	GetUsers() *entity.UserList
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
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

<<<<<<< HEAD
func (ps *User_service) GetUsers() *[]entity.User {
=======
// Função que retorna lista de users
func (ps *User_service) GetUsers() *entity.UserList {
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
<<<<<<< HEAD
	rows, err := database.Query("SELECT U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id")

=======
	rows, err := database.Query("SELECT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id")
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

<<<<<<< HEAD
	// variável do tipo ProductList (vazia)
	user_list := []entity.User{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Produto (vazia)
		user := entity.User{}

		// pega dados da query e atribui a variável produto, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&user.User_ID, &user.Name, &user.Email, &user.Level, &user.Created_At, &user.Status); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			user_list = append(user_list, user)
=======
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
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
		}

	}

<<<<<<< HEAD
	return &user_list
=======
	// retorna lista de produtos
	return lista_users
>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
}
