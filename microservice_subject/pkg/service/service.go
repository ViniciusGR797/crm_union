package service

import (
	"fmt"
	"log"
	"microservice_subject/pkg/entity"
	"microservice_user/pkg/database"
)

type SubjectServiceInterface interface {
	GetSubjectList(id uint64) (*entity.Subject_list, error)
	GetSubject(id uint64) (*entity.SubjectID, error)
	UpdateStatusSubjectFinished(id uint64) (int64, error)
	UpdateStatusSubjectCanceled(id uint64) (int64, error)
	CreateSubject(subject *entity.CreateSubject, id uint64) (*entity.SubjectID, error)
}

type Subject_service struct {
	dbp database.DatabaseInterface
}

func NewGroupService(dabase_pool database.DatabaseInterface) *Subject_service {
	return &Subject_service{
		dabase_pool,
	}
}

func (s *Subject_service) GetSubjectList(id uint64) (*entity.Subject_list, error) {

	database := s.dbp.GetDB()

	rows, err := database.Query("call pcGetAllUserSubjects (?)", id)
	if err != nil {
		log.Println(err.Error())
	}

	defer rows.Close()

	list_subjects := &entity.Subject_list{}

	for rows.Next() {

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
		); err != nil {
			log.Println(err.Error())
		}

		list_subjects.List = append(list_subjects.List, &subject)
	}

	return list_subjects, nil

}

func (s *Subject_service) GetSubject(id uint64) (*entity.SubjectID, error) {

	database := s.dbp.GetDB()

	rows, err := database.Query("call pcGetSubjectByID (?)", id)
	if err != nil {
		log.Println(err.Error())
	}

	defer rows.Close()

	subject := &entity.SubjectID{}

	for rows.Next() {

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
		); err != nil {
			log.Println(err.Error())
		}

	}

	return subject, nil

}

func (s *Subject_service) UpdateStatusSubjectFinished(id uint64) (int64, error) {

	database := s.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblSubject WHERE subject_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	var statusSubject uint64

	err = stmt.QueryRow(id).Scan(&statusSubject)
	if err != nil {
		log.Println(err.Error())
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("SUBJECT", "FINISHED").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	updt, err := database.Prepare("UPDATE tblSubject SET status_id = ? WHERE subject_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	result, err := updt.Exec(statusID, id)
	if err != nil {
		log.Println(err.Error())
	}

	roww, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return roww, nil

}

func (s *Subject_service) UpdateStatusSubjectCanceled(id uint64) (int64, error) {

	database := s.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblSubject WHERE subject_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	var statusSubject uint64

	err = stmt.QueryRow(id).Scan(&statusSubject)
	if err != nil {
		log.Println(err.Error())
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("SUBJECT", "CANCELED").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	updt, err := database.Prepare("UPDATE tblSubject SET status_id = ? WHERE subject_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	result, err := updt.Exec(statusID, id)
	if err != nil {
		log.Println(err.Error())
	}

	roww, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return roww, nil

}

func (s *Subject_service) CreateSubject(subject *entity.CreateSubject, id uint64) (*entity.SubjectID, error) {

	database := s.dbp.GetDB()

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("SUBJECT", "IN PROGRESS").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	stmt, err := database.Prepare("INSERT INTO tblSubject (subject_title, subject_text, subject_type,  client_id, release_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?,?)")
	if err != nil {
		log.Println(err.Error())
	}

	result, err := stmt.Exec(subject.Subject_title, subject.Subject_text, subject.Subject_type, subject.Client_id, subject.Release_id, id, statusID)
	if err != nil {
		log.Println(err.Error())
	}

	idresult, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := database.Query("call pcGetSubjectByID (?)", idresult)
	if err != nil {
		log.Println(err.Error())
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
		); err != nil {
			log.Println(err.Error())
		}

	}

	return subjectID, nil

}
