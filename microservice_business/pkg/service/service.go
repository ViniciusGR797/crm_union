package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"microservice_business/pkg/database"
	"microservice_business/pkg/entity"
)

// BusinessServiceInterface estrutura de dados para BusinessServiceInterface
type BusinessServiceInterface interface {
	// Pega todos os Businesss, logo lista todos os Businesss
	GetBusiness(ctx context.Context) *entity.BusinessList
	GetBusinessById(ID uint64, ctx context.Context) (*entity.Business, error)
	GetTagsBusiness(ID *uint64, ctx context.Context) ([]*entity.Tag, error)
	CreateBusiness(business *entity.Business_Update, logID *int, ctx context.Context) error
	UpdateBusiness(ID uint64, business *entity.Business_Update, logID *int, ctx context.Context) (uint64, error)
	UpdateStatusBusiness(ID *uint64, logID *int, ctx context.Context) int64
	GetBusinessByName(name *string, ctx context.Context) (*entity.BusinessList, error)
}

// Business_service Estrutura de dados para armazenar a pool de conexão do Database, onde oferece os serviços de CRUD
type Business_service struct {
	dbp database.DatabaseInterface
}

// NewBusinessService Cria um novo serviço de CRUD para pool de conexão
func NewBusinessService(dabase_pool database.DatabaseInterface) *Business_service {
	return &Business_service{
		dabase_pool,
	}
}

// GetBusiness traz todos os Business do banco de dados
func (ps *Business_service) GetBusiness(ctx context.Context) *entity.BusinessList {

	database := ps.dbp.GetDB()
	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT DISTINCT b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description FROM tblBusiness b INNER JOIN tblDomain d on b.segment_id = d.domain_id INNER JOIN  tblStatus s on b.status_id = s.status_id ORDER BY b.business_name")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_Business := &entity.BusinessList{}

	for rows.Next() {
		business := entity.Business{}

		if err := rows.Scan(&business.Business_id, &business.Business_code, &business.Business_name, &business.BusinessSegment.BusinessSegment_id, &business.BusinessSegment.BusinessSegment_description, &business.Status.Status_id, &business.Status.Status_description); err != nil {
			return &entity.BusinessList{}
		} else {
			rowsTags, err := database.QueryContext(ctx, "SELECT DISTINCT tag.tag_id ,tag.tag_name  from tblTags tag  inner join tblBusinessTag tRTT on tag.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag.tag_name", business.Business_id)
			if err != nil {
				return &entity.BusinessList{}
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
					return &entity.BusinessList{}
				} else {
					tags = append(tags, tag)
				}
			}

			business.Tags = tags

			list_Business.List = append(list_Business.List, &business)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil
	}

	return list_Business

}

// GetBusinessById traz um usuario no banco de dados pelo ID do mesmo
func (ps *Business_service) GetBusinessById(ID uint64, ctx context.Context) (*entity.Business, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description FROM tblBusiness b INNER JOIN tblDomain d on b.segment_id = d.domain_id INNER JOIN  tblStatus s on b.status_id = s.status_id WHERE b.business_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	Business := &entity.Business{}

	err = stmt.QueryRow(ID).Scan(&Business.Business_id, &Business.Business_code, &Business.Business_name, &Business.BusinessSegment.BusinessSegment_id, &Business.BusinessSegment.BusinessSegment_description, &Business.Status.Status_id, &Business.Status.Status_description)
	if err != nil {
		return &entity.Business{}, errors.New("error scanning rows")
	}

	rowsTags, err := tx.Query("SELECT DISTINCT tag.tag_id, tag.tag_name from tblTags tag INNER JOIN tblBusinessTag tRTT on tag.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag.tag_name", Business.Business_id)
	if err != nil {
		return &entity.Business{}, errors.New("error fetching tags from business by id")
	}
	defer rowsTags.Close()

	var tags []entity.Tag

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			return &entity.Business{}, errors.New("error scanning tags from business by id")
		} else {
			tags = append(tags, tag)
		}
	}

	Business.Tags = tags

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return Business, nil
}

