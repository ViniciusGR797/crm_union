package entity

import (
	"encoding/json"
	"log"
)

type UserInterface interface {
	String() string
}

// Estrutura de dados de User
type User struct {
	ID         uint64 `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Level      int    `json:"level,omitempty"`
	Created_At string `json:"created_at,omitempty"`
	Status     string `json:"status,omitempty"`
}

// Método de user - retorna string com json do user ou erro
func (p *User) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de Users
type UserList struct {
	List []*User `json:"list"`
}

// Método de UserList - retorna string com json da lista de users ou erro
func (pl *UserList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert UserList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Construtor de User - recebe dados no parâmetro e transforma em um user
func NewUser(name, email, created_at, status string, level int, id uint64) *User {
	return &User{
		ID:         id,
		Name:       name,
		Email:      email,
		Level:      level,
		Created_At: created_at,
		Status:     status,
	}
}
