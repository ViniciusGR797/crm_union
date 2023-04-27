package service

import (

	// Import interno de packages do próprio sistema
	"errors"
	"microservice_release/pkg/database"
	"microservice_release/pkg/entity"
)

// Estrutura interface para padronizar comportamento de CRUD Release (tudo que tiver os métodos abaixo do CRUD são serviços de release)
type ReleaseServiceInterface interface {
	GetReleasesTrain() (*entity.ReleaseList, error)
	GetReleaseTrainByID(ID uint64) (*entity.Release, error)
	UpdateReleaseTrain(ID uint64, release *entity.Release_Update, logID *int) (uint64, error)
	GetTagsReleaseTrain(ID *uint64) ([]*entity.Tag, error)
	InsertTagsReleaseTrain(ID uint64, tags []entity.Tag, logID *int) (uint64, error)
	UpdateStatusReleaseTrain(ID *uint64, logID *int) (int64, error)
	GetReleaseTrainByBusiness(businessID *uint64) (*entity.ReleaseList, error)
	CreateReleaseTrain(release *entity.Release_Update, logID *int) error
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

// Função que retorna lista de release train
func (ps *Release_service) GetReleasesTrain() (*entity.ReleaseList, error) {
	database := ps.dbp.GetDB()

	rows, err := database.Query("select DISTINCT * from vwGetAllReleaseTrains ORDER BY release_name")
	if err != nil {
		return &entity.ReleaseList{}, errors.New("error fetching release train")
	}

	defer rows.Close()

	list_release := &entity.ReleaseList{}
	hasResult := false
	for rows.Next() {
		hasResult = true
		release := entity.Release{}

		if err := rows.Scan(&release.ID, &release.Code, &release.Business_Name, &release.Business_Id, &release.Name, &release.Status_Description); err != nil {
			return &entity.ReleaseList{}, errors.New("error scanning release train")
		} else {
			rowsTags, err := database.Query("select DISTINCT tblTags.tag_id, tag_name from tblTags inner join tblReleaseTrainTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.release_id = ? ORDER BY tag_name ", release.ID)
			if err != nil {
				return &entity.ReleaseList{}, errors.New("error fetching tags")
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
					return &entity.ReleaseList{}, errors.New("error scanning tags")
				} else {
					tags = append(tags, tag)
				}
			}
			defer rowsTags.Close()

			release.Tags = tags

			list_release.List = append(list_release.List, &release)
		}
	}
	if !hasResult {
		return nil, errors.New("release not found")
	}

	return list_release, nil
}

// GetReleaseTrainByID busca release train por ID
func (ps *Release_service) GetReleaseTrainByID(ID uint64) (*entity.Release, error) {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("select * from vwGetAllReleaseTrains where release_id = ?")
	if err != nil {
		return &entity.Release{}, errors.New("error prepare fetching release train by id")
	}
	defer stmt.Close()

	release := &entity.Release{}

	err = stmt.QueryRow(ID).Scan(&release.ID, &release.Code, &release.Business_Name, &release.Business_Id, &release.Name, &release.Status_Description)
	if err != nil {
		return &entity.Release{}, errors.New("error scanning rows")
	}

	rowsTags, err := database.Query("select DISTINCT tblTags.tag_id, tag_name from tblTags inner join tblReleaseTrainTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.release_id = ? ORDER BY tag_name", ID)
	if err != nil {
		return &entity.Release{}, errors.New("error fetching tags from release train by id")
	}

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			return &entity.Release{}, errors.New("error scanning tags from release train by id")
		} else {
			tags = append(tags, tag)
		}
	}
	defer rowsTags.Close()

	release.Tags = tags

	return release, nil
}

