package service

import (
	"fmt"
	"microservice_subject/pkg/entity"
	"microservice_user/pkg/database"
)

type SubjectServiceInterface interface {
	GetSubmissiveSubjects(id int) (*entity.Subject_list, error)
	GetSubjectByID(id uint64) (*entity.SubjectID, error)
	UpdateStatusSubjectFinished(id uint64) (int64, error)
	UpdateStatusSubjectCanceled(id uint64) (int64, error)
	CreateSubject(subject *entity.CreateSubject, id uint64) (*entity.SubjectID, error)
	UpdateSubject(id uint64, subject *entity.UpdateSubject) (int64, error)
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
func (s *Subject_service) GetSubmissiveSubjects(id int) (*entity.Subject_list, error) {

	database := s.dbp.GetDB()

	rows, err := database.Query("call pcGetAllUserSubjects (?)", id)
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
			&subject.User,
			&subject.Release,
			&subject.Business,
			&subject.Client,
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

	return list_subjects, nil

}

// GetSubjectByID retorna um Subject pelo id
func (s *Subject_service) GetSubjectByID(id uint64) (*entity.SubjectID, error) {

	database := s.dbp.GetDB()

	rows, err := database.Query("call pcGetSubjectByID (?)", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hasResult := false

	subject := &entity.SubjectID{}

	for rows.Next() {

		hasResult = true

		if err := rows.Scan(
			&subject.Subject_id,
			&subject.Subject_title,
			&subject.Client.Client_id,
			&subject.Client.Client_email,
			&subject.Client.Client_name,
			&subject.Business_name,
			&subject.Release_name,
			&subject.Subject_text,
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

	return subject, nil

}

// pdateStatusSubjectFinished atualiza o status de um Subject para FINISHED
func (s *Subject_service) UpdateStatusSubjectFinished(id uint64) (int64, error) {

	database := s.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblSubject WHERE subject_id = ?")
	if err != nil {
		return 0, err
	}

	var statusSubject uint64

	err = stmt.QueryRow(id).Scan(&statusSubject)
	if err != nil {
		return 0, err
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}

	var statusID uint64

	err = status.QueryRow("SUBJECT", "FINISHED").Scan(&statusID)
	if err != nil {
		return 0, err
	}

	updt, err := database.Prepare("UPDATE tblSubject SET status_id = ? WHERE subject_id = ?")
	if err != nil {
		return 0, err
	}

	result, err := updt.Exec(statusID, id)
	if err != nil {
		return 0, err
	}

	roww, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return roww, nil

}

// UpdateStatusSubjectCanceled atualiza o status de um Subject para CANCELED
func (s *Subject_service) UpdateStatusSubjectCanceled(id uint64) (int64, error) {

	database := s.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblSubject WHERE subject_id = ?")
	if err != nil {
		return 0, err
	}

	var statusSubject uint64

	err = stmt.QueryRow(id).Scan(&statusSubject)
	if err != nil {
		return 0, err
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}

	var statusID uint64

	err = status.QueryRow("SUBJECT", "CANCELED").Scan(&statusID)
	if err != nil {
		return 0, err
	}

	updt, err := database.Prepare("UPDATE tblSubject SET status_id = ? WHERE subject_id = ?")
	if err != nil {
		return 0, err
	}

	result, err := updt.Exec(statusID, id)
	if err != nil {
		return 0, err
	}

	roww, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return roww, nil

}

// CreateSubject cria um novo Subject
func (s *Subject_service) CreateSubject(subject *entity.CreateSubject, id uint64) (*entity.SubjectID, error) {

	database := s.dbp.GetDB()

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return nil, err
	}

	var statusID uint64

	err = status.QueryRow("SUBJECT", "IN PROGRESS").Scan(&statusID)
	if err != nil {
		return nil, err
	}

	stmt, err := database.Prepare("INSERT INTO tblSubject (subject_title, subject_text, subject_type,  client_id, release_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?,?)")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(subject.Subject_title, subject.Subject_text, subject.Subject_type, subject.Client_id, subject.Release_id, id, statusID)
	if err != nil {
		return nil, err
	}

	idresult, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	rows, err := database.Query("call pcGetSubjectByID (?)", idresult)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	subjectID := &entity.SubjectID{}

	for rows.Next() {

		if err := rows.Scan(
			&subjectID.Subject_id,
			&subjectID.Subject_title,
			&subjectID.Client.Client_id,
			&subjectID.Client.Client_email,
			&subjectID.Client.Client_name,
			&subjectID.Business_name,
			&subjectID.Release_name,
			&subjectID.Subject_text,
			&subjectID.Created_at,
			&subjectID.Domain.Domain_id,
			&subjectID.Domain.Domain_value,
			&subjectID.Status.Status_id,
			&subjectID.Status.Status_description,
		); err != nil {
			return nil, err
		}

	}

	return subjectID, nil

}

// UpdateSubject atualiza um Subject
func (s *Subject_service) UpdateSubject(id uint64, subject *entity.UpdateSubject) (int64, error) {

	database := s.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblSubject SET subject_title = ?, subject_text = ? WHERE subject_id = ?")

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(subject.Subject_title, subject.Subject_text, id)
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

	return roww, nil

}
