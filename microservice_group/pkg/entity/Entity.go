package entity

import (
	"encoding/json"
	"log"
)

type Group struct {
	Group_id   uint64   `json:"group_id,omitempty"`
	Group_name string   `json:"group_name,omitempty"`
	Created_at string   `json:"created_at,omitempty"`
	Status     Status   `json:"-,omitempty"`
	Customer   Custumer `json:"customers,omitempty"`
}

type GroupID struct {
	Group_id   uint64   `json:"group_id,omitempty"`
	Group_name string   `json:"group_name,omitempty"`
	Customer   Custumer `json:"customers,omitempty"`
	User       []User   `json:"users,omitempty"`
}

// tabela customer
type Custumer struct {
	Customer_id   int    `json:"customer_id,omitempty"`
	Customer_name string `json:"customer_name,omitempty"`
}

type User struct {
	User_id   int    `json:"user_id,omitempty"`
	User_name string `json:"user_name,omitempty"`
}

type Status struct {
	Status_id          int    `json:"status_id,omitempty"`
	Status_description string `json:"status_name,omitempty"`
}

type CreateGroup struct {
	Group_name  string `json:"group_name,omitempty"`
	Customer_id int64  `json:"customer_id,omitempty"`
}

func (p *Group) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type ID struct {
	ID int `json:"id"`
}
type GroupIDList struct {
	List []*ID `json:"users_id"`
}
type GroupList struct {
	List []*Group `json:"group_list"`
}

type UserList struct {
	List []*User `json:"user_list"`
}

func (pl *GroupList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert UserList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

func NewGroup(group_name, created_at string, status_id, customer_id int, group_id uint64) *Group {
	return &Group{
		Group_id:   group_id,
		Group_name: group_name,
		Status:     Status{Status_id: status_id},
		Created_at: created_at,
		Customer:   Custumer{Customer_id: customer_id},
	}
}