// UpdateReleaseTrain atualiza a release train
func (ps *Release_service) UpdateReleaseTrain(ID uint64, release *entity.Release_Update, logID *int) (uint64, error) {
	database := ps.dbp.GetDB()

	// Definir a variável de sessão "@user"
	_, err := database.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	stmt, err := database.Prepare("UPDATE tblReleaseTrain SET release_code = ?, release_name = ?, business_id = ? WHERE release_id = ?")
	if err != nil {
		return 0, errors.New("error prepare update release train")
	}

	defer stmt.Close()

	var releaseID int64

	result, err := stmt.Exec(release.Code, release.Name, release.Business_ID, ID)
	if err != nil {
		return 0, errors.New("error exec update release train")
	}

	releaseID, err = result.RowsAffected()
	if err != nil {
		return 0, errors.New("error RowsAffected update release train")
	}

	return uint64(releaseID), nil

}

// GetTagsReleaseTrain busca as tags da release train
func (ps *Release_service) GetTagsReleaseTrain(ID *uint64) ([]*entity.Tag, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT DISTINCT T.tag_id, T.tag_name from tblTags T INNER JOIN tblReleaseTrainTag tRTT on T.tag_id = tRTT.tag_id WHERE release_id = ? ORDER BY T.tag_name")
	if err != nil {
		return []*entity.Tag{}, errors.New("error fetching on tag release train")
	}

	defer stmt.Close()

	var tags []*entity.Tag

	rowsTags, err := stmt.Query(ID)
	if err != nil {
		return []*entity.Tag{}, errors.New("error fetching on row tags query release train")
	}

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			return []*entity.Tag{}, errors.New("error fetching on row tags next release train")
		}

		tags = append(tags, &tag)
	}
	defer rowsTags.Close()

	return tags, nil
}

// InsertTagsReleaseTrain deleta relese train tag e dps insere novamente as alterações
func (ps *Release_service) InsertTagsReleaseTrain(ID uint64, tags []entity.Tag, logID *int) (uint64, error) {
	database := ps.dbp.GetDB()

	// Definir a variável de sessão "@user"
	_, err := database.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	stmt, err := database.Prepare("DELETE FROM tblReleaseTrainTag WHERE release_id = ?")
	if err != nil {
		return 0, errors.New("error prepare delete tags on release train")
	}

	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return 0, errors.New("error exec statement exec on release train")
	}

	stmt, err = database.Prepare("INSERT IGNORE tblReleaseTrainTag SET tag_id = ?, release_id = ?")
	if err != nil {
		return 0, errors.New("error insert a new row on tag_id and release_id")
	}

	defer stmt.Close()

	for _, tag := range tags {
		_, err := stmt.Exec(tag.Tag_ID, ID)
		if err != nil {
			return 0, errors.New("error insert data tag_ID and ID on database")
		}
	}

	return ID, nil
}

// UpdateStatusReleaseTrain atualiza o status da release train "softdelete"
func (ps *Release_service) UpdateStatusReleaseTrain(ID *uint64, logID *int) (int64, error) {
	database := ps.dbp.GetDB()

	// Definir a variável de sessão "@user"
	_, err := database.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	stmt, err := database.Prepare("SELECT status_id FROM tblReleaseTrain WHERE release_id = ?")
	if err != nil {
		return 0, errors.New("error preparing statement")
	}

	var statusID uint64

	err = stmt.QueryRow(ID).Scan(&statusID)
	if err != nil {
		return 0, errors.New("error preparing statement QueryRow")
	}

	if statusID == 7 {
		statusID = 8
	} else {
		statusID = 7
	}

	updt, err := database.Prepare("UPDATE tblReleaseTrain SET status_id = ? WHERE release_id = ?")
	if err != nil {
		return 0, errors.New("error preparing update status_id in release_id on database")
	}
	defer updt.Close()
	defer stmt.Close()

	result, err := updt.Exec(statusID, ID)
	if err != nil {
		return 0, errors.New("error preparing update on statusID and ID")
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("error fetching rows affected")
	}

	return rowsaff, nil
}

