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
	GetClientByID(ID *uint64) *entity.Client
	GetTagsClient(ID *uint64) []*entity.Tag
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

// Função que retorna lista de client pelo group
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
			rowsTags, err := database.Query("SELECT tag_name FROM tblTags INNER JOIN tblClientTag tCT ON tblTags.tag_id = tCT.tag_id WHERE tCT.client_id = ?", client.ID)
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

func (ps *Client_service) GetClientByID(ID *uint64) *entity.Client {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("call pcGetClientByID(?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	var client entity.Client

	err = stmt.QueryRow(ID).Scan(&client.ID, &client.Name, &client.Email, &client.Role, &client.Customer_Name, &client.Business_Name, &client.Release_Name, &client.User_Name, &client.Status_Description)
	if err != nil {
		fmt.Println(err.Error())
	}

	rowsTags, err := database.Query("SELECT tag_name FROM tblTags INNER JOIN tblClientTag tCT ON tblTags.tag_id = tCT.tag_id WHERE tCT.client_id = ?", ID)
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

	client.Tags = tags

	return &client
}

func (ps *Client_service) GetTagsClient(ID *uint64) []*entity.Tag {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("select T.tag_id, T.tag_name from tblTags T inner join tblClientTag TCT on T.tag_id = TCT.tag_id WHERE client_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	var tags []*entity.Tag

	rowsTags, err := stmt.Query(ID)

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			fmt.Println(err.Error())
		}

		tags = append(tags, &tag)
	}

	return tags

}

// Função que atualizar o status do client
func (ps *Client_service) UpdateStatusClient(ID *uint64) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblClient WHERE client_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

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
