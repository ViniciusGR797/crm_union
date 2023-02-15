package service

import (
	// Import interno de packages do próprio sistema
	"errors"
	"fmt"
	"log"
	"microservice_planner/pkg/database"
	"microservice_planner/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD User (tudo que tiver os métodos abaixo do CRUD são serviços de user)
type PlannerServiceInterface interface {
	// Pega todos os planners, logo lista todos os planners
	GetPlannerByID(ID *uint64) (*entity.Planner, error)
	CreatePlanner(planner *entity.PlannerUpdate) error
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

	rowsGuest, err := database.Query("SELECT C.client_name FROM tblClient C INNER JOIN tblEngagementPlannerGuestInvite G ON C.client_id = G.client_id WHERE planner_id = ?", planner.ID)
	if err != nil {
		return &entity.Planner{}, errors.New("error fetching tags from planner by id")
	}

	var guest []entity.Client

	for rowsGuest.Next() {
		client := entity.Client{}

		if err := rowsGuest.Scan(&client.Name); err != nil {
			return &entity.Planner{}, errors.New("error scanning guest from planners by id")
		} else {
			guest = append(guest, client)
		}

	}
	planner.Guest = guest

	return planner, nil
}

// CreateBlanner cria um blanner no banco
func (ps *Planner_service) CreatePlanner(planner *entity.PlannerUpdate) error {

	database := ps.dbp.GetDB()

	var statusID uint64

	rowStatus, err := database.Query("SELECT status_id FROM tblStatus WHERE status_dominio = 'PLANNER' AND status_description = 'SCHEDULED'")
	if err != nil {
		fmt.Println(err.Error())
	}

	rowStatus.Next()

	if err := rowStatus.Scan(&statusID); err != nil {
		fmt.Println(err.Error())
	}

	stmt, err := database.Prepare("INSERT INTO tblPlanner (planner_subject, planner_date, planner_duration, subject_id, client_id, release_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(planner.Name, planner.Date, planner.Duration, planner.Subject, planner.Client, planner.Release, planner.User, statusID)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = result.RowsAffected()
	if err != nil {
		return errors.New("error rowAffected insert into database")
	}

	ID, _ := result.LastInsertId()
	planner.ID = uint64(ID)

	stmt, err = database.Prepare("INSERT INTO tblEngagementPlannerGuestInvite (client_id, planner_id)  VALUES (?, ?)")
	if err != nil {
		return errors.New("error in prepare planner statement")
	}

	for _, guest := range planner.Guest {
		_, err := stmt.Exec(guest.ID, planner.ID)
		if err != nil {
			return errors.New("planner guest error")
		}
	}

	return nil

}
