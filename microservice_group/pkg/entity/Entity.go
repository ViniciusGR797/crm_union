package entity

import (
	"encoding/json"
	"log"
)

type Group struct {
	Group_id   uint64   `json:"group_id,omitempty"`
	Group_name string   `json:"group_name,omitempty"`
	Created_at string   `json:"created_at,omitempty"`
	Status     Status   `json:"status,omitempty"`
	Customer   Custumer `json:"customers,omitempty"`
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

type GroupID struct {
	Group_id   uint64   `json:"group_id,omitempty"`
	Group_name string   `json:"group_name,omitempty"`
	Customer   Custumer `json:"customers,omitempty"`
	User       []User   `json:"users,omitempty"`
}

func (p *GroupID) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// tabela customer
type Custumer struct {
	Customer_id   int    `json:"customer_id,omitempty"`
	Customer_name string `json:"customer_name,omitempty"`
}

func (p *Custumer) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type User struct {
	User_id   int    `json:"user_id,omitempty"`
	User_name string `json:"user_name,omitempty"`
}

func (p *User) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type Status struct {
	Status_id          int    `json:"status_id,omitempty"`
	Status_description string `json:"status_name,omitempty"`
}

func (p *Status) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type CreateGroup struct {
	Group_name  string `json:"group_name,omitempty"`
	Customer_id int64  `json:"customer_id,omitempty"`
	GroupIDList `json:"users,omitempty"`
}

func (p *CreateGroup) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type CountUser struct {
	Grup_id uint64 `json:"group_id,omitempty"`
	Qnt     int    `json:"qnt,omitempty"`
}

func (p *CountUser) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type CountUsersList struct {
	List []*CountUser `json:"count_users_list,omitempty"`
}

func (p *CountUsersList) String() string {
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

func (p *ID) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type GroupIDList struct {
	List []*ID `json:"users_id"`
}

func (pl *GroupIDList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert UserList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type GroupList struct {
	List []*Group `json:"group_list"`
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

type UserList struct {
	List []*User `json:"user_list"`
}

func (pl *UserList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert UserList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}
