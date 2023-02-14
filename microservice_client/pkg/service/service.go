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
	CreateClient(client *entity.ClientUpdate) int64
	UpdateStatusClient(ID *uint64) int64
	InsertTagClient(ID *uint64, client *entity.ClientUpdate) *uint64
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
			rowsTags, err := database.Query("SELECT tT.tag_id, tT.tag_name FROM tblTags tT INNER JOIN tblClientTag tCT ON tT.tag_id = tCT.tag_id WHERE tCT.client_id = ?", client.ID)
			if err != nil {
				fmt.Println(err.Error())
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
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

// Função que retorna um client pelo ID
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

	rowsTags, err := database.Query("SELECT tT.tag_id, tT.tag_name FROM tblTags tT INNER JOIN tblClientTag tCT ON tT.tag_id = tCT.tag_id WHERE tCT.client_id = ?", ID)
	if err != nil {
		log.Println(err.Error())
	}

	defer rowsTags.Close()

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			fmt.Println(err.Error())
		} else {
			tags = append(tags, tag)
		}
	}

	client.Tags = tags

	return &client
}

// Função que retorna uma lista de tag pelo ID do client
func (ps *Client_service) GetTagsClient(ID *uint64) []*entity.Tag {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("select T.tag_id, T.tag_name from tblTags T inner join tblClientTag TCT on T.tag_id = TCT.tag_id WHERE client_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	var tags []*entity.Tag

	rowsTags, err := stmt.Query(ID)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			fmt.Println(err.Error())
		}

		tags = append(tags, &tag)
	}

	return tags

}

// Função utilizada para criar um novo client
func (ps *Client_service) CreateClient(client *entity.ClientUpdate) int64 {
	database := ps.dbp.GetDB()

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("CLIENT", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	stmt, err := database.Prepare("INSERT INTO tblClient(client_name, client_email, client_role, customer_id, release_id, business_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(&client.Name, &client.Email, &client.Role, &client.Customer_ID, &client.Release_ID, &client.Business_ID, &client.User_ID, statusID)
	if err != nil {
		fmt.Println(err.Error())
	}

	ID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}

	newID := uint64(ID)

	if client.Tags != nil {
		ps.InsertTagClient(&newID, client)
	}

	return ID
}

func (ps *Client_service) InsertTagClient(ID *uint64, client *entity.ClientUpdate) *uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO tblClientTag(client_id, tag_id) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	for _, tag := range client.Tags {
		_, err := stmt.Exec(ID, tag.Tag_ID)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return ID
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
