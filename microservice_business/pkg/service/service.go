package service

import (
	"fmt"
	"microservice_business/pkg/database"
	"microservice_business/pkg/entity"
)

type BusinessServiceInterface interface {
	// Pega todos os Businesss, logo lista todos os Businesss
	GetBusiness() *entity.BusinessList
	GetBusinessByID(id uint64) (*entity.Business, error)
	CreateBusiness(business *entity.CreateBusiness) int64
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

	rows, err := database.Query("select b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description from tblBusiness b inner join tblDomain d on b.segment_id = d.domain_id inner join  tblStatus s on b.status_id = s.status_id;")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	lista_Business := &entity.BusinessList{}

	for rows.Next() {

		business := entity.Business{}

		if err := rows.Scan(
			&business.Business_id,
			&business.Business_code,
			&business.Business_name,
			&business.BusinessSegment.BusinessSegment_id,
			&business.BusinessSegment.BusinessSegment_description,
			&business.Status.Status_id,
			&business.Status.Status_description); err != nil {
			fmt.Println(err.Error())
		} else {

			lista_Business.List = append(lista_Business.List, &business)
		}

	}

	return lista_Business

}

func (ps *Business_service) GetBusinessByID(id uint64) (*entity.Business, error) {

	database := ps.dbp.GetDB()

	rows, err := database.Query("select b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description from tblBusiness b inner join tblDomain d on b.segment_id = d.domain_id inner join  tblStatus s on b.status_id = s.status_id where b.business_id = ?", id)

	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	Business := &entity.Business{}

	if rows.Next() {
		if err := rows.Scan(
			&Business.Business_id,
			&Business.Business_code,
			&Business.Business_name,
			&Business.BusinessSegment.BusinessSegment_id,
			&Business.BusinessSegment.BusinessSegment_description,
			&Business.Status.Status_id,
			&Business.Status.Status_description); err != nil {
			return &entity.Business{}, err
		}
	}

	return Business, nil

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
