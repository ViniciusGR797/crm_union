package service

import (
	"errors"
	"fmt"
	"log"
	"microservice_business/pkg/database"
	"microservice_business/pkg/entity"
)

// BusinessServiceInterface estrutura de dados para BusinessServiceInterface
type BusinessServiceInterface interface {
	// Pega todos os Businesss, logo lista todos os Businesss
	GetBusiness() *entity.BusinessList
	GetBusinessById(ID uint64) (*entity.Business, error)
	GetTagsBusiness(ID *uint64) ([]*entity.Tag, error)
	CreateBusiness(business *entity.Business_Update) error
	UpdateBusiness(ID uint64, business *entity.Business_Update) (uint64, error)
	UpdateStatusBusiness(ID *uint64) int64
	GetBusinessByName(name *string) (*entity.BusinessList, error)
	InsertTagsBusiness(ID uint64, tags []entity.Tag) error
}

// Business_service Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type Business_service struct {
	dbp database.DatabaseInterface
}

// NewBusinessService Cria um novo serviço de CRUD para pool de conexão
func NewBusinessService(dabase_pool database.DatabaseInterface) *Business_service {
	return &Business_service{
		dabase_pool,
	}
}

// GetBusiness traz todos os Business do banco de dados
func (ps *Business_service) GetBusiness() *entity.BusinessList {

	database := ps.dbp.GetDB()

	rows, err := database.Query("SELECT DISTINCT b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description FROM tblBusiness b INNER JOIN tblDomain d on b.segment_id = d.domain_id INNER JOIN  tblStatus s on b.status_id = s.status_id ORDER BY b.business_name")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_Business := &entity.BusinessList{}

	for rows.Next() {
		business := entity.Business{}

		if err := rows.Scan(&business.Business_id, &business.Business_code, &business.Business_name, &business.BusinessSegment.BusinessSegment_id, &business.BusinessSegment.BusinessSegment_description, &business.Status.Status_id, &business.Status.Status_description); err != nil {
			return &entity.BusinessList{}
		} else {
			rowsTags, err := database.Query("SELECT DISTINCT tag_name from tblTags inner join tblBusinessTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag_name", business.Business_id)
			if err != nil {
				return &entity.BusinessList{}
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
					return &entity.BusinessList{}
				} else {
					tags = append(tags, tag)
				}
			}

			business.Tags = tags

			list_Business.List = append(list_Business.List, &business)
		}
	}

	return list_Business

}

// GetBusinessById traz um usuario no banco de dados pelo ID do mesmo
func (ps *Business_service) GetBusinessById(ID uint64) (*entity.Business, error) {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description FROM tblBusiness b INNER JOIN tblDomain d on b.segment_id = d.domain_id INNER JOIN  tblStatus s on b.status_id = s.status_id WHERE b.business_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	Business := &entity.Business{}

	err = stmt.QueryRow(ID).Scan(&Business.Business_id, &Business.Business_code, &Business.Business_name, &Business.BusinessSegment.BusinessSegment_id, &Business.BusinessSegment.BusinessSegment_description, &Business.Status.Status_id, &Business.Status.Status_description)
	if err != nil {
		return &entity.Business{}, errors.New("error scanning rows")
	}

	rowsTags, err := database.Query("SELECT DISTINCT tag_name from tblTags INNER JOIN tblBusinessTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag_name", Business.Business_id)
	if err != nil {
		return &entity.Business{}, errors.New("error fetching tags from business by id")
	}

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
			return &entity.Business{}, errors.New("error scanning tags from business by id")
		} else {
			tags = append(tags, tag)
		}
	}

	Business.Tags = tags

	return Business, nil
}

