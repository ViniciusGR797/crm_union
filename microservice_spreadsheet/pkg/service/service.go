package service

import (
	"context"
	"errors"

	// Import interno de packages do próprio sistema

	"microservice_spreadsheet/pkg/database"
	"microservice_spreadsheet/pkg/entity"
)

// RemarkServiceInterface Estrutura interface para padronizar comportamento de CRUD Remark (tudo que tiver os métodos abaixo do CRUD são serviços de Remark)
type RemarkServiceInterface interface {
	// Pega todos os Remarks, logo lista todos os Remarks
	RemarkFilter(ID *int, filter *entity.RemarkFilter, ctx context.Context) (*entity.RemarkList, error)
}

// remark_service Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type remark_service struct {
	dbp database.DatabaseInterface
}

// NewRemarkService Cria novo serviço de CRUD para pool de conexão
func NewRemarkService(dabase_pool database.DatabaseInterface) *remark_service {
	return &remark_service{
		dabase_pool,
	}
}

// GetSubmissiveRemarks Função que retorna lista de Remarks
func (ps *remark_service) RemarkFilter(ID *int, filter *entity.RemarkFilter, ctx context.Context) (*entity.RemarkList, error) {
	// pega database
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// manda uma query para ser executada no database
	rows, err := tx.Query("call pcGetAllRemarks(?, ?, ?)", filter.User_ID, filter.Date, filter.Date_Return)
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
		if err := rows.Scan(&remark.ID, &remark.Subject_Name, &remark.Text, &remark.Date, &remark.Date_Return, &remark.Subject_Title, &remark.Client_Name, &remark.Release_Name, &remark.User_Name, &remark.CreatedBy_name, &remark.Status_Description); err != nil {
			return nil, errors.New("error scan remark")
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			lista_Remarks.List = append(lista_Remarks.List, &remark)
		}

	}

	if !hasResult {
		return nil, errors.New("remarks not found")
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// lista_Remarks retorna lista de produtos
	return lista_Remarks, nil
}
