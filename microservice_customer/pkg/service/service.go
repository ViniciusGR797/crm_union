package service

import (
	"errors"

	// Import interno de packages do próprio sistema
	"microservice_customer/pkg/database"
	"microservice_customer/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Customer (tudo que tiver os métodos abaixo do CRUD são serviços de customer)
type CustomerServiceInterface interface {
	// Pega todos os users, logo lista todos os customer
	GetCustomers() (*entity.CustomerList, error)
	GetCustomerByID(ID *uint64) (*entity.Customer, error)
	CreateCustomer(customer *entity.Customer, logID *int) error
	UpdateCustomer(ID *uint64, customer *entity.Customer, logID *int) error
	UpdateStatusCustomer(ID *uint64, logID *int) error
}

// ustomer_service Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type customer_service struct {
	dbp database.DatabaseInterface
}

// NewCostumerService Cria novo serviço de CRUD para pool de conexão
func NewCostumerService(dabase_pool database.DatabaseInterface) *customer_service {
	return &customer_service{
		dabase_pool,
	}
}

// GetCustomers Função que retorna lista de users
func (ps *customer_service) GetCustomers() (*entity.CustomerList, error) {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("SELECT DISTINCT C.customer_id, C.customer_name, S.status_description FROM tblCustomer C INNER JOIN tblStatus S ON C.status_id = S.status_id ORDER BY C.customer_name")
	// verifica se teve erro
	if err != nil {
		return nil, err
	}

	// Close fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo CostumerList (vazia)
	lista_customer := &entity.CustomerList{}

	hasResult := false

	// Next Pega todo resultado da query linha por linha
	for rows.Next() {

		hasResult = true

		// variável do tipo Customer (vazia)
		customer := entity.Customer{}

		// Scan pega dados da query e atribui a variável user, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Status); err != nil {
			return nil, errors.New("error scan customer")
		} else {
			rowsTags, err := database.Query("SELECT DISTINCT tag_name FROM tblTags INNER JOIN tblCustomerTag tCT ON tblTags.tag_id = tCT.tag_id WHERE tCT.customer_id = ? ORDER BY tag_name", customer.ID)
			if err != nil {
				return nil, err
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
					return nil, errors.New("error scan tag")
				} else {
					tags = append(tags, tag)
				}
			}

			customer.Tags = tags

			lista_customer.List = append(lista_customer.List, &customer)
		}
	}

	if !hasResult {
		return nil, errors.New("customer not found")
	}

	// retorna lista de customer
	return lista_customer, nil
}

// GetCustomerByID é responsável por buscar um cliente no banco de dados pelo seu ID.
func (ps *customer_service) GetCustomerByID(ID *uint64) (*entity.Customer, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT C.customer_id, C.customer_name, S.status_description FROM tblCustomer C INNER JOIN tblStatus S ON C.status_id = S.status_id WHERE C.customer_id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	customer := entity.Customer{}

	err = stmt.QueryRow(ID).Scan(&customer.ID, &customer.Name, &customer.Status)
	if err != nil {
		return nil, errors.New("error get id")
	}

	rowsTags, err := database.Query("SELECT DISTINCT tag_name FROM tblTags INNER JOIN tblCustomerTag tCT ON tblTags.tag_id = tCT.tag_id WHERE tCT.customer_id = ? ORDER BY tag_name", ID)
	if err != nil {
		return nil, err
	}

	defer rowsTags.Close()

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_Name); err != nil {
			return nil, errors.New("tag not found")
		} else {
			tags = append(tags, tag)
		}
	}

	customer.Tags = tags

	return &customer, nil
}

// CreatCustomer Esta é uma função que cria um novo registro de cliente no banco de dados.
func (ps *customer_service) CreateCustomer(customer *entity.Customer, logID *int) error {
	database := ps.dbp.GetDB()

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return err
	}

	// Definir a variável de sessão "@user"
	_, err = database.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	var statusID uint64

	err = status.QueryRow("CUSTOMER", "ATIVO").Scan(&statusID)
	if err != nil {
		return err
	}

	stmt, err := database.Prepare("INSERT INTO tblCustomer(customer_name,  status_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(customer.Name, statusID)
	if err != nil {
		return errors.New("error create customer")
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	stmt, err = database.Prepare("INSERT INTO tblCustomerTag (customer_id, tag_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	for _, tag := range customer.Tags {
		_, err := stmt.Exec(lastId, tag.Tag_ID)
		if err != nil {
			return errors.New("error insert tag")
		}

	}
	return nil
}

// UpdateCustomer é responsável por atualizar um registro de cliente em um banco de dados.
func (ps *customer_service) UpdateCustomer(ID *uint64, customer *entity.Customer, logID *int) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblCustomer SET customer_name = ? WHERE customer_id = ?")
	if err != nil {
		return nil
	}

	// Definir a variável de sessão "@user"
	_, err = database.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, ID)
	if err != nil {
		return errors.New("error update customer")
	}

	return nil
}

// UpdateStatusCustomer atualiza o status do cliente no banco de dados.
func (ps *customer_service) UpdateStatusCustomer(ID *uint64, logID *int) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblCustomer WHERE customer_id = ?")
	if err != nil {
		return err
	}

	// Definir a variável de sessão "@user"
	_, err = database.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	defer stmt.Close()

	var statusCustomer uint64

	err = stmt.QueryRow(ID).Scan(&statusCustomer)
	if err != nil {
		return err
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return err
	}

	var statusID uint64

	err = status.QueryRow("CUSTOMER", "ATIVO").Scan(&statusID)
	if err != nil {
		return err
	}

	if statusID == statusCustomer {
		statusCustomer++
	} else {
		statusCustomer--
	}

	updt, err := database.Prepare("UPDATE tblCustomer SET status_id = ? WHERE customer_id = ?")
	if err != nil {
		return err
	}

	_, err = updt.Exec(statusCustomer, ID)
	if err != nil {
		return errors.New("error update status")
	}

	return nil
}
