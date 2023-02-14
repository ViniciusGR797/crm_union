package service

import (
	"errors"
	"fmt"
	"log"
	"microservice_business/pkg/database"
	"microservice_business/pkg/entity"
)

type BusinessServiceInterface interface {
	// Pega todos os Businesss, logo lista todos os Businesss
	GetBusiness() *entity.BusinessList
	GetBusinessByID(ID *uint64) *entity.Business
	CreateBusiness(business *entity.CreateBusiness) int64
	UpdateBusiness(ID *uint64, business *entity.Business) uint64
	SoftDeleteBusiness(ID *uint64) int64
	GetBusinessByName(name *string) (*entity.BusinessList, error)
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type Business_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewBusinessService(dabase_pool database.DatabaseInterface) *Business_service {
	return &Business_service{
		dabase_pool,
	}
}

// GetBusiness implements BusinessServiceInterface
func (ps *Business_service) GetBusiness() *entity.BusinessList {

	database := ps.dbp.GetDB()

	rows, err := database.Query("select b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description from tblBusiness b inner join tblDomain d on b.segment_id = d.domain_id inner join  tblStatus s on b.status_id = s.status_id")
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
			rowsTags, err := database.Query("select DISTINCT tag_name from tblTags inner join tblBusinessTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag_name", business.Business_id)
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

func (ps *Business_service) GetBusinessByID(ID *uint64) *entity.Business {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("select b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description from tblBusiness b inner join tblDomain d on b.segment_id = d.domain_id inner join  tblStatus s on b.status_id = s.status_id where b.business_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	Business := entity.Business{}

	err = stmt.QueryRow(ID).Scan(&Business.Business_id, &Business.Business_code, &Business.Business_name, &Business.BusinessSegment.BusinessSegment_id, &Business.BusinessSegment.BusinessSegment_description, &Business.Status.Status_id, &Business.Status.Status_description)
	if err != nil {
		log.Println("error: cannot find business", err.Error())
	}

	rowsTags, err := database.Query("select DISTINCT tag_name from tblTags inner join tblBusinessTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag_name", Business.Business_id)
	if err != nil {
		return &entity.Business{}
	}

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
			return &entity.Business{}
		} else {
			tags = append(tags, tag)
		}
	}

	Business.Tags = tags

	return &Business

}

func (ps *Business_service) CreateBusiness(business *entity.CreateBusiness) int64 {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("insert into tblBusiness (business_code, business_name, segment_id, status_id) values (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(business.Busines_code, business.Business_name, business.Business_Segment_id, business.Business_Status_id)
	if err != nil {
		fmt.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}

	return rowsaff

}

func (ps *Business_service) UpdateBusiness(ID *uint64, business *entity.Business) uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblBusiness SET business_name = ?, business_code = ?, segment_id = ?, status_id = ? WHERE business_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(business.Business_name, business.Business_code, business.BusinessSegment.BusinessSegment_id, business.Status.Status_id, business.Business_id)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return uint64(rowsaff)
}

func (ps *Business_service) SoftDeleteBusiness(ID *uint64) int64 {
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
	lista_Businesss := &entity.BusinessList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Business (vazia)
		Business := entity.Business{}

		// pega dados da query e atribui a variável Business, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&Business.Business_id,
			&Business.Business_code,
			&Business.Business_name,
			&Business.BusinessSegment.BusinessSegment_id,
			&Business.BusinessSegment.BusinessSegment_description,
			&Business.Status.Status_id,
			&Business.Status.Status_description); err != nil {
			log.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a lista de Businesss
			lista_Businesss.List = append(lista_Businesss.List, &Business)
		}

	}

	// retorna lista de Businesss
	return lista_Businesss, nil
}
