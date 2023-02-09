package service

import (
	"fmt"
	"log"

	// Import interno de packages do próprio sistema
	"microservice_client/pkg/database"
	"microservice_client/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Client (tudo que tiver os métodos abaixo do CRUD são serviços de client)
type ClientServiceInterface interface {
	GetClientsMyGroups(ID *uint64) *entity.ClientList
	UpdateStatusClient(ID *uint64) int64
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type Client_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewClientService(dabase_pool database.DatabaseInterface) *Client_service {
	return &Client_service{
		dabase_pool,
	}
}

// Função que retorna lista de client
func (ps *Client_service) GetClientsMyGroups(ID *uint64) *entity.ClientList {
	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcGetAllClientsGroup (?)", ID)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_client := &entity.ClientList{}

	for rows.Next() {
		client := entity.Client{}

		if err := rows.Scan(&client.Name, &client.Email, &client.Role, &client.Customer_Name, &client.Business_Name, &client.Release_Name, &client.Status_Description); err != nil {
			fmt.Println(err.Error())
		} else {
			list_client.List = append(list_client.List, &client)
		}

	}

	return list_client
}

func (ps *Client_service) UpdateStatusClient(ID *uint64) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblClient WHERE client_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = stmt.QueryRow(ID).Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	if statusID == 11 {
		statusID = 12
	} else {
		statusID = 11
	}

	updt, err := database.Prepare("UPDATE tblClient SET status_id = ? WHERE client_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := updt.Exec(statusID, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}
