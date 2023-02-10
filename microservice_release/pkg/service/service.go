package service

import (

	// Import interno de packages do próprio sistema
	"fmt"
	"microservice_release/pkg/database"
	"microservice_release/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Release (tudo que tiver os métodos abaixo do CRUD são serviços de release)
type ReleaseServiceInterface interface {
	GetReleasesTrain() *entity.ReleaseList
}

// Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type Release_service struct {
	dbp database.DatabaseInterface
}

// Cria novo serviço de CRUD para pool de conexão
func NewReleaseService(dabase_pool database.DatabaseInterface) *Release_service {
	return &Release_service{
		dabase_pool,
	}
}

// Função que retorna lista de client
func (ps *Release_service) GetReleasesTrain() *entity.ReleaseList {
	database := ps.dbp.GetDB()

	rows, err := database.Query("select * from vwGetAllReleaseTrains")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_release := &entity.ReleaseList{}

	for rows.Next() {
		release := entity.Release{}

		if err := rows.Scan(&release.ID, &release.Code, &release.Name, &release.Business_Name, &release.Status_Description); err != nil {
			fmt.Println(err.Error())
		} else {
			rowsTags, err := database.Query("select tag_name from tblTags inner join tblReleaseTrainTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.release_id = ?", release.ID)
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

			release.Tags = tags

			list_release.List = append(list_release.List, &release)
		}
	}

	return list_release
}
