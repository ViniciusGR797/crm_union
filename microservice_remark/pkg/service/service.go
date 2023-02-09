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
	GetRemarkByID(ID *int64) *entity.Remark
	CreateRemark(remark *entity.Remark) int64
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
		Remark := entity.Remark{}

		// pega dados da query e atribui a variável Remark, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&Remark.ID, &Remark.User_Name, &Remark.Subject, &Remark.Client_Name, &Remark.Business_Name, &Remark.Release_Name, &Remark.Text, &Remark.Date, &Remark.Date_Return, &Remark.Status_Description); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_Remarks.List = append(lista_Remarks.List, &Remark)
		}

	}

	// retorna lista de produtos
	return lista_Remarks
}

func (ps *remark_service) GetRemarkByID(ID *int64) *entity.Remark {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("select remark_id, remark_subject, remark_text from tblRemark WHERE remark_id = ?")

	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	remark := entity.Remark{}

	err = stmt.QueryRow(ID).Scan(&remark.ID, &remark.Subject, &remark.Text)
	if err != nil {
		log.Println("error: cannot find remarByName", err.Error())
	}

	return &remark
}

func (ps *remark_service) CreateRemark(remark *entity.Remark) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO remark (remark_text, remark_date, remark_date_return) VALUES (?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(remark.ID, remark.Text, remark.Date, remark.Date_Return)
	if err != nil {
		log.Println(err.Error())
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	return lastId
}
