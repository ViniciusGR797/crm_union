package service

import (
	"fmt"
	"log"
	"microservice_group/pkg/database"
	"microservice_group/pkg/entity"
)

type GroupServiceInterface interface {
	GetGroups(id uint64) (*entity.GroupList, error)
	GetGroupByID(id uint64) (*entity.GroupID, error)
	UpdateStatusGroup(id uint64) int64
	GetUsersGroup(id uint64) (*entity.UserList, error)
	CreateGroup(group *entity.CreateGroup) int64
	AttachUserGroup(user *entity.GroupIDList, id uint64) int64
	DetachUserGroup(user *entity.GroupIDList, id uint64) int64
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
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_groups := &entity.GroupList{}

	for rows.Next() {

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

	return list_groups, nil

}

func (ps *Group_service) GetGroupByID(id uint64) (*entity.GroupID, error) {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("call pcGetGroupDataById (?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	group := entity.GroupID{}

	stmt.QueryRow(id).Scan(
		&group.Group_id,
		&group.Group_name,
		&group.Customer.Customer_id,
		&group.Customer.Customer_name,
	)

	result, err := database.Query("call pcGetAllUserGroup (?)", id)
	if err != nil {
		fmt.Println(err.Error())
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

func (ps *Group_service) UpdateStatusGroup(id uint64) int64 {
	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("SELECT status_id FROM tblGroup WHERE group_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusGroup uint64

	err = stmt.QueryRow(id).Scan(&statusGroup)
	if err != nil {
		log.Println(err.Error())
	}

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("GROUP", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	if statusID == statusGroup {
		statusGroup++
	} else {
		statusGroup--
	}

	updt, err := database.Prepare("UPDATE tblGroup SET status_id = ? WHERE group_id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := updt.Exec(statusGroup, id)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	if rowsaff == 0 {
		return 0
	}

	currentStatus, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	if currentStatus.QueryRow("GROUP", "ATIVO").Scan(&statusID) == nil {
		if statusID == statusGroup {
			return 1
		}
	}

	if currentStatus.QueryRow("GROUP", "INATIVO").Scan(&statusID) == nil {
		if statusID == statusGroup {
			return 2
		}

	}

	return 0
}

func (ps *Group_service) GetUsersGroup(id uint64) (*entity.UserList, error) {

	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcGetAllUserGroup (?)", id)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	list_users := &entity.UserList{}

	for rows.Next() {

		user := entity.User{}

		if err := rows.Scan(
			&user.User_id,
			&user.User_name,
		); err != nil {
			fmt.Println(err.Error())
		} else {

			list_users.List = append(list_users.List, &user)

		}

	}

	return list_users, nil

}

func (ps *Group_service) CreateGroup(group *entity.CreateGroup) int64 {

	database := ps.dbp.GetDB()

	status, err := database.Prepare("SELECT status_id FROM tblStatus WHERE status_dominio = ? AND status_description = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	var statusID uint64

	err = status.QueryRow("GROUP", "ATIVO").Scan(&statusID)
	if err != nil {
		log.Println(err.Error())
	}

	stmt, err := database.Prepare("INSERT INTO tblGroup (group_name, customer_id, status_id) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(group.Group_name, group.Customer_id, statusID)
	if err != nil {
		fmt.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}

	return rowsaff

}

func (ps *Group_service) AttachUserGroup(users *entity.GroupIDList, id uint64) int64 {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("INSERT INTO tblUserGroup (group_id, user_id) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	for _, user := range users.List {
		_, err := stmt.Exec(id, user.ID)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

	return int64(id)

}

func (ps *Group_service) DetachUserGroup(users *entity.GroupIDList, id uint64) int64 {

	database := ps.dbp.GetDB()

	stmt, err := database.Prepare("DELETE FROM tblUserGroup WHERE group_id = ? AND user_id = ?")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	for _, user := range users.List {
		_, err := stmt.Exec(id, user.ID)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

	return int64(id)

}
