package service

import (
	"errors"
	"fmt"
	"log"

	// Import interno de packages do próprio sistema
	"microservice_user/pkg/database"
	"microservice_user/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD User (tudo que tiver os métodos abaixo do CRUD são serviços de user)
type UserServiceInterface interface {
	// Pega todos os users, logo lista todos os users
	GetUsers() (*entity.UserList, error)
	// Pega user em específico passando o id dele como parâmetro
	GetUserByID(ID *int) (*entity.User, error)
	// Pega users em específico passando o name dele como parâmetro
	GetUserByName(name *string) (*entity.UserList, error)
	// Pega users submissos passando o id de um user como parâmetro
	GetSubmissiveUsers(ID *int, level int) (*entity.UserList, error)
	// Cadastra users passando suas informações
	CreateUser(user *entity.User, logID *int) (uint64, error)
	// Altera status do user
	UpdateStatusUser(ID *uint64, logID *int) (int64, error)
	// Atualiza dados de um usuário, passando id do usuário e dados a serem alterados por parâmetro
	UpdateUser(ID *int, user *entity.User, logID *int) (int, error)
	// Busca o hash do usuário por email
	Login(user *entity.User) (string, error)
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
func (ps *User_service) GetUsers() (*entity.UserList, error) {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("SELECT DISTINCT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id ORDER BY U.user_name")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return &entity.UserList{}, errors.New("error fetching users")
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
			log.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_users.List = append(lista_users.List, &user)
		}

	}

	// retorna lista de users
	return lista_users, nil
}

// Função que retorna user
func (ps *User_service) GetUserByID(ID *int) (*entity.User, error) {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("SELECT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE U.user_id = ?")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return &entity.User{}, errors.New("error preparing statement")
	}
	// fecha linha da query, quando sair da função
	defer stmt.Close()
	// variável do tipo user (vazia)
	user := entity.User{}

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	err = stmt.QueryRow(ID).Scan(&user.ID, &user.Name, &user.Email, &user.Level, &user.Created_At, &user.Status)
	// verifica se teve erro
	if err != nil {
		return &entity.User{}, nil
	}

	// retorna user
	return &user, nil
}

// Função que retorna lista de users
func (ps *User_service) GetUserByName(name *string) (*entity.UserList, error) {
	nameString := fmt.Sprint("%", *name, "%")
	query := "SELECT DISTINCT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE U.user_name LIKE ? ORDER BY U.user_name"

	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query(query, nameString)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return &entity.UserList{}, errors.New("error fetching users")
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
			log.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a lista de users
			lista_users.List = append(lista_users.List, &user)
		}

	}

	// retorna lista de users
	return lista_users, nil
}

// Função que retorna lista de users
func (ps *User_service) GetSubmissiveUsers(ID *int, level int) (*entity.UserList, error) {
	query := "SELECT group_id FROM tblUserGroup WHERE user_id = ?"

	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query(query, ID)
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
		return &entity.UserList{}, errors.New("error fetching user's groups")
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
		query := "SELECT DISTINCT U.user_id, U.user_name, U.user_email, U.user_level, U.created_at, S.status_description FROM tblUser U INNER JOIN tblUserGroup UG ON U.user_id = UG.user_id INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE UG.group_id = ? AND U.user_level < ? ORDER BY U.user_level DESC, U.user_name"

		// manda uma query para ser executada no database
		rows, err := database.Query(query, groupID.ID, level)
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
	return lista_users, nil
}

// Função que retorna user
func (ps *User_service) CreateUser(user *entity.User, logID *int) (uint64, error) {
	// pega database
	database := ps.dbp.GetDB()

	// Definir a variável de sessão "@user"
	_, err := database.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	// prepara query para ser executada no database
	stmt, err := database.Prepare("INSERT INTO tblUser (user_name, user_email, user_pwd, user_level, status_id) VALUES (?, ?, ?, ?, ?)")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error preparing statement")
	}
	// fecha linha da query, quando sair da função
	defer stmt.Close()

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	result, err := stmt.Exec(user.Name, user.Email, user.Hash, user.Level, 9) // TODO implement status
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	// pega id do último usuário inserido
	lastId, err := result.LastInsertId()
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error fetching last id")
	}

	// coloca o id do usuário de volta no modelo
	user.ID = uint64(lastId)

	// retorna user
	return uint64(lastId), nil
}

func (ps *User_service) UpdateStatusUser(ID *uint64, logID *int) (int64, error) {
	database := ps.dbp.GetDB()

	// Definir a variável de sessão "@user"
	_, err := database.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	stmt, err := database.Prepare("SELECT status_id FROM tblUser WHERE user_id = ?")
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error preparing statement")
	}

	var statusID uint64

	err = stmt.QueryRow(ID).Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}

	if statusID == 9 {
		statusID = 10
	} else {
		statusID = 9
	}

	updt, err := database.Prepare("UPDATE tblUser SET status_id = ? WHERE user_id = ?")
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error preparing statement")
	}

	defer stmt.Close()

	result, err := updt.Exec(statusID, ID)
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error fetching rows affected")
	}

	return rowsaff, nil
}

// Função que altera o usuário
func (ps *User_service) UpdateUser(ID *int, user *entity.User, logID *int) (int, error) {
	// pega database
	database := ps.dbp.GetDB()

	// Definir a variável de sessão "@user"
	_, err := database.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	// prepara query para ser executada no database
	stmt, err := database.Prepare("UPDATE tblUser SET user_name = ?, user_email = ?, user_pwd = ?, user_level = ? WHERE user_id = ? ")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error preparing statement")
	}
	// fecha linha da query, quando sair da função
	defer stmt.Close()

	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	result, err := stmt.Exec(user.Name, user.Email, user.Hash, user.Level, ID)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}

	// RowsAffected retorna número de linhas afetadas com update
	rowsaff, err := result.RowsAffected()
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error fetching rows affected")
	}

	// retorna rowsaff (converte int64 para int)
	return int(rowsaff), nil
}

func (ps *User_service) Login(user *entity.User) (string, error) {
	// pega database
	database := ps.dbp.GetDB()

	// prepara query para ser executada no database
	stmt, err := database.Prepare("SELECT U.user_id, U.user_pwd, U.user_level, S.status_description FROM tblUser U INNER JOIN tblStatus S ON U.status_id = S.status_id WHERE user_email = ?")
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("error preparing statement")
	}
	// fecha linha da query, quando sair da função
	defer stmt.Close()

	hash := ""
	// substitui ? da query pelos valores passados por parâmetro de Exec, executa a query e retorna um resultado
	err = stmt.QueryRow(user.Email).Scan(&user.ID, &hash, &user.Level, &user.Status)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return "", nil
	}

	return hash, nil
}