// CreateBusiness cria um Business no banco de dados
func (ps *Business_service) CreateBusiness(business *entity.Business_Update, logID *int, ctx context.Context) error {

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

	result, err := tx.ExecContext(ctx, "INSERT INTO tblBusiness (business_code, business_name, segment_id, status_id) VALUES (?, ?, ?, ?)", business.Code, business.Name, business.Segment_Id, 1)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = result.RowsAffected()
	if err != nil {
		return errors.New("error rowAffected insert into database")
	}

	ID, _ := result.LastInsertId()
	business.ID = uint64(ID)

	if business.Tags != nil {

		for _, tag := range business.Tags {
			_, err := tx.ExecContext(ctx, "INSERT tblBusinessTag SET tag_id = ?, business_id = ?", tag.Tag_ID, business.ID)
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

// UpdateBusiness atualiza os dados de um Bussines no banco pelo ID do mesmo
func (ps *Business_service) UpdateBusiness(ID uint64, business *entity.Business_Update, logID *int, ctx context.Context) (uint64, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, err
	}

	result, err := tx.ExecContext(ctx, "UPDATE tblBusiness SET business_name = ?, business_code = ?, segment_id = ? WHERE business_id = ?", business.Name, business.Code, business.Segment_Id, ID)
	if err != nil {
		return 0, errors.New("error prepare update business")
	}

	_, err = result.RowsAffected()
	if err != nil {
		return 0, errors.New("error RowsAffected update business")
	}

	if business.Tags != nil {
		_, err = tx.ExecContext(ctx, "DELETE FROM tblBusinessTag WHERE business_id = ?", ID)
		if err != nil {
			return 0, errors.New("error prepare delete tags on client train")
		}

		for _, tag := range business.Tags {
			_, err := tx.ExecContext(ctx, "INSERT IGNORE tblBusinessTag SET tag_id = ?, business_id = ?", tag.Tag_ID, ID)
			if err != nil {
				return 0, errors.New("error insert data tag_ID and ID on database")
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// UpdateStatusBusiness altera o status de um usuario no banco
func (ps *Business_service) UpdateStatusBusiness(ID *uint64, logID *int, ctx context.Context) int64 {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT status_id FROM tblBusiness WHERE business_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0
	}

	var statusBusiness uint64

	err = stmt.QueryRow(ID).Scan(&statusBusiness)
	if err != nil {
		log.Println(err.Error())
	}

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer status.Close()

	var statusID uint64

	err = status.QueryRow("Business", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	if statusID == statusBusiness {
		statusBusiness++
	} else {
		statusBusiness--
	}

	updt, err := tx.ExecContext(ctx, "UPDATE tblBusiness SET status_id = ? WHERE business_id = ?", statusBusiness, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := updt.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return 0
	}

	return rowsaff
}

// GetBusinessByName busca Business no banco de dados pelo nome passado como parâmetro no query.
func (ps *Business_service) GetBusinessByName(name *string, ctx context.Context) (*entity.BusinessList, error) {
	nameString := fmt.Sprint("%", *name, "%")
	query := "SELECT DISTINCT b.business_id, b.business_code, b.business_name, b.segment_id, d.domain_value, b.status_id, s.status_description FROM tblBusiness b inner join tblDomain d on b.segment_id = d.domain_id inner join  tblStatus s on b.status_id = s.status_id WHERE b.business_name LIKE ? ORDER BY b.business_name"

	// pega database
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// manda uma query para ser executada no database
	rows, err := tx.Query(query, nameString)
	// verifica se teve erro
	if err != nil {
		log.Println(err.Error())
		return &entity.BusinessList{}, errors.New("error fetching Businesss")
	}

	// fecha linha da query, quando sair da função
	defer rows.Close()

	// variável do tipo BusinessList (vazia)
	list_Business := &entity.BusinessList{}

	for rows.Next() {
		business := entity.Business{}

		if err := rows.Scan(&business.Business_id, &business.Business_code, &business.Business_name, &business.BusinessSegment.BusinessSegment_id, &business.BusinessSegment.BusinessSegment_description, &business.Status.Status_id, &business.Status.Status_description); err != nil {
			return &entity.BusinessList{}, nil
		} else {

			rowsTags, err := database.QueryContext(ctx, "SELECT DISTINCT tag.tag_id, tag.tag_name from tblTags tag inner join tblBusinessTag tRTT on tag.tag_id = tRTT.tag_id WHERE tRTT.business_id = ? ORDER BY tag.tag_name", business.Business_id)
			if err != nil {
				return &entity.BusinessList{}, nil
			}

			var tags []entity.Tag

			for rowsTags.Next() {
				tag := entity.Tag{}

				if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
					return &entity.BusinessList{}, nil
				} else {
					tags = append(tags, tag)
				}
			}

			business.Tags = tags

			list_Business.List = append(list_Business.List, &business)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return list_Business, nil
}

// GetTagsBusiness busca as tags de Business
func (ps *Business_service) GetTagsBusiness(ID *uint64, ctx context.Context) ([]*entity.Tag, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT DISTINCT T.tag_id, T.tag_name from tblTags T INNER JOIN tblBusinessTag B on T.tag_id = B.tag_id WHERE business_id = ? ORDER BY T.tag_name")
	if err != nil {
		return []*entity.Tag{}, errors.New("error fetching on tag business")
	}

	defer stmt.Close()

	var tags []*entity.Tag

	rowsTags, err := stmt.Query(ID)
	if err != nil {
		return []*entity.Tag{}, errors.New("error fetching on row tags query release train")
	}
	defer rowsTags.Close()

	for rowsTags.Next() {
		tag := entity.Tag{}

		if err := rowsTags.Scan(&tag.Tag_ID, &tag.Tag_Name); err != nil {
			return []*entity.Tag{}, errors.New("error fetching on row tags next release train")
		}

		tags = append(tags, &tag)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tags, nil
}
