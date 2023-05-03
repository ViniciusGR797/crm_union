package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"microservice_tags/pkg/database"
	"microservice_tags/pkg/entity"
)

// TagsServiceInterface estrutura de dados para TagsServiceInterface
type TagsServiceInterface interface {
	// Pega todos os Tagss, logo lista todos os Tagss
	GetTags(ctx context.Context) *entity.TagsList
	GetTagsById(ID uint64, ctx context.Context) (*entity.Tags, error)
	GetDomains(ctx context.Context) *entity.DomainList
	GetDomainById(ID uint64, ctx context.Context) (*entity.Domain, error)
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
func (ps *Tags_service) GetTags(ctx context.Context) *entity.TagsList {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT tag_id, tag_name, tag_type FROM tblTags")
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
		if err := rows.Scan(&tag.Tag_ID, &tag.Tag_Name, &tag.Tag_Type); err != nil {
			log.Println(err.Error())
		} else {
			list_Tags.List = append(list_Tags.List, &tag)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil
	}

	return list_Tags
}

func (ps *Tags_service) GetDomains(ctx context.Context) *entity.DomainList {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT domain_id, domain_name, domain_code, domain_value FROM tblDomain")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_Domain := &entity.DomainList{}

	for rows.Next() {
		// variável do tipo Tag(vazia)
		domain := entity.Domain{}

		// pega dados da query e atribui a variável tag, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&domain.Domain_ID, &domain.Domain_Name, &domain.Domain_Code, &domain.Domain_Value); err != nil {
			log.Println(err.Error())
		} else {
			list_Domain.List = append(list_Domain.List, &domain)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil
	}

	return list_Domain
}

func (ps *Tags_service) GetDomainById(ID uint64, ctx context.Context) (*entity.Domain, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT domain_id, domain_name, domain_code, domain_value FROM tblDomain WHERE domain_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	Domain := &entity.Domain{}

	err = stmt.QueryRow(ID).Scan(&Domain.Domain_ID, &Domain.Domain_Name, &Domain.Domain_Code, &Domain.Domain_Value)
	if err != nil {
		return &entity.Domain{}, errors.New("error scanning rows")
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return Domain, nil
}

// GetTagsById traz um usuario no banco de dados pelo ID do mesmo
func (ps *Tags_service) GetTagsById(ID uint64, ctx context.Context) (*entity.Tags, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT tag_id, tag_name, tag_type FROM tblTags WHERE tag_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	Tags := &entity.Tags{}

	err = stmt.QueryRow(ID).Scan(&Tags.Tag_ID, &Tags.Tag_Name, &Tags.Tag_Type)
	if err != nil {
		return &entity.Tags{}, errors.New("error scanning rows")
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return Tags, nil
}
