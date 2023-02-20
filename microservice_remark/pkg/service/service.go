package service

import (
	"errors"
	"log"

	// Import interno de packages do próprio sistema

	"microservice_remark/pkg/database"
	"microservice_remark/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Remark (tudo que tiver os métodos abaixo do CRUD são serviços de Remark)
type RemarkServiceInterface interface {
	// Pega todos os Remarks, logo lista todos os Remarks
	GetSubmissiveRemarks(ID *int) (*entity.RemarkList, error)
	GetRemarkByID(ID *uint64) (*entity.Remark, error)
	CreateRemark(remark *entity.RemarkUpdate) error
	GetBarChartRemark(ID *uint64) *entity.Remark
	GetPieChartRemark(ID *uint64) *entity.Remark
	UpdateStatusRemark(ID *uint64, remark *entity.Remark) error
	UpdateRemark(ID *uint64, remark *entity.RemarkUpdate) error
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type remark_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewRemarkService(dabase_pool database.DatabaseInterface) *remark_service {
	return &remark_service{
		dabase_pool,
	}
}

// Função que retorna lista de Remarks
func (ps *remark_service) GetSubmissiveRemarks(ID *int) (*entity.RemarkList, error) {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("call pcGetAllRemarkUserGroup (?)", ID)
	// verifica se teve erro
	if err != nil {
		return nil, err
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo RemarkList (vazia)
	lista_Remarks := &entity.RemarkList{}

	hasResult := false

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		hasResult = true
		// variável do tipo Remark (vazia)
		remark := entity.Remark{}

		// pega dados da query e atribui a variável Remark, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&remark.ID, &remark.Remark_Name, &remark.User_Name, &remark.Subject_Name, &remark.Client_Name, &remark.Business_Name, &remark.Release_Name, &remark.Text, &remark.Date, &remark.Date_Return, &remark.Status_Description, &remark.User_ID); err != nil {
			return nil, errors.New("error scan remark")
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_Remarks.List = append(lista_Remarks.List, &remark)
		}

	}

	if !hasResult {
		return nil, errors.New("Remarks not found")
	}

	// retorna lista de produtos
	return lista_Remarks, nil
}

// Função que retorna um Remark pelo ID
func (ps *remark_service) GetRemarkByID(ID *uint64) (*entity.Remark, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("call pcGetRemarkByID (?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	remark := entity.Remark{}

	err = stmt.QueryRow(ID).Scan(&remark.ID, &remark.Client_Name, &remark.Client_Email, &remark.Remark_Name, &remark.Subject_Name, &remark.Business_Name, &remark.Release_Name, &remark.Date, &remark.Date_Return, &remark.Text, &remark.Status_Description)
	if err != nil {
		return nil, errors.New("remark not found")
	}

	return &remark, nil
}

// Função que cria um Remark
func (ps *remark_service) CreateRemark(remark *entity.RemarkUpdate) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO tblRemark (remark_subject, remark_text, remark_date, remark_return, subject_id, client_id, release_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(remark.Remark_Name, remark.Text, remark.Date, remark.Date_Return, remark.Subject_ID, remark.Client_ID, remark.Release_ID, remark.User_ID, 21)
	if err != nil {
		return errors.New("error insert remark")
	}

	return nil
}

func (ps *remark_service) GetBarChartRemark(ID *uint64) *entity.Remark {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("call pcGetRemarkByID (?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	remark := entity.Remark{}

	err = stmt.QueryRow(ID).Scan(&remark.ID, &remark.Client_Name, &remark.Client_Email, &remark.Remark_Name, &remark.Business_Name, &remark.Release_Name, &remark.Date, &remark.Date_Return, &remark.Text, &remark.Status_Description)
	if err != nil {
		log.Println("error: cannot find remarkByID", err.Error())
	}

	return &remark

}

func (ps *remark_service) GetPieChartRemark(ID *uint64) *entity.Remark {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("call pcGetRemarkByID (?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	remark := entity.Remark{}

	err = stmt.QueryRow(ID).Scan(&remark.ID, &remark.Client_Name, &remark.Client_Email, &remark.Remark_Name, &remark.Business_Name, &remark.Release_Name, &remark.Date, &remark.Date_Return, &remark.Text, &remark.Status_Description)
	if err != nil {
		log.Println("error: cannot find remarkPieID", err.Error())
	}

	return &remark

}

// Função que atualiza o Status do Remark
func (ps *remark_service) UpdateStatusRemark(ID *uint64, remark *entity.Remark) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblRemark WHERE remark_id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	var statusRemark uint64

	err = stmt.QueryRow(ID).Scan(&statusRemark)
	if err != nil {
		return errors.New("error select status_remark")
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return err
	}

	var statusID uint64

	err = status.QueryRow("REMARK", remark.Status_Description).Scan(&statusID)
	if err != nil {
		return errors.New("error select status")
	}

	if statusRemark == statusID {
		return errors.New("unable to update with the same id, 400")
	}

	updt, err := database.Prepare("UPDATE tblRemark SET status_id = ? WHERE remark_id = ?")
	if err != nil {
		return err
	}

	_, err = updt.Exec(statusID, ID)
	if err != nil {
		return errors.New("error update status")
	}

	return nil
}

// Função que atualiza um Remark
func (ps *remark_service) UpdateRemark(ID *uint64, remark *entity.RemarkUpdate) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblRemark SET remark_subject = ?, remark_text = ?, remark_date = ?, remark_return = ?, subject_id = ?, client_id = ?, release_id = ?, user_id = ?, status_id = ?  WHERE remark_id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(remark.Remark_Name, remark.Text, remark.Date, remark.Date_Return, remark.Subject_ID, remark.Client_ID, remark.Release_ID, remark.User_ID, remark.Status_ID, ID)
	if err != nil {
		return errors.New("error update remark")
	}

	return nil

}
