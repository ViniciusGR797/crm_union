package service

import (

	// Import interno de packages do próprio sistema
	"errors"
	"fmt"
	"log"
	"microservice_release/pkg/database"
	"microservice_release/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Release (tudo que tiver os métodos abaixo do CRUD são serviços de release)
type ReleaseServiceInterface interface {
	GetReleasesTrain() *entity.ReleaseList
	GetReleaseTrainByID(ID uint64) *entity.Release
	UpdateReleaseTrain(ID uint64, release *entity.Release_Update) uint64
	GetTagsReleaseTrain(ID *uint64) []*entity.Tag
	InsertTagsReleaseTrain(ID uint64, tags []entity.Tag) (uint64, error)
	UpdateStatusReleaseTrain(ID *uint64) (int64, error)
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

func (ps *Release_service) GetReleaseTrainByID(ID uint64) *entity.Release {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("select * from vwGetAllReleaseTrains where release_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	release := &entity.Release{}

	err = stmt.QueryRow(ID).Scan(&release.ID, &release.Code, &release.Name, &release.Business_Name, &release.Status_Description)
	if err != nil {
		log.Println(err.Error())
	}

	rowsTags, err := database.Query("select tag_name from tblTags inner join tblReleaseTrainTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.release_id = ?", ID)
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

	return release

}
func (ps *Release_service) UpdateReleaseTrain(ID uint64, release *entity.Release_Update) uint64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("UPDATE tblReleaseTrain SET release_code = ?, release_name = ?, business_id = ? WHERE release_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	var releaseID int64

	result, err := stmt.Exec(release.Code, release.Name, release.Business_ID, ID)
	if err != nil {
		log.Println(err.Error())
	}

	releaseID, err = result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return uint64(releaseID)
}

func (ps *Release_service) GetTagsReleaseTrain(ID *uint64) []*entity.Tag {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT T.tag_id, T.tag_name from tblTags T INNER JOIN tblReleaseTrainTag tRTT on T.tag_id = tRTT.tag_id WHERE release_id = ?")
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

func (ps *Release_service) InsertTagsReleaseTrain(ID uint64, tags []entity.Tag) (uint64, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM tblReleaseTrainTag WHERE release_id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	stmt, err = database.Prepare("INSERT IGNORE tblReleaseTrainTag SET tag_id = ?, release_id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	defer stmt.Close()

	for _, tag := range tags {
		_, err := stmt.Exec(tag.Tag_ID, ID)
		if err != nil {
			log.Println(err.Error())
			return 0, err
		}
	}

	return ID, nil
}

func (ps *Release_service) UpdateStatusReleaseTrain(ID *uint64) (int64, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblReleaseTrain WHERE release_id = ?")
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error preparing statement")
	}

	var statusID uint64

	err = stmt.QueryRow(ID).Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}

	if statusID == 7 {
		statusID = 8
	} else {
		statusID = 7
	}

	updt, err := database.Prepare("UPDATE tblReleaseTrain SET status_id = ? WHERE release_id = ?")
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error preparing statement")
	}

	defer stmt.Close()

	result, err := updt.Exec(statusID, ID)
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("error fetching rows affected")
	}

	return rowsaff, nil
}
