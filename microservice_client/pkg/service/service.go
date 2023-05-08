package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	// Import interno de packages do próprio sistema
	"microservice_client/pkg/database"
	"microservice_client/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Client (tudo que tiver os métodos abaixo do CRUD são serviços de client)
type ClientServiceInterface interface {
	GetAllClients(ID *int, ctx context.Context) (*entity.ClientList, error)
	GetClientsMyGroups(ID *int, ctx context.Context) (*entity.ClientList, error)
	GetClientByID(ID *uint64, ctx context.Context) (*entity.Client, error)
	GetClientByReleaseID(ID *uint64, ctx context.Context) (*entity.ClientList, error)
	GetTagsClient(ID *uint64, ctx context.Context) ([]*entity.Tag, error)
	CreateClient(client *entity.ClientUpdate, logID *int, ctx context.Context) error
	UpdateClient(ID *uint64, client *entity.ClientUpdate, logID *int, ctx context.Context) error
	UpdateStatusClient(ID *uint64, logID *int, ctx context.Context) error
	GetRoles(ctx context.Context) *entity.RoleList
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

func (ps *Client_service) GetAllClients(ID *int, ctx context.Context) (*entity.ClientList, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list_client := &entity.ClientList{}

	for rows.Next() {
		client := entity.Client{}

		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Role_ID, &client.Role, &client.Customer_ID, &client.Customer_Name, &client.Business_Name, &client.Business_ID, &client.Release_Name, &client.Release_ID, &client.User_ID, &client.User_Name, &client.Status_Description); err != nil {
			return nil, errors.New("error scan client")
		} else {
			rowsTags, err := database.QueryContext(ctx, "SELECT DISTINCT tT.tag_id, tT.tag_name FROM tblTags tT INNER JOIN tblClientTag tCT ON tT.tag_id = tCT.tag_id WHERE tCT.client_id = ? ORDER BY tT.tag_name", client.ID)
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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return list_client, nil
}

// GetClientsMyGroups: Retorna lista de client pelo group
func (ps *Client_service) GetClientsMyGroups(ID *int, ctx context.Context) (*entity.ClientList, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("call pcGetAllClientsGroup(?)", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list_client := &entity.ClientList{}

	for rows.Next() {
		client := entity.Client{}

		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Role_ID, &client.Role, &client.Customer_ID, &client.Customer_Name, &client.Business_Name, &client.Business_ID, &client.Release_Name, &client.Release_ID, &client.User_ID, &client.User_Name, &client.Status_Description); err != nil {
			return nil, errors.New("error scan client")
		} else {
			rowsTags, err := database.QueryContext(ctx, "SELECT DISTINCT tT.tag_id, tT.tag_name FROM tblTags tT INNER JOIN tblClientTag tCT ON tT.tag_id = tCT.tag_id WHERE tCT.client_id = ? ORDER BY tT.tag_name", client.ID)
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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return list_client, nil

}

// GetClientByID: Retorna um client pelo ID
func (ps *Client_service) GetClientByID(ID *uint64, ctx context.Context) (*entity.Client, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("call pcGetClientByID(?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var client entity.Client

	err = stmt.QueryRow(ID).Scan(&client.ID, &client.Name, &client.Email, &client.Role, &client.Customer_Name, &client.Business_Name, &client.Release_ID, &client.Release_Name, &client.User_Name, &client.Status_Description)
	if err != nil {
		return nil, errors.New("client not found")
	}

	rowsTags, err := tx.Query("SELECT DISTINCT tT.tag_id, tT.tag_name FROM tblTags tT INNER JOIN tblClientTag tCT ON tT.tag_id = tCT.tag_id WHERE tCT.client_id = ? ORDER BY tT.tag_name", ID)
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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// GetClientByReleaseID: Retorna uma lista de clients pelo ID da release
func (ps *Client_service) GetClientByReleaseID(ID *uint64, ctx context.Context) (*entity.ClientList, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT tC.client_id, tC.client_name, tC.client_email FROM tblClient tC WHERE tC.release_id = ?", ID)
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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return list_client, nil
}

// GetTagsClient: Retorna uma lista de tag pelo ID do client
func (ps *Client_service) GetTagsClient(ID *uint64, ctx context.Context) ([]*entity.Tag, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("select DISTINCT T.tag_id, T.tag_name from tblTags T inner join tblClientTag TCT on T.tag_id = TCT.tag_id WHERE client_id = ? ORDER BY T.tag_name")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var tags []*entity.Tag

	rowsTags, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}
	defer rowsTags.Close()

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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tags, nil

}

// CreateClient: Cria um novo client
func (ps *Client_service) CreateClient(client *entity.ClientUpdate, logID *int, ctx context.Context) error {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return err
	}
	defer status.Close()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	var statusID uint64

	err = status.QueryRow("CLIENT", "Active").Scan(&statusID)
	if err != nil {
		return errors.New("status not found")
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO tblClient(client_name, client_email, client_role, customer_id, release_id, business_id, user_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", &client.Name, &client.Email, &client.Role, &client.Customer_ID, &client.Release_ID, &client.Business_ID, &client.User_ID, statusID)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if client.Tags != nil {
		_, err = tx.ExecContext(ctx, "DELETE FROM tblClientTag WHERE client_id = ?", ID)
		if err != nil {
			return errors.New("error prepare delete tags on client train")
		}

		for _, tag := range client.Tags {
			_, err := tx.ExecContext(ctx, "INSERT IGNORE tblClientTag SET tag_id = ?, client_id = ?", tag.Tag_ID, ID)
			if err != nil {
				return errors.New("error insert data tag_ID and ID on database")
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// UpdateClient: Atualiza as informações do client
func (ps *Client_service) UpdateClient(ID *uint64, client *entity.ClientUpdate, logID *int, ctx context.Context) error {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	result, err := tx.ExecContext(ctx, "UPDATE tblClient SET client_name = ?, client_email = ?, client_role = ?, customer_id = ?, business_id = ?, user_id = ? WHERE client_id = ?", client.Name, client.Email, client.Role, client.Customer_ID, client.Business_ID, client.User_ID, ID)
	if err != nil {
		return err
	}

	// aqui não esta sendo usado
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	if client.Tags != nil {
		_, err = tx.ExecContext(ctx, "DELETE FROM tblClientTag WHERE client_id = ?", ID)
		if err != nil {
			return errors.New("error prepare delete tags on client train")
		}

		for _, tag := range client.Tags {
			_, err := tx.ExecContext(ctx, "INSERT IGNORE tblClientTag SET tag_id = ?, client_id = ?", tag.Tag_ID, ID)
			if err != nil {
				return errors.New("error insert data tag_ID and ID on database")
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// UpdateStatusClient: Atualizar o status do client
func (ps *Client_service) UpdateStatusClient(ID *uint64, logID *int, ctx context.Context) error {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT status_id FROM tblClient WHERE client_id = ?")
	if err != nil {
		return err
	}

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	defer stmt.Close()

	var statusClient uint64

	err = stmt.QueryRow(ID).Scan(&statusClient)
	if err != nil {
		return errors.New("status client not found")
	}

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return err
	}
	defer status.Close()

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

	result, err := tx.ExecContext(ctx, "UPDATE tblClient SET status_id = ? WHERE client_id = ?", statusClient, ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	// aqui não esta sendo usado
	_, err = result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

// GetRoles traz todos os Roles do banco de dados
func (ps *Client_service) GetRoles(ctx context.Context) *entity.RoleList {

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
