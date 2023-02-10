package service

import (
	"fmt"
	"microservice_group/pkg/database"
	"microservice_group/pkg/entity"
)

type GroupServiceInterface interface {
	GetGroups(id uint64) (*entity.GroupList, error)
	GetGroupByID(id uint64) (*entity.Group, error)
	SoftDelete(id uint64) error
}

type Group_service struct {
	dbp database.DatabaseInterface
}

// GetGroup implements GroupServiceInterface

func NewGroupService(dabase_pool database.DatabaseInterface) *Group_service {
	return &Group_service{
		dabase_pool,
	}
}

func (ps *Group_service) GetGroups(id uint64) (*entity.GroupList, error) {

	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcGetAllUserGroup (?)", id)
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
			&group.Customer.Customer_name,
			&group.User.User_id,
			&group.User.User_name,
			&group.Status.Status_description,
		); err != nil {
			fmt.Println(err.Error())
		} else {

			list_groups.List = append(list_groups.List, &group)

		}

	}

	return list_groups, nil

}

func (ps *Group_service) GetGroupByID(id uint64) (*entity.Group, error) {

	database := ps.dbp.GetDB()

	rows, err := database.Query("call pcGetGroupDataById (?)", id)

	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	group := &entity.Group{}

	if rows.Next() {
		if err := rows.Scan(
			&group.Group_name,
			&group.Customer.Customer_name,
			&group.User.User_name,
			&group.User.User_id); err != nil {

			return &entity.Group{}, err
		}
	}

	return group, nil

}

func (ps *Group_service) SoftDelete(id uint64) error {

	database := ps.dbp.GetDB()

	_, err := database.Exec("update tblGroup set status_id = 6 if( ) where group_id = ?", id)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil

}
