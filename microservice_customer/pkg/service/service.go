package service

import (
	"fmt"
	"log"

	// Import interno de packages do próprio sistema
	"microservice_customer/pkg/database"
	"microservice_customer/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Customer (tudo que tiver os métodos abaixo do CRUD são serviços de customer)
type CustomerServiceInterface interface {
	// Pega todos os users, logo lista todos os customer
	GetAllCustomer() *entity.CustomerList
	GetCustomerByID(ID *uint64) *entity.Customer
	CreateCustomer(customer *entity.Customer) uint64
	UpdateCustomer(ID *uint64, customer *entity.Customer) uint64
	SoftDeleteCustomer(ID *uint64) int64
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type customer_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewCostumerService(dabase_pool database.DatabaseInterface) *customer_service {
	return &customer_service{
		dabase_pool,
	}
}

// Função que retorna lista de users
func (ps *customer_service) GetAllCustomer() *entity.CustomerList {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("SELECT C.customer_id, C.customer_name, S.status_description FROM tblCustomer C INNER JOIN tblStatus S ON C.status_id = S.status_id")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo CostumerList (vazia)
	lista_customer := &entity.CustomerList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Customer (vazia)
		customer := entity.Customer{}

		// pega dados da query e atribui a variável user, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Status); err != nil {
			fmt.Println(err.Error())
		} else {
			rowsTags, err := database.Query("SELECT tag_name FROM tblTags INNER JOIN tblCustomerTag tCT ON tblTags.tag_id = tCT.tag_id WHERE tCT.customer_id = ?", customer.ID)
			if err != nil {
				fmt.Println(err.Error())
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
					fmt.Println(err.Error())
				} else {
					tags = append(tags, tag)
				}
			}

			customer.Tags = tags

			lista_customer.List = append(lista_customer.List, &customer)
		}
	}

	// retorna lista de produtos
	return lista_customer
}

func (ps *customer_service) GetCustomerByID(ID *uint64) *entity.Customer {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT C.customer_id, C.customer_name, S.status_description FROM tblCustomer C INNER JOIN tblStatus S ON C.status_id = S.status_id WHERE C.customer_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	customer := entity.Customer{}

	err = stmt.QueryRow(ID).Scan(&customer.ID, &customer.Name, &customer.Status)
	if err != nil {
		log.Println("error: cannot find customer", err.Error())
	}

	rowsTags, err := database.Query("SELECT tag_name FROM tblTags INNER JOIN tblCustomerTag tCT ON tblTags.tag_id = tCT.tag_id WHERE tCT.customer_id = ?", ID)
	if err != nil {
		log.Println(err.Error())
	}

	defer rowsTags.Close()

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
			fmt.Println(err.Error())
		} else {
			tags = append(tags, tag)
		}
	}

	customer.Tags = tags

	return &customer
}

func (ps *customer_service) CreateCustomer(customer *entity.Customer) uint64 {
	database := ps.dbp.GetDB()

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("CUSTOMER", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	stmt, err := database.Prepare("INSERT INTO tblCustomer(customer_name,  status_id) VALUES (?, ?)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(customer.Name, statusID)
	if err != nil {
		log.Println(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	stmt, err = database.Prepare("INSERT INTO tblCustomerTag (customer_id, tag_id) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, tag := range customer.Tags {
		_, err := stmt.Exec(lastId, tag.Tag_ID)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
	return uint64(lastId)
}

func (ps *customer_service) UpdateCustomer(ID *uint64, customer *entity.Customer) uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblCustomer SET customer_name = ?, status_id = ? WHERE customer_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(customer.Name, customer.Status, customer.ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return uint64(rowsaff)
}

func (ps *customer_service) SoftDeleteCustomer(ID *uint64) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblCustomer WHERE customer_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	var statusCustomer uint64

	err = stmt.QueryRow(ID).Scan(&statusCustomer)
	if err != nil {
		log.Println(err.Error())
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("CUSTOMER", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	if statusID == statusCustomer {
		statusCustomer++
	} else {
		statusCustomer--
	}

	updt, err := database.Prepare("UPDATE tblCustomer SET status_id = ? WHERE customer_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	result, err := updt.Exec(statusCustomer, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}
