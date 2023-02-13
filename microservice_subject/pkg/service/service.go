package service

import (
	"log"
	"microservice_subject/pkg/entity"
	"microservice_user/pkg/database"
)

type SubjectServiceInterface interface {
	GetSubjectList(id uint64) (*entity.Subject_list, error)
	GetSubject(id uint64) (*entity.SubjectID, error)
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
