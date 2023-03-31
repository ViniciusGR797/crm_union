package service

import (
	"errors"
	"fmt"
	"log"

	// Import interno de packages do próprio sistema
	"microservice_client/pkg/database"
	"microservice_client/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Client (tudo que tiver os métodos abaixo do CRUD são serviços de client)
type ClientServiceInterface interface {
	GetClientsMyGroups(ID *int) (*entity.ClientList, error)
	GetClientByID(ID *uint64) (*entity.Client, error)
	GetClientByReleaseID(ID *uint64) (*entity.ClientList, error)
	GetTagsClient(ID *uint64) ([]*entity.Tag, error)
	CreateClient(client *entity.ClientUpdate, logID *int) error
	UpdateClient(ID *uint64, client *entity.ClientUpdate, logID *int) error
	UpdateStatusClient(ID *uint64, logID *int) error
	InsertTagClient(ID *uint64, tags *[]entity.Tag, logID *int) error
	GetRoles() *entity.RoleList
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

// GetClientsMyGroups: Retorna lista de client pelo group
func (ps *Client_service) GetClientsMyGroups(ID *int) (*entity.ClientList, error) {
	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcGetAllClientsGroup(?)", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list_client := &entity.ClientList{}

	for rows.Next() {
		client := entity.Client{}

		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Role, &client.Customer_Name, &client.Business_Name, &client.Business_ID, &client.Release_Name, &client.Release_ID, &client.User_Name, &client.Status_Description); err != nil {
			return nil, errors.New("error scan client")
		} else {
			rowsTags, err := database.Query("SELECT DISTINCT tT.tag_id, tT.tag_name FROM tblTags tT INNER JOIN tblClientTag tCT ON tT.tag_id = tCT.tag_id WHERE tCT.client_id = ? ORDER BY tT.tag_name", client.ID)
			if err != nil {
				return nil, errors.New("error get tag")
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
					return nil, errors.New("error scan tag")
				} else {
					tags = append(tags, tag)
				}
			}

			client.Tags = tags

			list_client.List = append(list_client.List, &client)
		}
	}

	return list_client, nil

}

// GetClientByID: Retorna um client pelo ID
func (ps *Client_service) GetClientByID(ID *uint64) (*entity.Client, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("call pcGetClientByID(?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var client entity.Client

	err = stmt.QueryRow(ID).Scan(&client.ID, &client.Name, &client.Email, &client.Role, &client.Customer_Name, &client.Business_Name, &client.Release_ID, &client.Release_Name, &client.User_Name, &client.Status_Description)
	if err != nil {
		return nil, errors.New("client not found")
	}

	rowsTags, err := database.Query("SELECT DISTINCT tT.tag_id, tT.tag_name FROM tblTags tT INNER JOIN tblClientTag tCT ON tT.tag_id = tCT.tag_id WHERE tCT.client_id = ? ORDER BY tT.tag_name", ID)
	if err != nil {
		return nil, errors.New("error get tags")
	}

	defer rowsTags.Close()

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			return nil, errors.New("error scan tag")
		} else {
			tags = append(tags, tag)
		}
	}

	client.Tags = tags

	return &client, nil
}

// GetClientByReleaseID: Retorna uma lista de clients pelo ID da release
func (ps *Client_service) GetClientByReleaseID(ID *uint64) (*entity.ClientList, error) {
	database := ps.dbp.GetDB()

	rows, err := database.Query("SELECT tC.client_id, tC.client_name, tC.client_email FROM tblClient tC WHERE tC.release_id = ?", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list_client := &entity.ClientList{}

	for rows.Next() {
		client := entity.Client{}

		if err := rows.Scan(&client.ID, &client.Name, &client.Email); err != nil {
			return nil, errors.New("error scan client")
		}

		list_client.List = append(list_client.List, &client)
	}

	return list_client, nil
}

// GetTagsClient: Retorna uma lista de tag pelo ID do client
func (ps *Client_service) GetTagsClient(ID *uint64) ([]*entity.Tag, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("select DISTINCT T.tag_id, T.tag_name from tblTags T inner join tblClientTag TCT on T.tag_id = TCT.tag_id WHERE client_id = ? ORDER BY T.tag_name")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var tags []*entity.Tag

	rowsTags, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}

	hasResult := false

	for rowsTags.Next() {
		hasResult = true

		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			return nil, errors.New("error scan tag")
		}

		tags = append(tags, &tag)
	}

	if !hasResult {
		return nil, errors.New("tags not found")
	}

	return tags, nil

}