// CreateBusiness cria um Business no banco de dados
func (ps *Business_service) CreateBusiness(business *entity.Business_Update) error {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO tblBusiness (business_code, business_name, segment_id, status_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(business.Code, business.Name, business.Segment_Id, 1)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = result.RowsAffected()
	if err != nil {
		return errors.New("error rowAffected insert into database")
	}

	ID, _ := result.LastInsertId()
	business.ID = uint64(ID)

	stmt, err = database.Prepare("INSERT tblBusinessTag SET tag_id = ?, business_id = ?")
	if err != nil {
		return errors.New("error in prepare rusiness tags statement")
	}

	for _, tag := range business.Tags {
		_, err := stmt.Exec(tag.Tag_ID, business.ID)
		if err != nil {
			return errors.New("business tags exec error")
		}
	}

	return nil

}

// UpdateBusiness atualiza os dados de um Bussines no banco pelo ID do mesmo
func (ps *Business_service) UpdateBusiness(ID uint64, business *entity.Business_Update) (uint64, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblBusiness SET business_name = ?, business_code = ?, segment_id = ? WHERE business_id = ?")
	if err != nil {
		return 0, errors.New("error prepare update business")
	}

	defer stmt.Close()

	var businessID int64

	result, err := stmt.Exec(business.Name, business.Code, business.Segment_Id, ID)
	if err != nil {
		return 0, errors.New("error exec update business")
	}

	businessID, err = result.RowsAffected()
	if err != nil {
		return 0, errors.New("error RowsAffected update business")
	}

	if business.Tags != nil {
		err = ps.InsertTagsBusiness(ID, business.Tags)
		if err != nil {
			return 0, errors.New("error in update business tags statement")
		}
	}

	return uint64(businessID), nil
}

// UpdateStatusBusiness altera o status de um usuario no banco
func (ps *Business_service) UpdateStatusBusiness(ID *uint64) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblBusiness WHERE business_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	var statusBusiness uint64

	err = stmt.QueryRow(ID).Scan(&statusBusiness)
	if err != nil {
		log.Println(err.Error())
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("Business", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	if statusID == statusBusiness {
		statusBusiness++
	} else {
		statusBusiness--
	}

	updt, err := database.Prepare("UPDATE tblBusiness SET status_id = ? WHERE business_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	result, err := updt.Exec(statusBusiness, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}

// GetBusinessByName busca Business no banco de dados pelo nome passado como parâmetro no query.
func (ps *Business_service) GetBusinessByName(name *string) (*entity.BusinessList, error) {
	nameString := fmt.Sprint("%", *name, "%")
	query := "SELECT DISTINCT b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description FROM tblBusiness b inner join tblDomain d on b.segment_id = d.domain_id inner join  tblStatus s on b.status_id = s.status_id WHERE b.business_name LIKE ? ORDER BY b.business_name"

	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query(query, nameString)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return &entity.BusinessList{}, errors.New("error fetching Businesss")
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo BusinessList (vazia)
	list_Business := &entity.BusinessList{}

	for rows.Next() {
		business := entity.Business{}

		if err := rows.Scan(&business.Business_id, &business.Business_code, &business.Business_name, &business.BusinessSegment.BusinessSegment_id, &business.BusinessSegment.BusinessSegment_description, &business.Status.Status_id, &business.Status.Status_description); err != nil {
			return &entity.BusinessList{}, nil
		} else {

			rowsTags, err := database.Query("SELECT DISTINCT tag_name from tblTags inner join tblBusinessTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag_name", business.Business_id)
			if err != nil {
				return &entity.BusinessList{}, nil
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
					return &entity.BusinessList{}, nil
				} else {
					tags = append(tags, tag)
				}
			}

			business.Tags = tags

			list_Business.List = append(list_Business.List, &business)
		}
	}

	return list_Business, nil
}

// InsertTagsBusiness insere tags na Business
func (ps *Business_service) InsertTagsBusiness(ID uint64, tags []entity.Tag) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM tblBusinessTag WHERE business_id = ?")
	if err != nil {
		return errors.New("error prepare delete tags on business")
	}

	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {

		return errors.New("error exec statement exec on business")
	}

	stmt, err = database.Prepare("INSERT IGNORE tblBusinessTag SET tag_id = ?, business_id = ?")
	if err != nil {
		return errors.New("error insert a new row on tag_id and business_id")
	}

	defer stmt.Close()

	for _, tag := range tags {
		_, err := stmt.Exec(tag.Tag_ID, ID)
		if err != nil {
			return errors.New("error insert data tag_ID and ID on database")
		}
	}

	return nil
}

// GetTagsBusiness busca as tags de Business
func (ps *Business_service) GetTagsBusiness(ID *uint64) ([]*entity.Tag, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT DISTINCT T.tag_id, T.tag_name from tblTags T INNER JOIN tblBusinessTag B on T.tag_id = B.tag_id WHERE business_id = ? ORDER BY T.tag_name")
	if err != nil {
		return []*entity.Tag{}, errors.New("error fetching on tag business")
	}

	defer stmt.Close()

	var tags []*entity.Tag

	rowsTags, err := stmt.Query(ID)
	if err != nil {
		return []*entity.Tag{}, errors.New("error fetching on row tags query release train")
	}

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			return []*entity.Tag{}, errors.New("error fetching on row tags next release train")
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}