// Função que retorna lista de releases, filtrando pelo ID business
func (ps *Release_service) GetReleaseTrainByBusiness(businessID *uint64) (*entity.ReleaseList, error) {
	query := "SELECT DISTINCT V.release_id, V.release_code, V.release_name, V.business_name, V.business_id, V.status_description FROM vwGetAllReleaseTrains V INNER JOIN tblReleaseTrain R ON V.release_id = R.release_id WHERE R.business_id = ? ORDER BY V.release_name"

	// pega database
	database := ps.dbp.GetDB()

	// manda uma query para ser executada no database
	rows, err := database.Query(query, businessID)
	// verifica se teve erro
	if err != nil {
		return &entity.ReleaseList{}, errors.New("error fetching release's business")
	}
	defer rows.Close()

	// variável do tipo ReleaseList (vazia)
	releaseList := &entity.ReleaseList{}

	// Pega todo resultado da query linha por linha
	for rows.Next() {
		// variável do tipo Release (vazia)
		release := entity.Release{}

		// pega dados da query e atribui ao release, além de verificar se teve erro ao pegar dados
		if err := rows.Scan(&release.ID,
			&release.Code,
			&release.Name,
			&release.Business_Name,
			&release.Business_Id,
			&release.Status_Description); err != nil {
			return &entity.ReleaseList{}, errors.New("error scanning release data")
		} else {
			// caso não tenha erro, adiciona a lista de users
			releaseList.List = append(releaseList.List, &release)
		}
	}

	// For para pegar tags da lista de releases
	for _, release := range releaseList.List {
		query := "SELECT DISTINCT tag_name FROM tblTags INNER JOIN tblReleaseTrainTag tRTT on tblTags.tag_id = tRTT.tag_id WHERE tRTT.release_id = ? ORDER BY tag_name"

		// manda uma query para ser executada no database
		rows, err := database.Query(query, release.ID)
		// verifica se teve erro
		if err != nil {
			return &entity.ReleaseList{}, errors.New("error fetching tags")
		}

		// Variável do tipo slice de Tag
		var tags []entity.Tag

		// Pega todo resultado da query linha por linha
		for rows.Next() {
			// variável do tipo User (vazia)
			tag := entity.Tag{}

			// pega dados da query e atribui a variável groupID, além de verificar se teve erro ao pegar dados
			if err := rows.Scan(&tag.Tag_Name); err != nil {
				return &entity.ReleaseList{}, errors.New("error scanning tag data")
			} else {
				// caso não tenha erro, adiciona a lista de users
				tags = append(tags, tag)
			}
		}

		// Atribui slice de tag para um release
		release.Tags = tags
	}

	// Retornar lista de releases
	return releaseList, nil
}

// CreateReleaseTrain cria release train
func (ps *Release_service) CreateReleaseTrain(release *entity.Release_Update, logID *int) error {
	database := ps.dbp.GetDB()

	// Definir a variável de sessão "@user"
	_, err := database.Exec("SET @user = ?", logID)
	if err != nil {
		return errors.New("session variable error")
	}

	stmt, err := database.Prepare("INSERT INTO tblReleaseTrain (release_code, release_name, business_id, status_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return errors.New("error in prepare release statement")
	}
	defer stmt.Close()

	result, err := stmt.Exec(release.Code, release.Name, release.Business_ID, 7)
	if err != nil {
		return errors.New("release exec error")
	}

	_, err = result.RowsAffected()
	if err != nil {
		return errors.New("error rowAffected insert into database")
	}

	ID, _ := result.LastInsertId()
	release.ID = uint64(ID)

	stmt, err = database.Prepare("INSERT INTO tblReleaseTrainTag (tag_id, release_id) VALUES (?, ?)")
	if err != nil {
		return errors.New("error in prepare release tags statement")
	}

	for _, tag := range release.Tags {
		_, err := stmt.Exec(tag.Tag_ID, release.ID)
		if err != nil {
			return errors.New("release tags exec error")
		}
	}

	return nil
}
