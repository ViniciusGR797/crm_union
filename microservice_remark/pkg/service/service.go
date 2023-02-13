package service

import (
	"fmt"
	"log"

	// Import interno de packages do próprio sistema

	"microservice_remark/pkg/database"
	"microservice_remark/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Remark (tudo que tiver os métodos abaixo do CRUD são serviços de Remark)
type RemarkServiceInterface interface {
	// Pega todos os Remarks, logo lista todos os Remarks
	GetSubmissiveRemarks(ID *uint64) *entity.RemarkList
	GetRemarkByID(ID *uint64) *entity.Remark
	CreateRemark(remark *entity.RemarkUpdate) uint64
	GetBarChartRemark(ID *uint64) *entity.Remark
	GetPieChartRemark(ID *uint64) *entity.Remark
	UpdateStatusRemark(ID *uint64, remark *entity.RemarkUpdate) uint64
	UpdateRemark(ID *uint64, remark *entity.RemarkUpdate) uint64
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
func (ps *remark_service) GetSubmissiveRemarks(ID *uint64) *entity.RemarkList {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("call pcGetAllRemarkUserGroup (?)", ID)
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo RemarkList (vazia)
	lista_Remarks := &entity.RemarkList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Remark (vazia)
		remark := entity.Remark{}

		// pega dados da query e atribui a variável Remark, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&remark.ID, &remark.User_Name, &remark.Remark_Name, &remark.Client_Name, &remark.Business_Name, &remark.Release_Name, &remark.Text, &remark.Date, &remark.Date_Return, &remark.Status_Description); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_Remarks.List = append(lista_Remarks.List, &remark)
		}

	}

	// retorna lista de produtos
	return lista_Remarks
}

func (ps *remark_service) GetRemarkByID(ID *uint64) *entity.Remark {
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

func (ps *remark_service) CreateRemark(remark *entity.RemarkUpdate) uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO tblRemark (remark_subject, remark_text, remark_date, remark_return, subject_id, client_id, release_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(remark.Remark_Name, remark.Text, remark.Date, remark.Date_Return, remark.Subject_ID, remark.Client_ID, remark.Release_ID, remark.User_ID, remark.Status_ID)
	if err != nil {
		log.Println(err.Error())
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	return uint64(lastId)
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

func (ps *remark_service) UpdateStatusRemark(ID *uint64, remark *entity.RemarkUpdate) uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblRemark SET status_id = ? WHERE remark_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(remark.Status_ID, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return uint64(rowsaff)

}

func (ps *remark_service) UpdateRemark(ID *uint64, remark *entity.RemarkUpdate) uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblRemark SET remark_subject = ?, remark_text = ?, remark_date = ?, remark_return = ?, subject_id = ?, client_id = ?, release_id = ?, user_id = ?  WHERE remark_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(remark.Remark_Name, remark.Text, remark.Date, remark.Date_Return, remark.Subject_ID, remark.Client_ID, remark.Release_ID, remark.User_ID, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	newID := uint64(rowsaff)

	return newID

}
