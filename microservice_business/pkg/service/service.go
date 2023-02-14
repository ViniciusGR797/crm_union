package service

import (
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
		log.Println("error: cannot find customer", err.Error())
	}

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
