package service

import (
	"fmt"

	// Import interno de packages do próprio sistema
	"microservice_client/pkg/database"
	"microservice_client/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Client (tudo que tiver os métodos abaixo do CRUD são serviços de client)
type ClientServiceInterface interface {
	// Pega todos os clients, logo lista todos os clients
	GetClientsMyGroups(id uint64) *entity.ClientList
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
func (ps *Client_service) GetClientsMyGroups(id uint64) *entity.ClientList {
	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query("call SelectAllClients (?)", id)
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo ClientList (vazia)
	list_client := &entity.ClientList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Client (vazia)
		client := entity.Client{}

		// pega dados da query e atribui a variável user, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&client.Name, &client.Email, &client.Role, &client.Costumer_ID, &client.Business_ID, &client.Relase_ID, &client.Status_ID); err != nil {
			fmt.Println(err.Error())
		} else {
			// caso não tenha erro, adiciona a variável log na lista de logs
			list_client.List = append(list_client.List, &client)
		}

	}

	// retorna lista de client
	return list_client
}
