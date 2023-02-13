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
			// caso não tenha erro, adiciona a variável log na lista de logs
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

	return &customer
}

func (ps *customer_service) CreateCustomer(customer *entity.Customer) uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO tblCustomer(customer_name,  status_id) VALUES (?, 3)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(customer.Name)
	if err != nil {
		log.Println(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
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

// CRIAR INSERT DE CUSTOMERTAGS ASSOCIATIVA
//UPDATE CUSTOMERTAG?

// stmt, err := database.Prepare("INSERT INTO tblCustomerTag (customer_id, customerTag_id) VALUES (?, ?)")
//     if err != nil {
//         fmt.Println(err.Error())
//     }
// for , user := range users.List {
// 	, err := stmt.Exec(id, user.ID)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// }
