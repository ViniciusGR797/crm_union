package entity

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/badoux/checkmail"
)

type UserInterface interface {
	String() string
}

// Estrutura de dados de User
type User struct {
	ID         uint64 `json:"id,omitempty"`
	TCS_ID     uint64 `json:"tcs_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Level      uint   `json:"level,omitempty"`
	Created_At string `json:"created_at,omitempty"`
	Status     string `json:"status,omitempty"`
	Password   string `json:"password,omitempty"`
	Hash       string `json:"-"`
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
func NewUser(name, email, created_at, status string, level uint, id, tcs_id uint64) *User {
	return &User{
		ID:         id,
		TCS_ID:     tcs_id,
		Name:       name,
		Email:      email,
		Level:      level,
		Created_At: created_at,
		Status:     status,
	}
}

// Estrutura de dados de groupID
type GroupID struct {
	ID uint64
}

// Estrutura de dados para lista de groupID
type GroupIDList struct {
	List []*GroupID
}

func (user *User) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	if erro := user.format(); erro != nil {
		return erro
	}

	return nil
}

func (user *User) format() error {
	user.Name = strings.TrimSpace(user.Name)

	return nil
}

func (user *User) validate() error {
	// Verifica se tcs_id está vazio
	if user.TCS_ID == 0 {
		return errors.New("the tcs id is mandatory and cannot be blank")

	}

	// Verifica se nome está vazio
	if user.Name == "" {
		return errors.New("the name is mandatory and cannot be blank")

	}

	// Verifica se level é válido
	if user.Level < 1 || user.Level > 5 {
		return errors.New("invalid level")
	}

	// Verifica se o status é válido
	if user.Status != "ATIVO" && user.Status != "INATIVO" && user.Status != "" {
		return errors.New("invalid status")
	}

	// Verifica se email formato válido
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email")
	}

	// Verifica se senha tem o tamanho mínimo de caracteres
	if len(user.Password) < 8 {
		return errors.New("password too short")
	}

	return nil
}
