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
	// Pega user em específico passando o id dele como parâmetro
	GetUserByID(ID *int) *entity.User
	// Pega users em específico passando o name dele como parâmetro
	GetUserByName(name *string) *entity.UserList
	// Pega users submissos passando o id de um user como parâmetro
	GetSubmissiveUsers(ID *int) *entity.UserList
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

// Função que retorna lista de users
func (ps *User_service) GetUserByName(name *string) *entity.UserList {
	nameString := fmt.Sprint("%", *name, "%")
	query := fmt.Sprint("SELECT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE U.user_name LIKE ?")

	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query(query, nameString)
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
			// caso não tenha erro, adiciona a lista de users
			lista_users.List = append(lista_users.List, &user)
		}

	}

	// retorna lista de users
	return lista_users
}

// Função que retorna lista de users
func (ps *User_service) GetSubmissiveUsers(ID *int) *entity.UserList {
	query := fmt.Sprint("SELECT group_id FROM tblUserGroup WHERE user_id = ?")

	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query(query, ID)
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// variável do tipo UserList (vazia)
	groupIDList := &entity.GroupIDList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo User (vazia)
		groupID := entity.GroupID{}

		// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&groupID.ID); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a lista de users
			groupIDList.List = append(groupIDList.List, &groupID)
		}
	}

	// variável do tipo UserList (vazia)
	lista_users := &entity.UserList{}

	for _, groupID := range groupIDList.List {
		query := fmt.Sprint("SELECT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblUserGroup UG ON U.user_id = UG.user_id INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE UG.group_id = ? AND U.user_level < (SELECT user_level FROM tblUser WHERE user_id = ?)")

		// manda uma query para ser executada no database
		rows, err := database.Query(query, groupID.ID, ID)
		// verifica se teve erro
		if err != nil {
			fmt.Println(err.Error())
		}

		// Pega todo resultado da query linha por linha
		for rows.Next() {
			// variável do tipo User (vazia)
			user := entity.User{}

			// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
			if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Level, &user.Created_At, &user.Status); err != nil {
				fmt.Println(err.Error())
			} else {
				// caso não tenha erro, adiciona a lista de users
				lista_users.List = append(lista_users.List, &user)
			}
		}
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// retorna lista de users
	return lista_users
}

// Função que retorna user
func (ps *User_service) CreateUser(user *entity.User) uint64 {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("INSERT INTO tblUser (user_name, user_email, user_pwd, status_id) VALUES (?, ?, ?, ?);")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}
	// fecha linha da query, quando sair da função
	defer stmt.Close()

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	result, err := database.Exec(user.Name, user.Email, user.Password, 9) // TODO implement status
	if err != nil {
		log.Println(err.Error())
	}

	// pega id do último product inserido
	lastId, err := result.LastInsertId()
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
	}

	// coloca o id do usuário de volta no modelo
	user.ID = uint64(lastId)

	// retorna user
	return uint64(lastId)
}
