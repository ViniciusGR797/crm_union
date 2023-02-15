package service

import (
	"fmt"
	"microservice_group/pkg/database"
	"microservice_group/pkg/entity"
)

type GroupServiceInterface interface {
	GetGroups(id uint64) (*entity.GroupList, error)
	GetGroupByID(id uint64) (*entity.GroupID, error)
	UpdateStatusGroup(id uint64) (int64, error)
	GetUsersGroup(id uint64) (*entity.UserList, error)
	CreateGroup(group *entity.CreateGroup) (int64, error)
	AttachUserGroup(user *entity.GroupIDList, id uint64) (int64, error)
	DetachUserGroup(user *entity.GroupIDList, id uint64) (int64, error)
	CountUsersGroup(id uint64) (*entity.CountUsersList, error)
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

			list_groups.List = append(list_groups.List, &group)

		}

	}

	if !hasResult {
		return nil, fmt.Errorf("no groups found")
	}

	return list_groups, nil

}

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

func (ps *Group_service) UpdateStatusGroup(id uint64) (int64, error) {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblGroup WHERE group_id = ?")
	if err != nil {
		return 0, err
	}

	var statusGroup uint64

	err = stmt.QueryRow(id).Scan(&statusGroup)
	if err != nil {
		return 0, err
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}

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

	updt, err := database.Prepare("UPDATE tblGroup SET status_id = ? WHERE group_id = ?")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := updt.Exec(statusGroup, id)
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

	currentStatus, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}

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

	return 0, nil
}

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

func (ps *Group_service) CreateGroup(group *entity.CreateGroup) (int64, error) {

	database := ps.dbp.GetDB()

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		return 0, err
	}

	var statusID uint64

	err = status.QueryRow("GROUP", "ATIVO").Scan(&statusID)
	if err != nil {
		return 0, err
	}

	stmt, err := database.Prepare("INSERT INTO tblGroup (group_name, customer_id, status_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(group.Group_name, group.Customer_id, statusID)
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

		ps.AttachUserGroup(&group.GroupIDList, uint64(newid))

	}

	return rowsaff, nil

}

func (ps *Group_service) AttachUserGroup(users *entity.GroupIDList, id uint64) (int64, error) {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO tblUserGroup (group_id, user_id) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	for _, user := range users.List {
		_, err := stmt.Exec(id, user.ID)
		if err != nil {
			return 0, err
		}

	}

	return int64(id), nil

}

func (ps *Group_service) DetachUserGroup(users *entity.GroupIDList, id uint64) (int64, error) {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM tblUserGroup WHERE group_id = ? AND user_id = ?")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	for _, user := range users.List {
		_, err := stmt.Exec(id, user.ID)
		if err != nil {
			return 0, err
		}

	}

	return int64(id), nil

}

// count users in group
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
