package service

import (
	// Import interno de packages do próprio sistema
	"context"
	"errors"
	"fmt"
	"log"
	"microservice_planner/pkg/database"
	"microservice_planner/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD User (tudo que tiver os métodos abaixo do CRUD são serviços de user)
type PlannerServiceInterface interface {
	// Pega todos os planners, logo lista todos os planners
	GetPlannerByID(ID *uint64, ctx context.Context) (*entity.Planner, error)
	CreatePlanner(planner *entity.CreatePlanner, logID *int, ctx context.Context) error
	GetPlannerByName(ID *int, level int, name *string, ctx context.Context) (*entity.PlannerList, error)
	GetSubmissivePlanners(ID *int, level int, ctx context.Context) (*entity.PlannerList, error)
	GetPlannerByBusiness(name *string, ctx context.Context) (*entity.PlannerList, error)
	GetGuestClientPlanners(ID *uint64, ctx context.Context) ([]*entity.Client, error)
	UpdatePlanner(ID uint64, planner *entity.PlannerUpdate, logID *int, ctx context.Context) (uint64, error)
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
func (ps *Planner_service) GetPlannerByID(ID *uint64, ctx context.Context) (*entity.Planner, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT DISTINCT * FROM vwGetAllPlanners WHERE planner_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	planner := &entity.Planner{}

	err = stmt.QueryRow(ID).Scan(
		&planner.ID, &planner.Name, &planner.Date, &planner.Duration, &planner.Subject_id, &planner.Subject, &planner.Client_id, &planner.Client, &planner.Client_email, &planner.Business_id, &planner.Business, &planner.Release_id, &planner.Release, &planner.Remark_subject, &planner.Remark_text, &planner.User_id, &planner.User, &planner.CreatedBy_id, &planner.CreatedBy_name, &planner.Status)
	if err != nil {
		return &entity.Planner{}, err
	}

	rowsGuest, err := tx.Query("SELECT C.client_id, C.client_name FROM tblClient C INNER JOIN tblEngagementPlannerGuestInvite G ON C.client_id = G.client_id WHERE planner_id = ?", planner.ID)
	if err != nil {
		return &entity.Planner{}, errors.New("error fetching tags from planner by id")
	}
	defer rowsGuest.Close()

	var guest []entity.Client

	for rowsGuest.Next() {
		client := entity.Client{}

		if err := rowsGuest.Scan(&client.ID, &client.Name); err != nil {
			return &entity.Planner{}, errors.New("error scanning guest from planners by id")
		} else {
			guest = append(guest, client)
		}

	}
	planner.Guest = guest

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return planner, nil
}

// CreateBlanner cria um blanner no banco
func (ps *Planner_service) CreatePlanner(planner *entity.CreatePlanner, logID *int, ctx context.Context) error {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rowStatus, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rowStatus.Close()
	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	var statusID uint64

	err = rowStatus.QueryRow("PLANNER", "SCHEDULED").Scan(&statusID)
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO tblPlanner (planner_subject, planner_date, planner_duration, subject_id, remark_id, client_id, release_id, user_id, created_by, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", planner.Name, planner.Date, planner.Duration, planner.Subject, planner.Remark, planner.Client, planner.Release, planner.User, logID, statusID)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	planner.ID = uint64(ID)

	if planner.Guest != nil {

		for _, guest := range planner.Guest {
			_, err := tx.ExecContext(ctx, "INSERT INTO tblEngagementPlannerGuestInvite (client_id, planner_id)  VALUES (?, ?)", guest.ID, planner.ID)
			if err != nil {
				return errors.New("error insert data tag_ID and ID on database")
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (ps *Planner_service) GetPlannerByName(ID *int, level int, name *string, ctx context.Context) (*entity.PlannerList, error) {

	query := "SELECT group_id FROM tblUserGroup WHERE user_id = ?"

	// pega database
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// manda uma query para ser executada no database
	rows, err := tx.QueryContext(ctx, query, ID)
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
		rows, err := tx.QueryContext(ctx, query, groupID.ID, level)
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

	user := entity.User{}
	user.ID = uint64(*ID)
	lista_users.List = append(lista_users.List, &user)

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo PlannerList (vazia)
	lista_planners := &entity.PlannerList{}

	for _, userID := range lista_users.List {

		// query := "SELECT DISTINCT P.planner_id, P.planner_subject, P.planner_date, P.planner_duration, SU.subject_title, C.client_name, B.business_name, R.release_name, U.user_name, P.created_at, S.status_description FROM tblPlanner P INNER JOIN tblSubject SU ON P.subject_id = SU.subject_id INNER JOIN tblClient C ON P.client_id = C.client_id INNER JOIN tblReleaseTrain R ON P.release_id = R.release_id INNER JOIN tblBusiness B ON R.business_id = B.business_id INNER JOIN tblUser U ON P.user_id = U.user_id INNER JOIN tblStatus S ON P.status_id = S.status_id WHERE P.user_id = ? AND P.planner_subject LIKE ? ORDER BY P.planner_subject"
		query = "SELECT DISTINCT vP.* FROM vwGetAllPlanners vP INNER JOIN tblPlanner P ON vP.planner_id = P.planner_id INNER JOIN tblRemark R ON P.remark_id = R.remark_id WHERE P.user_id = ? AND P.planner_subject LIKE ? ORDER BY P.planner_subject"
		nameString := fmt.Sprint("%", *name, "%")
		// manda uma query para ser executada no database
		rows, err := tx.Query(query, userID.ID, nameString)
		// verifica se teve erro
		if err != nil {
			return &entity.PlannerList{}, errors.New("error fetching planners")
		}

		// Pega todo resultado da query linha por linha
		for rows.Next() {
			// variável do tipo User (vazia)
			planner := entity.Planner{}

			// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
			if err := rows.Scan(&planner.ID, &planner.Name, &planner.Date, &planner.Duration, &planner.Subject_id, &planner.Subject, &planner.Client_id, &planner.Client, &planner.Client_email, &planner.Business_id, &planner.Business, &planner.Release_id, &planner.Release, &planner.Remark_subject, &planner.Remark_text, &planner.User_id, &planner.User, &planner.Created_At, &planner.Status); err != nil {
				return &entity.PlannerList{}, errors.New("error scan planners")
			} else {
				// caso não tenha erro, adiciona a lista de users
				lista_planners.List = append(lista_planners.List, &planner)
			}
		}
	}

	for _, planner := range lista_planners.List {
		rowsGuest, err := tx.Query("SELECT C.client_id, C.client_name FROM tblClient C INNER JOIN tblEngagementPlannerGuestInvite G ON C.client_id = G.client_id WHERE planner_id = ?", planner.ID)
		if err != nil {
			return &entity.PlannerList{}, errors.New("error fetching guests")
		}

		var guest []entity.Client

		for rowsGuest.Next() {
			client := entity.Client{}

			if err := rowsGuest.Scan(&client.ID, &client.Name); err != nil {
				return &entity.PlannerList{}, errors.New("error scan guests")
			} else {
				guest = append(guest, client)
			}

		}
		planner.Guest = guest
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// retorna lista de users
	return lista_planners, nil

}

// Função que retorna lista de users
func (ps *Planner_service) GetSubmissivePlanners(ID *int, level int, ctx context.Context) (*entity.PlannerList, error) {
	query := "SELECT group_id FROM tblUserGroup WHERE user_id = ?"

	// pega database
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// manda uma query para ser executada no database
	rows, err := tx.Query(query, ID)
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
		rows, err := tx.Query(query, groupID.ID, level)
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

	user := entity.User{}
	user.ID = uint64(*ID)
	lista_users.List = append(lista_users.List, &user)

	// fecha linha da query, quando sair da função
	defer rows.Close()
	// variável do tipo PlannerList (vazia)
	lista_planners := &entity.PlannerList{}

	for _, userID := range lista_users.List {
		// query := "SELECT DISTINCT P.planner_id, P.planner_subject, P.planner_date, P.planner_duration, SU.subject_title, C.client_name, B.business_name, R.release_name, P.remark_subject, P.remark_text, U.user_name, P.created_at, S.status_description FROM tblPlanner P INNER JOIN tblSubject SU ON P.subject_id = SU.subject_id INNER JOIN tblClient C ON P.client_id = C.client_id INNER JOIN tblReleaseTrain R ON P.release_id = R.release_id INNER JOIN tblBusiness B ON R.business_id = B.business_id INNER JOIN tblUser U ON P.user_id = U.user_id INNER JOIN tblStatus S ON P.status_id = S.status_id WHERE P.user_id = ? ORDER BY P.planner_subject"
		query := "SELECT DISTINCT vP.* FROM vwGetAllPlanners vP INNER JOIN tblPlanner P ON vP.planner_id = P.planner_id LEFT JOIN tblRemark R ON P.remark_id = R.remark_id WHERE P.user_id = ? ORDER BY P.planner_subject"
		// manda uma query para ser executada no database
		rows, err := tx.Query(query, userID.ID)
		// verifica se teve erro
		if err != nil {
			return &entity.PlannerList{}, err
			//  errors.New("error fetching planners")
		}

		// Pega todo resultado da query linha por linha
		for rows.Next() {
			// variável do tipo User (vazia)
			planner := entity.Planner{}

			// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
			if err := rows.Scan(&planner.ID, &planner.Name, &planner.Date, &planner.Duration, &planner.Subject_id, &planner.Subject, &planner.Client_id, &planner.Client, &planner.Client_email, &planner.Client_role, &planner.Domain_value, &planner.Business_id, &planner.Business, &planner.Release_id, &planner.Release, &planner.Remark_subject, &planner.Remark_text, &planner.User_id, &planner.User, &planner.CreatedBy_id, &planner.CreatedBy_name, &planner.Status); err != nil {
				return &entity.PlannerList{}, errors.New("error scan planners")
			} else {
				// caso não tenha erro, adiciona a lista de users
				lista_planners.List = append(lista_planners.List, &planner)
			}
		}
	}

	for _, planner := range lista_planners.List {
		rowsGuest, err := tx.Query("SELECT C.client_name, C.client_id, C.client_email FROM tblClient C INNER JOIN tblEngagementPlannerGuestInvite G ON C.client_id = G.client_id WHERE planner_id = ?", planner.ID)
		if err != nil {
			return &entity.PlannerList{}, errors.New("error fetching guests")
		}

		var guest []entity.Client

		for rowsGuest.Next() {
			client := entity.Client{}

			if err := rowsGuest.Scan(&client.Name, &client.ID, &client.Email); err != nil {
				return &entity.PlannerList{}, errors.New("error scan guests")
			} else {
				guest = append(guest, client)
			}

		}
		planner.Guest = guest
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// retorna lista de users
	return lista_planners, nil
}

func (ps *Planner_service) GetPlannerByBusiness(name *string, ctx context.Context) (*entity.PlannerList, error) {

	nameString := fmt.Sprint("%", *name, "%")

	// Consulta SQL
	query := "SELECT DISTINCT * FROM vwGetAllPlanners WHERE business_name LIKE ? ORDER BY business_name"

	// Atribui o banco de dados
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query(query, nameString)
	if err != nil {
		log.Println(err.Error())
		return &entity.PlannerList{}, errors.New("error fetching Planner")
	}

	defer rows.Close()

	planner_list := &entity.PlannerList{}

	// O método Next() prepara a próxima linha do retorno da consulta para a leitura do método Scan()
	for rows.Next() {

		planner := entity.Planner{}

		// O método Scan() atribui o valor das colunas da linha atual e atribui em ordem  as variáveis informadas
		// no parâmetro. Se ocorrer um erro, este será atribuído a variável 'err'
		if err := rows.Scan(&planner.ID, &planner.Name, &planner.Date, &planner.Duration, &planner.Subject_id, &planner.Subject, &planner.Client_id, &planner.Client, &planner.Client_email, &planner.Business_id, &planner.Business, &planner.Release_id, &planner.Release, &planner.Remark_subject, &planner.Remark_text, &planner.User_id, &planner.User, &planner.CreatedBy_id, &planner.CreatedBy_name, &planner.Status); err != nil {

			return nil, errors.New("error scan planner")

		} else {
			// Adiciona o planner na lista a cada iteração
			planner_list.List = append(planner_list.List, &planner)
		}
	}

	for _, planner := range planner_list.List {
		rowsGuest, err := tx.Query("SELECT C.client_id, C.client_name FROM tblClient C INNER JOIN tblEngagementPlannerGuestInvite G ON C.client_id = G.client_id WHERE planner_id = ?", planner.ID)
		if err != nil {
			return &entity.PlannerList{}, errors.New("error fetching guests")
		}

		var guest []entity.Client

		for rowsGuest.Next() {
			client := entity.Client{}

			if err := rowsGuest.Scan(&client.ID, &client.Name); err != nil {
				return &entity.PlannerList{}, errors.New("error scan guests")
			} else {
				guest = append(guest, client)
			}

		}
		planner.Guest = guest
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return planner_list, nil
}

// GetTagsBusiness busca as tags de business
func (ps *Planner_service) GetGuestClientPlanners(ID *uint64, ctx context.Context) ([]*entity.Client, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("call pcGetClientGuest(?)")
	if err != nil {
		return []*entity.Client{}, errors.New("error fetching on tag business")
	}

	defer stmt.Close()

	var guests []*entity.Client

	rowsGuests, err := stmt.Query(ID)
	if err != nil {
		return []*entity.Client{}, errors.New("error fetching on Guests clients")
	}

	for rowsGuests.Next() {
		client := entity.Client{}

		if err := rowsGuests.Scan(&client.ID, &client.Name, &client.Email); err != nil {
			return []*entity.Client{}, errors.New("error fetching on row tags next release train")
		}

		guests = append(guests, &client)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return guests, nil
}

func (ps *Planner_service) UpdatePlanner(ID uint64, planner *entity.PlannerUpdate, logID *int, ctx context.Context) (uint64, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	result, err := tx.ExecContext(ctx, "UPDATE tblPlanner SET planner_subject = ?, planner_date = ?, planner_duration = ?, subject_id = ?, client_id = ?, release_id = ?, remark_id = ?, user_id = ?, status_id = ? WHERE planner_id = ?", planner.Name, planner.Date, planner.Duration, planner.Subject, planner.Client, planner.Release, planner.Remark, planner.User, planner.Status, ID)
	if err != nil {
		return 0, err
	}

	plannerID, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("error rowAffected update into database")
	}

	planner.ID = uint64(ID)

	if planner.Guest != nil {
		_, err = tx.ExecContext(ctx, "DELETE FROM tblEngagementPlannerGuestInvite WHERE planner_id = ?", ID)
		if err != nil {
			return 0, errors.New("error prepare delete guests on planner")
		}

		for _, guest := range planner.Guest {
			_, err := tx.ExecContext(ctx, "INSERT IGNORE tblEngagementPlannerGuestInvite SET client_id = ?, planner_id = ?", guest.ID, planner.ID)
			if err != nil {
				return 0, errors.New("error insert data tag_ID and ID on database")
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint64(plannerID), nil
}
