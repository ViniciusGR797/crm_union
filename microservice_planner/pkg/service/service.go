package service

import (
	// Import interno de packages do próprio sistema
	"errors"
	"log"
	"microservice_planner/pkg/database"
	"microservice_planner/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD User (tudo que tiver os métodos abaixo do CRUD são serviços de user)
type PlannerServiceInterface interface {
	// Pega todos os planners, logo lista todos os planners
	GetPlannerByID(ID *uint64) (*entity.Planner, error)
	GetSubmissivePlanners(ID *int, level int) (*entity.PlannerList, error)
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type Planner_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewPlannerService(dabase_pool database.DatabaseInterface) *Planner_service {
	return &Planner_service{
		dabase_pool,
	}
}

// Função que retorna lista de planners
func (ps *Planner_service) GetPlannerByID(ID *uint64) (*entity.Planner, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT * FROM vwGetAllPlanners WHERE planner_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	planner := &entity.Planner{}

	err = stmt.QueryRow(ID).Scan(&planner.ID, &planner.Name, &planner.Date, &planner.Duration, &planner.Subject, &planner.Client, &planner.Release, &planner.User, &planner.Created_At, &planner.Status)
	if err != nil {
		return &entity.Planner{}, errors.New("error scanning rows")
	}

	return planner, nil
}

// Função que retorna lista de users
func (ps *Planner_service) GetSubmissivePlanners(ID *int, level int) (*entity.PlannerList, error) {
	query := "SELECT group_id FROM tblUserGroup WHERE user_id = ?"

	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query(query, ID)
	// verifica se teve erro
	if err != nil {
		return &entity.PlannerList{}, errors.New("error fetching user's groups")
	}

	// variável do tipo UserList (vazia)
	groupIDList := &entity.GroupIDList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo User (vazia)
		groupID := entity.GroupID{}

		// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&groupID.ID); err != nil {
			return &entity.PlannerList{}, errors.New("error scan user's groups")
		} else {
			// caso não tenha erro, adiciona a lista de users
			groupIDList.List = append(groupIDList.List, &groupID)
		}
	}

	// variável do tipo UserList (vazia)
	lista_users := &entity.UserList{}

	for _, groupID := range groupIDList.List {
		query := "SELECT DISTINCT U.user_id FROM tblUser U INNER JOIN tblUserGroup UG ON U.user_id = UG.user_id WHERE UG.group_id = ? AND U.user_level < ?"

		// manda uma query para ser executada no database
		rows, err := database.Query(query, groupID.ID, level)
		// verifica se teve erro
		if err != nil {
			return &entity.PlannerList{}, errors.New("error fetching users")
		}

		// Pega todo resultado da query linha por linha
		for rows.Next() {
			// variável do tipo User (vazia)
			user := entity.User{}

			// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
			if err := rows.Scan(&user.ID); err != nil {
				return &entity.PlannerList{}, errors.New("error scan users")
			} else {
				// caso não tenha erro, adiciona a lista de users
				lista_users.List = append(lista_users.List, &user)
			}
		}
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo PlannerList (vazia)
	lista_planners := &entity.PlannerList{}

	for _, userID := range lista_users.List {
		query := "SELECT P.planner_id, P.planner_subject, P.planner_date, P.planner_duration, SU.subject_title, C.client_name, R.release_name, U.user_name, P.created_at, S.status_description FROM tblPlanner P INNER JOIN tblSubject SU ON P.subject_id = SU.subject_id INNER JOIN tblClient C ON P.client_id = C.client_id INNER JOIN tblReleaseTrain R ON P.release_id = R.release_id INNER JOIN tblUser U ON P.user_id = U.user_id INNER JOIN tblStatus S ON P.status_id = S.status_id WHERE P.user_id = ?"

		// manda uma query para ser executada no database
		rows, err := database.Query(query, userID.ID)
		// verifica se teve erro
		if err != nil {
			return &entity.PlannerList{}, errors.New("error fetching planners")
		}

		// Pega todo resultado da query linha por linha
		for rows.Next() {
			// variável do tipo User (vazia)
			planner := entity.Planner{}

			// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
			if err := rows.Scan(&planner.ID, &planner.Name, &planner.Date, &planner.Duration, &planner.Subject, &planner.Client, &planner.Release, &planner.User, &planner.Created_At, &planner.Status); err != nil {
				return &entity.PlannerList{}, errors.New("error scan planners")
			} else {
				// caso não tenha erro, adiciona a lista de users
				lista_planners.List = append(lista_planners.List, &planner)
			}
		}
	}

	// retorna lista de users
	return lista_planners, nil
}
