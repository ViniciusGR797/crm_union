package service

import (
	"errors"
	"fmt"
	"log"
	"microservice_tags/pkg/database"
	"microservice_tags/pkg/entity"
)

// TagsServiceInterface estrutura de dados para TagsServiceInterface
type TagsServiceInterface interface {
	// Pega todos os Tagss, logo lista todos os Tagss
	GetTags() *entity.TagsList
	GetTagsById(ID uint64) (*entity.Tags, error)
}

// Tags_service Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type Tags_service struct {
	dbp database.DatabaseInterface
}

// NewTagsService Cria um novo serviço de CRUD para pool de conexão
func NewTagsService(dabase_pool database.DatabaseInterface) *Tags_service {
	return &Tags_service{
		dabase_pool,
	}
}

// GetTags traz todos os Tags do banco de dados
func (ps *Tags_service) GetTags() *entity.TagsList {

	database := ps.dbp.GetDB()
	defer database.Close()

	rows, err := database.Query("SELECT tag_id, tag_name FROM tblTags")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_Tags := &entity.TagsList{}

	for rows.Next() {
		// variável do tipo Tag(vazia)
		tag := entity.Tags{}

		// pega dados da query e atribui a variável tag, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			log.Println(err.Error())
		} else {
			list_Tags.List = append(list_Tags.List, &tag)
		}
	}

	return list_Tags
}

// GetTagsById traz um usuario no banco de dados pelo ID do mesmo
func (ps *Tags_service) GetTagsById(ID uint64) (*entity.Tags, error) {

	database := ps.dbp.GetDB()
	defer database.Close()

	stmt, err := database.Prepare("SELECT tag_id, tag_name FROM tblTags WHERE tag_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	Tags := &entity.Tags{}

	err = stmt.QueryRow(ID).Scan(&Tags.Tag_ID, &Tags.Tag_Name)
	if err != nil {
		return &entity.Tags{}, errors.New("error scanning rows")
	}

	return Tags, nil
}