// CreateClient: Cria um novo client
func (ps *Client_service) CreateClient(client *entity.ClientUpdate, logID *int) error {
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

	err = status.QueryRow("CLIENT", "ATIVO").Scan(&statusID)
	if err != nil {
		return errors.New("status not found")
	}

	stmt, err := database.Prepare("INSERT INTO tblClient(client_name, client_email, client_role, customer_id, release_id, business_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(&client.Name, &client.Email, &client.Role, &client.Customer_ID, &client.Release_ID, &client.Business_ID, &client.User_ID, statusID)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	newID := uint64(ID)

	if client.Tags != nil {
		err := ps.InsertTagClient(&newID, &client.Tags, logID)
		if err != nil {
			return errors.New("could not insert tag in clients")
		}
	}

	return nil
}

// UpdateClient: Atualiza as informações do client
func (ps *Client_service) UpdateClient(ID *uint64, client *entity.ClientUpdate, logID *int) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblClient SET client_name = ?, client_email = ?, client_role = ?, customer_id = ?, business_id = ?, user_id = ? WHERE client_id = ?")
	if err != nil {
		return err
	}

	// Definir a variável de sessão "@user"
	_, err = database.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	defer stmt.Close()

	result, err := stmt.Exec(client.Name, client.Email, client.Role, client.Customer_ID, client.Business_ID, client.User_ID, ID)
	if err != nil {
		return errors.New("unable to update client")
	}

	// aqui não esta sendo usado
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	if client.Tags != nil {
		err := ps.InsertTagClient(ID, &client.Tags, logID)
		if err != nil {
			return errors.New("could not insert tag in clients")
		}
	}

	return nil
}

// UpdateStatusClient: Atualizar o status do client
func (ps *Client_service) UpdateStatusClient(ID *uint64, logID *int) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblClient WHERE client_id = ?")
	if err != nil {
		return err
	}

	// Definir a variável de sessão "@user"
	_, err = database.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	defer stmt.Close()

	var statusClient uint64

	err = stmt.QueryRow(ID).Scan(&statusClient)
	if err != nil {
		return errors.New("status client not found")
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return err
	}

	var statusID uint64

	err = status.QueryRow("CLIENT", "Active").Scan(&statusID)
	if err != nil {
		return errors.New("status not found")
	}

	if statusID == statusClient {
		statusClient++
	} else {
		statusClient--
	}

	updt, err := database.Prepare("UPDATE tblClient SET status_id = ? WHERE client_id = ?")
	if err != nil {
		return err
	}

	result, err := updt.Exec(statusClient, ID)
	if err != nil {
		return errors.New("unable to update client")
	}

	// aqui não esta sendo usado
	_, err = result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

// InsertTagClient: Função auxiliar para adicionar tag ao client
func (ps *Client_service) InsertTagClient(ID *uint64, tags *[]entity.Tag, logID *int) error {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM tblClientTag WHERE client_id = ?")
	if err != nil {
		return errors.New("error prepare delete tags on client train")
	}

	// Definir a variável de sessão "@user"
	_, err = database.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return errors.New("error exec statement exec on client train")
	}

	stmt, err = database.Prepare("INSERT IGNORE tblClientTag SET tag_id = ?, client_id = ?")
	if err != nil {
		return errors.New("error insert a new row on tag_id and client")
	}

	defer stmt.Close()

	for _, tag := range *tags {
		_, err := stmt.Exec(tag.Tag_ID, ID)
		if err != nil {
			return errors.New("error insert data tag_ID and ID on database")
		}
	}

	return nil
}

// GetRoles traz todos os Roles do banco de dados
func (ps *Client_service) GetRoles() *entity.RoleList {

	database := ps.dbp.GetDB()

	rows, err := database.Query("SELECT domain_id, domain_value FROM tblDomain where domain_name = 'ROLE'")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_Role := &entity.RoleList{}

	for rows.Next() {
		// variável do tipo Tag(vazia)
		role := entity.Role{}

		// pega dados da query e atribui a variável tag, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&role.Role_ID, &role.Role_Name); err != nil {
			log.Println(err.Error())
		} else {
			list_Role.List = append(list_Role.List, &role)
		}
	}

	return list_Role
}
