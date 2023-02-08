package service

import (
	"fmt"
	"microservice_group/pkg/database"
	"microservice_group/pkg/entity"
)

type GroupServiceInterface interface {
	GetGroups() *entity.GroupList
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

func (ps *Group_service) GetGroups() *entity.GroupList {

	database := ps.dbp.GetDB()

	rows, err := database.Query("select g.group_id, g.group_name, g.status_id, s.status_description, g.created_at,g.customer_id, c.customer_name from tblGroup g inner join tblCustomer c on g.customer_id = c.customer_id inner join tblStatus s on g.status_id = s.status_id")
	// verifica se teve erro
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	lista_groups := &entity.GroupList{}

	for rows.Next() {

		group := entity.Group{}

		if err := rows.Scan(
			&group.Group_id,
			&group.Group_name,
			&group.Status.Status_id,
			&group.Status.Status_description,
			&group.Created_at,
			&group.Customer.Customer_id,
			&group.Customer.Customer_name); err != nil {
			fmt.Println(err.Error())
		} else {

			lista_groups.List = append(lista_groups.List, &group)
		}

	}

	return lista_groups

}
