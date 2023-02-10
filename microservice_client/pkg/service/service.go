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

		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Role, &client.Customer_Name, &client.Business_Name, &client.Release_Name, &client.User_Name, &client.Status_Description); err != nil {
			fmt.Println(err.Error())
		} else {
			rowsTags, err := database.Query("select tag_name from tblTags inner join tblClientTag tCT on tblTags.tag_id = tCT.tag_id WHERE tCT.client_id = ?", client.ID)
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

			client.Tags = tags

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

	var statusClient uint64

	err = stmt.QueryRow(ID).Scan(&statusClient)
	if err != nil {
		log.Println(err.Error())
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("CLIENT", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	if statusID == statusClient {
		statusClient++
	} else {
		statusClient--
	}

	updt, err := database.Prepare("UPDATE tblClient SET status_id = ? WHERE client_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := updt.Exec(statusClient, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}
