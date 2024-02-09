package service

import (
	"context"
	"database/sql"
	"errors"

	// Import interno de packages do próprio sistema

	"microservice_remark/pkg/database"
	"microservice_remark/pkg/entity"
)

// RemarkServiceInterface Estrutura interface para padronizar comportamento de CRUD Remark (tudo que tiver os métodos abaixo do CRUD são serviços de Remark)
type RemarkServiceInterface interface {
	// Pega todos os Remarks, logo lista todos os Remarks
	GetSubmissiveRemarks(ID *int, ctx context.Context) (*entity.RemarkList, error)
	GetAllRemarkUser(ID *uint64, ctx context.Context) (*entity.RemarkList, error)
	GetRemarkByID(ID *uint64, ctx context.Context) (*entity.Remark, error)
	CreateRemark(remark *entity.RemarkUpdate, logID *int, ctx context.Context) (*entity.Remark, error)
	GetBarChartRemark(ID *uint64, ctx context.Context) (*entity.Remark, error)
	GetPieChartRemark(ID *uint64, ctx context.Context) (*entity.Remark, error)
	UpdateStatusRemark(ID *uint64, remark *entity.Remark, logID *int, ctx context.Context) error
	UpdateRemark(ID *uint64, remark *entity.RemarkUpdate, logID *int, ctx context.Context) error
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
func (ps *remark_service) GetSubmissiveRemarks(ID *int, ctx context.Context) (*entity.RemarkList, error) {
	// pega database
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// manda uma query para ser executada no database
	rows, err := tx.Query("call pcGetAllRemarkUserGroup (?)", ID)
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
		if err := rows.Scan(&remark.ID, &remark.Remark_Name, &remark.User_Name, &remark.Subject_ID, &remark.Subject_Name, &remark.Client_ID, &remark.Client_Name, &remark.Client_Role, &remark.Domain_Value, &remark.Business_ID, &remark.Business_Name, &remark.Release_ID, &remark.Release_Name, &remark.Text, &remark.Date, &remark.Date_Return, &remark.Status_Description, &remark.User_ID, &remark.CreatedBy_id, &remark.CreatedBy_name); err != nil {
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

// GetAllRemarkUser Função que retorna os Remarks de um ID
func (ps *remark_service) GetAllRemarkUser(ID *uint64, ctx context.Context) (*entity.RemarkList, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("call pcGetAllRemarkUser (?)", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	listRemark := entity.RemarkList{}

	for rows.Next() {
		remark := entity.Remark{}
		err = rows.Scan(&remark.ID, &remark.Remark_Name, &remark.Subject_Name, &remark.Client_Name, &remark.Client_Email, &remark.Business_Name, &remark.Release_Name, &remark.Text, &remark.Date, &remark.Date_Return, &remark.CreatedBy_id, &remark.CreatedBy_name)
		if err != nil {
			return nil, errors.New("remark not found")
		} else {
			listRemark.List = append(listRemark.List, &remark)
		}

	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &listRemark, nil
}

// GetRemarkByID Função que retorna um Remark pelo ID
func (ps *remark_service) GetRemarkByID(ID *uint64, ctx context.Context) (*entity.Remark, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("call pcGetRemarkByID (?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	remark := entity.Remark{}

	err = stmt.QueryRow(ID).Scan(&remark.ID, &remark.Client_ID, &remark.Client_Name, &remark.Client_Email, &remark.Remark_Name, &remark.Subject_ID, &remark.Subject_Title, &remark.Business_ID, &remark.Business_Name, &remark.Release_ID, &remark.Release_Name, &remark.Date, &remark.Date_Return, &remark.Text, &remark.CreatedBy_id, &remark.CreatedBy_name, &remark.Status_Description)
	if err != nil {
		return nil, errors.New("remark not found")
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &remark, nil
}

// CreateRemark que usa uma estrutura RemarkUpdate como argumento e retorna um erro. Função que cria um Remark
func (ps *remark_service) CreateRemark(remark *entity.RemarkUpdate, logID *int, ctx context.Context) (*entity.Remark, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return nil, errors.New("session variable error")
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO tblRemark (remark_subject, remark_text, remark_date, remark_return, subject_id, client_id, release_id, user_id, created_by, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", remark.Remark_Name, remark.Text, remark.Date, remark.Date_Return, remark.Subject_ID, remark.Client_ID, remark.Release_ID, remark.User_ID, logID, 21)
	if err != nil {
		return nil, err
	}

	idresult, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("call pcGetRemarkByID (?)", idresult)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	remarkID := &entity.Remark{}

	for rows.Next() {
		if err := rows.Scan(&remarkID.ID,
			&remarkID.Client_ID,
			&remarkID.Client_Name,
			&remarkID.Client_Email,
			&remarkID.Subject_Name,
			&remarkID.Subject_ID,
			&remarkID.Subject_Title,
			&remarkID.Business_ID,
			&remarkID.Business_Name,
			&remarkID.Release_ID,
			&remarkID.Release_Name,
			&remarkID.Date,
			&remarkID.Date_Return,
			&remarkID.Text,
			&remarkID.CreatedBy_id,
			&remarkID.CreatedBy_name,
			&remarkID.Status_Description,
		); err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return remarkID, nil
}

// GetBarChartRemark retorna um gráfico de barras mostrando a contagem de avaliações em relação ao tempo (atrasado, próximo, no prazo) para o usuário com o ID especificado na URL, disparando o método controller.GetBarChartRemark.
func (ps *remark_service) GetBarChartRemark(ID *uint64, ctx context.Context) (*entity.Remark, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return &entity.Remark{}, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("call pcGetRemarkByID (?)")
	if err != nil {
		return &entity.Remark{}, err
	}

	defer stmt.Close()

	remark := entity.Remark{}

	err = stmt.QueryRow(ID).Scan(&remark.ID, &remark.Client_ID, &remark.Client_Name, &remark.Client_Email, &remark.Remark_Name, &remark.Subject_ID, &remark.Subject_Title, &remark.Business_ID, &remark.Business_Name, &remark.Release_ID, &remark.Release_Name, &remark.Date, &remark.Date_Return, &remark.Text, &remark.CreatedBy_id, &remark.CreatedBy_name, &remark.Status_Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return &entity.Remark{}, nil
		} else {
			return &entity.Remark{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return &entity.Remark{}, err
	}

	return &remark, nil
}

// GetPieChartRemark retorna um gráfico de pizza mostrando a contagem de avaliações em relação ao status (pendente, aprovado, rejeitado) para o usuário com o ID especificado na URL, disparando o método controller.GetPieChartRemark
func (ps *remark_service) GetPieChartRemark(ID *uint64, ctx context.Context) (*entity.Remark, error) {
	database := ps.dbp.GetDB()

	println("ok - 1")

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return &entity.Remark{}, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("call pcGetRemarkByID (?)")
	if err != nil {
		return &entity.Remark{}, err
	}

	defer stmt.Close()

	remark := entity.Remark{}

	err = stmt.QueryRow(ID).Scan(&remark.ID, &remark.Client_ID, &remark.Client_Name, &remark.Client_Email, &remark.Remark_Name, &remark.Subject_ID, &remark.Subject_Title, &remark.Business_ID, &remark.Business_Name, &remark.Release_ID, &remark.Release_Name, &remark.Date, &remark.Date_Return, &remark.Text, &remark.CreatedBy_id, &remark.CreatedBy_name, &remark.Status_Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return &entity.Remark{}, nil
		} else {
			return &entity.Remark{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return &entity.Remark{}, err
	}

	return &remark, nil
}

// UpdateStatusRemark Função que atualiza o Status do Remark
func (ps *remark_service) UpdateStatusRemark(ID *uint64, remark *entity.Remark, logID *int, ctx context.Context) error {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	stmt, err := tx.Prepare("SELECT status_id FROM tblRemark WHERE remark_id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	var statusRemark uint64

	err = stmt.QueryRow(ID).Scan(&statusRemark)
	if err != nil {
		return errors.New("error select status_remark")
	}

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return err
	}

	var statusID uint64

	err = status.QueryRow("REMARK", remark.Status_Description).Scan(&statusID)
	if err != nil {
		return errors.New("error select status")
	}

	if statusRemark == statusID {
		return nil
	}

	_, err = tx.ExecContext(ctx, "UPDATE IGNORE tblRemark SET status_id = ? WHERE remark_id = ?", statusID, ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// UpdateRemark Função que atualiza um Remark
func (ps *remark_service) UpdateRemark(ID *uint64, remark *entity.RemarkUpdate, logID *int, ctx context.Context) error {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	_, err = tx.ExecContext(ctx, "UPDATE IGNORE tblRemark SET remark_subject = ?, remark_text = ?, remark_date = ?, remark_return = ?, subject_id = ?, client_id = ?, release_id = ?, user_id = ?, status_id = ?  WHERE remark_id = ?", remark.Remark_Name, remark.Text, remark.Date, remark.Date_Return, remark.Subject_ID, remark.Client_ID, remark.Release_ID, remark.User_ID, remark.Status_ID, ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
