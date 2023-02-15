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
