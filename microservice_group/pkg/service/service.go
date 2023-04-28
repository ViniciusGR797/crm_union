package service

import (
	"context"
	"errors"
	"fmt"
	"microservice_group/pkg/database"
	"microservice_group/pkg/entity"
)

type GroupServiceInterface interface {
	GetGroups(id uint64) (*entity.GroupList, error)
	GetGroupByID(id uint64) (*entity.GroupID, error)
	UpdateStatusGroup(id uint64, logID *int, ctx context.Context) (int64, error)
	GetUsersGroup(id uint64) (*entity.UserList, error)
	CreateGroup(group *entity.CreateGroup, logID *int, ctx context.Context) (int64, error)
	AttachUserGroup(users *entity.GroupIDList, id uint64, logID *int, ctx context.Context) (int64, error)
	DetachUserGroup(users *entity.GroupIDList, id uint64, logID *int, ctx context.Context) (int64, error)
	CountUsersGroup(id uint64) (*entity.CountUsersList, error)
	EditGroup(group *entity.EditGroup, id uint64, logID *int, ctx context.Context) (int64, error)
}

type Group_service struct {
	dbp database.DatabaseInterface
}

// InsertUserList implements GroupServiceInterface

// GetGroup implements GroupServiceInterface

func NewGroupService(dabase_pool database.DatabaseInterface) *Group_service {
	return &Group_service{
		dabase_pool,
	}
}

// GetGroups retorna todos os grupos do usuario
func (ps *Group_service) GetGroups(id uint64) (*entity.GroupList, error) {

	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcGetAllGroupUserNoId (?)", id)
	// verifica se teve erro
	if err != nil {
		return nil, err
	}

	hasResult := false

	defer rows.Close()

	list_groups := &entity.GroupList{}

	for rows.Next() {

		hasResult = true

		group := entity.Group{}

		if err := rows.Scan(
			&group.Group_id,
			&group.Group_name,
			&group.Status.Status_id,
			&group.Status.Status_description,
			&group.Created_at,
			&group.Customer.Customer_id,
			&group.Customer.Customer_name,
		); err != nil {
			fmt.Println(err.Error())
		} else {
			rows2, err := database.Query("call pcGetAllUserGroup (?)", group.Group_id)
			if err != nil {
				return nil, err
			}
			var user_list []entity.User

			for rows2.Next() {
				user := entity.User{}

				if err := rows2.Scan(
					&user.User_id,
					&user.User_name); err != nil {
					fmt.Println(err.Error())
				} else {
					user_list = append(user_list, user)
				}

			}

			group.Users = user_list

			list_groups.List = append(list_groups.List, &group)

		}

	}

	if !hasResult {
		return nil, fmt.Errorf("no groups found")
	}

	return list_groups, nil

}

