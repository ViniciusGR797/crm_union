package service

import (
	"context"
	"errors"
	"fmt"
	"microservice_subject/pkg/database"
	"microservice_subject/pkg/entity"
)

type SubjectServiceInterface interface {
	GetSubmissiveSubjects(id int, ctx context.Context) (*entity.Subject_list, error)
	GetSubjectByID(id uint64, ctx context.Context) (*entity.SubjectID, error)
	UpdateStatusSubjectFinished(id uint64, logID *int, ctx context.Context) (int64, error)
	UpdateStatusSubjectCanceled(id uint64, logID *int, ctx context.Context) (int64, error)
	CreateSubject(subject *entity.CreateSubject, id uint64, logID *int, ctx context.Context) (*entity.SubjectID, error)
	UpdateSubject(id uint64, subject *entity.UpdateSubject, logID *int, ctx context.Context) (int64, error)
	UpdateStatusSubject(status *entity.UpdateStatus, id uint64, logID *int, ctx context.Context) (int64, error)
}

type Subject_service struct {
	dbp database.DatabaseInterface
}

func NewGroupService(dabase_pool database.DatabaseInterface) *Subject_service {
	return &Subject_service{
		dabase_pool,
	}
}

// GetSubmissiveSubjects retorna uma lista de Subjects de um determinado usuario
func (s *Subject_service) GetSubmissiveSubjects(id int, ctx context.Context) (*entity.Subject_list, error) {
	database := s.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("call pcGetAllUserSubjects (?)", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hasResult := false

	list_subjects := &entity.Subject_list{}

	for rows.Next() {

		hasResult = true

		subject := entity.Subject{}

		if err := rows.Scan(
			&subject.Subject_id,
			&subject.Subject_title,
			&subject.Subject_text,
			&subject.CreatedBy_id,
			&subject.CreatedBy_name,
			&subject.User_ID,
			&subject.User,
			&subject.Release_id,
			&subject.Release,
			&subject.Business_id,
			&subject.Business,
			&subject.Client_id,
			&subject.Client,
			&subject.Client_email,
			&subject.Status.Status_id,
			&subject.Status.Status_description,
			&subject.Created_at,
			&subject.Domain.Domain_id,
			&subject.Domain.Domain_value,
		); err != nil {
			return nil, err
		}

		list_subjects.List = append(list_subjects.List, &subject)
	}

	if !hasResult {
		return nil, fmt.Errorf("no subjects found")
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return list_subjects, nil

}

// GetSubjectByID retorna um Subject pelo id
func (s *Subject_service) GetSubjectByID(id uint64, ctx context.Context) (*entity.SubjectID, error) {
	database := s.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("call pcGetSubjectByID (?)", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hasResult := false

	subject := &entity.SubjectID{}

	for rows.Next() {

		hasResult = true

		if err := rows.Scan(
			&subject.User_id,
			&subject.User_name,
			&subject.Subject_id,
			&subject.Subject_title,
			&subject.Subject_text,
			&subject.CreatedBy_id,
			&subject.CreatedBy_name,
			&subject.Client.Client_id,
			&subject.Client.Client_email,
			&subject.Client.Client_name,
			&subject.Business_id,
			&subject.Business_name,
			&subject.Release_id,
			&subject.Release_name,
			&subject.Created_at,
			&subject.Domain.Domain_id,
			&subject.Domain.Domain_value,
			&subject.Status.Status_id,
			&subject.Status.Status_description,
		); err != nil {
			return nil, err
		}

	}

	if !hasResult {
		return nil, fmt.Errorf("no subjects found")
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return subject, nil

}

// pdateStatusSubjectFinished atualiza o status de um Subject para FINISHED
func (s *Subject_service) UpdateStatusSubjectFinished(id uint64, logID *int, ctx context.Context) (int64, error) {

	database := s.dbp.GetDB()

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

	stmt, err := tx.Prepare("SELECT status_id FROM tblSubject WHERE subject_id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var statusSubject uint64

	err = stmt.QueryRow(id).Scan(&statusSubject)
	if err != nil {
		return 0, err
	}

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}

	var statusID uint64

	err = status.QueryRow("SUBJECT", "FINISHED").Scan(&statusID)
	if err != nil {
		return 0, err
	}

	updt, err := tx.ExecContext(ctx, "UPDATE tblSubject SET status_id = ? WHERE subject_id = ?", statusID, id)
	if err != nil {
		return 0, err
	}

	roww, err := updt.RowsAffected()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return roww, nil

}

// UpdateStatusSubjectCanceled atualiza o status de um Subject para CANCELED
func (s *Subject_service) UpdateStatusSubjectCanceled(id uint64, logID *int, ctx context.Context) (int64, error) {

	database := s.dbp.GetDB()

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

	stmt, err := tx.Prepare("SELECT status_id FROM tblSubject WHERE subject_id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var statusSubject uint64

	err = stmt.QueryRow(id).Scan(&statusSubject)
	if err != nil {
		return 0, err
	}

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}
	defer status.Close()

	var statusID uint64

	err = status.QueryRow("SUBJECT", "CANCELED").Scan(&statusID)
	if err != nil {
		return 0, err
	}

	updt, err := tx.ExecContext(ctx, "UPDATE tblSubject SET status_id = ? WHERE subject_id = ?", statusID, id)
	if err != nil {
		return 0, err
	}

	roww, err := updt.RowsAffected()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return roww, nil

}

// CreateSubject cria um novo Subject
func (s *Subject_service) CreateSubject(subject *entity.CreateSubject, id uint64, logID *int, ctx context.Context) (*entity.SubjectID, error) {

	database := s.dbp.GetDB()

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

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return nil, err
	}
	defer status.Close()

	var statusID uint64

	err = status.QueryRow("SUBJECT", "NOT STARTED").Scan(&statusID)
	if err != nil {
		return nil, err
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO tblSubject (subject_title, subject_text, created_by, subject_type,  client_id, release_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", subject.Subject_title, subject.Subject_text, logID, subject.Subject_type, subject.Client_id, subject.Release_id, id, statusID)
	if err != nil {
		return nil, err
	}

	idresult, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("call pcGetSubjectByID (?)", idresult)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	subjectID := &entity.SubjectID{}

	for rows.Next() {

		if err := rows.Scan(
			&subjectID.User_id,
			&subjectID.User_name,
			&subjectID.Subject_id,
			&subjectID.Subject_title,
			&subjectID.Subject_text,
			&subjectID.CreatedBy_id,
			&subjectID.CreatedBy_name,
			&subjectID.Client.Client_id,
			&subjectID.Client.Client_email,
			&subjectID.Client.Client_name,
			&subjectID.Business_id,
			&subjectID.Business_name,
			&subjectID.Release_id,
			&subjectID.Release_name,
			&subjectID.Created_at,
			&subjectID.Domain.Domain_id,
			&subjectID.Domain.Domain_value,
			&subjectID.Status.Status_id,
			&subjectID.Status.Status_description,
		); err != nil {
			return nil, err
		}

	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return subjectID, nil

}

// UpdateSubject atualiza um Subject
func (s *Subject_service) UpdateSubject(id uint64, subject *entity.UpdateSubject, logID *int, ctx context.Context) (int64, error) {

	database := s.dbp.GetDB()

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

	result, err := tx.ExecContext(ctx, "UPDATE tblSubject SET subject_title = ?, subject_text = ? WHERE subject_id = ?", subject.Subject_title, subject.Subject_text, id)

	if err != nil {
		return 0, err
	}

	roww, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if roww == 0 {
		return 0, fmt.Errorf("no subject found")
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return roww, nil

}

func (s *Subject_service) UpdateStatusSubject(status *entity.UpdateStatus, id uint64, logID *int, ctx context.Context) (int64, error) {

	database := s.dbp.GetDB()

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

	status_id_stmt, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}
	defer status_id_stmt.Close()

	var status_id uint64
	err = status_id_stmt.QueryRow("SUBJECT", &status.Status_description).Scan(&status_id)
	if err != nil {
		return 0, err
	}

	updt, err := tx.ExecContext(ctx, "UPDATE tblSubject SET status_id = ? WHERE subject_id = ?", status_id, id)
	if err != nil {
		return 0, err
	}

	roww, err := updt.RowsAffected()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return roww, nil

}