// GetGroupByID retorna um grupo pelo id
func (ps *Group_service) GetGroupByID(id uint64) (*entity.GroupID, error) {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("call pcGetGroupDataById (?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	group := entity.GroupID{}

	stmt.QueryRow(id).Scan(
		&group.Group_id,
		&group.Group_name,
		&group.Customer.Customer_id,
		&group.Customer.Customer_name,
	)

	if group.Group_id == 0 {
		return nil, fmt.Errorf("no group found")
	}

	result, err := database.Query("call pcGetAllUserGroup (?)", id)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	user_list := []entity.User{}

	for result.Next() {
		user := entity.User{}

		if err := result.Scan(
			&user.User_id,
			&user.User_name); err != nil {
			fmt.Println(err.Error())
		} else {
			user_list = append(user_list, user)
		}

	}

	group.User = user_list

	return &group, nil

}

// UpdateStatusGroup atualiza o status do grupo  ATIVO/INATIVO
func (ps *Group_service) UpdateStatusGroup(id uint64, logID *int, ctx context.Context) (int64, error) {
	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT status_id FROM tblGroup WHERE group_id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	var statusGroup uint64

	err = stmt.QueryRow(id).Scan(&statusGroup)
	if err != nil {
		return 0, err
	}

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}
	defer status.Close()

	var statusID uint64

	err = status.QueryRow("GROUP", "ATIVO").Scan(&statusID)
	if err != nil {
		return 0, err
	}

	if statusID == statusGroup {
		statusGroup++
	} else {
		statusGroup--
	}

	result, err := tx.ExecContext(ctx, "UPDATE tblGroup SET status_id = ? WHERE group_id = ?", statusGroup, id)
	if err != nil {
		return 0, err
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsaff == 0 {
		return 0, nil
	}

	currentStatus, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}
	defer currentStatus.Close()

	if currentStatus.QueryRow("GROUP", "ATIVO").Scan(&statusID) == nil {
		if statusID == statusGroup {
			return 1, nil
		}
	}

	if currentStatus.QueryRow("GROUP", "INATIVO").Scan(&statusID) == nil {
		if statusID == statusGroup {
			return 2, nil
		}

	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// GetUsersGroup retorna todos os usuarios do grupo
func (ps *Group_service) GetUsersGroup(id uint64) (*entity.UserList, error) {

	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcGetAllUserGroup (?)", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list_users := &entity.UserList{}

	hasResult := false

	for rows.Next() {

		hasResult = true

		user := entity.User{}

		if err := rows.Scan(
			&user.User_id,
			&user.User_name,
		); err != nil {
			return nil, err
		} else {

			list_users.List = append(list_users.List, &user)

		}

	}

	if !hasResult {
		return nil, fmt.Errorf("no users found")
	}

	return list_users, nil

}

// CreateGroup cria um novo grupo
func (ps *Group_service) CreateGroup(group *entity.CreateGroup, logID *int, ctx context.Context) (int64, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	status, err := tx.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}
	defer status.Close()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	var statusID uint64

	err = status.QueryRow("GROUP", "ATIVO").Scan(&statusID)
	if err != nil {
		return 0, err
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO tblGroup (group_name, customer_id, status_id) VALUES (?, ?, ?)", group.Group_name, group.Customer_id, statusID)
	if err != nil {
		return 0, err
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		return 0, err

	}

	newid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	if group.GroupIDList.List != nil {

		ps.AttachUserGroup(&group.GroupIDList, uint64(newid), logID, ctx)

	}

	return rowsaff, nil

}

// AttachUserGroup adiciona usuarios ao grupo
func (ps *Group_service) AttachUserGroup(users *entity.GroupIDList, id uint64, logID *int, ctx context.Context) (int64, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO tblUserGroup (group_id, user_id) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	defer stmt.Close()

	for _, user := range users.List {
		_, err := tx.ExecContext(ctx, "INSERT INTO tblUserGroup (group_id, user_id) VALUES (?, ?)", id, user.ID)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int64(id), nil

}

// DetachUserGroup remove usuarios do grupo
func (ps *Group_service) DetachUserGroup(users *entity.GroupIDList, id uint64, logID *int, ctx context.Context) (int64, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	for _, user := range users.List {
		_, err := tx.ExecContext(ctx, "DELETE FROM tblUserGroup WHERE group_id = ? AND user_id = ?", id, user.ID)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int64(id), nil

}

// CountUsersGroup retorna a quantidade de usuarios do grupo
func (ps *Group_service) CountUsersGroup(id uint64) (*entity.CountUsersList, error) {

	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcCountUserGroup (?)", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	CountUserList := &entity.CountUsersList{}

	hasResult := false

	for rows.Next() {
		CountUser := entity.CountUser{}

		hasResult = true

		if err := rows.Scan(
			&CountUser.Grup_id,
			&CountUser.Qnt,
		); err != nil {
			return nil, err
		} else {

			CountUserList.List = append(CountUserList.List, &CountUser)

		}
	}

	if !hasResult {
		return nil, fmt.Errorf("no users found")
	}

	return CountUserList, nil

}

func (ps *Group_service) EditGroup(group *entity.EditGroup, id uint64, logID *int, ctx context.Context) (int64, error) {

	database := ps.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	_, err = database.ExecContext(ctx, "UPDATE tblGroup SET group_name = ?, customer_id = ? WHERE group_id = ?", group.Group_name, group.Customer, id)
	if err != nil {
		return 0, err
	}

	// Definir a variável de sessão "@user"
	_, err = tx.Exec("SET @user = ?", logID)
	if err != nil {
		return 0, errors.New("session variable error")
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM tblUserGroup WHERE group_id = ?", id)
	if err != nil {
		return 0, err
	}

	if group.Ids != nil {
		for _, user := range group.Ids {
			_, err := tx.ExecContext(ctx, "INSERT INTO tblUserGroup (group_id, user_id) VALUES (?, ?)", id, user.ID)
			if err != nil {
				return 0, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}
