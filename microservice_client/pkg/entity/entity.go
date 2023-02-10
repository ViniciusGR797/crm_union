package entity

import (
	"encoding/json"
	"log"
	"time"
)

type ClienteInteface interface {
	String() string
}

// Estrutura de dados de Client
type Client struct {
	ID                 uint64    `json:"client_id,omitempty"`
	Name               string    `json:"client_name,omitempty"`
	Email              string    `json:"client_email,omitempty"`
	Role               string    `json:"client_role,omitempty"`
	Customer_Name      string    `json:"customer_name,omitempty"`
	Release_Name       string    `json:"release_name,omitempty"`
	Business_Name      string    `json:"business_name,omitempty"`
	User_Name          string    `json:"user_name,omitempty"`
	Created_At         time.Time `json:"created_at,omitempty"`
	Status_Description string    `json:"status_description,omitempty"`
	Tags               []Tag     `json:"tags,omitempty"`
}

type Tag struct {
	Tag_Name string `json:"tag_name,omitempty"`
}

// Estrutura de dados de Client para softdelete, create e update
type ClientUpdate struct {
	ID          uint64   `json:"client_id,omitempty"`
	Name        string   `json:"client_name,omitempty"`
	Email       string   `json:"client_email,omitempty"`
	Role        uint64   `json:"client_role,omitempty"`
	Customer_ID uint64   `json:"customer_id,omitempty"`
	Release_ID  uint64   `json:"release_id,omitempty"`
	Business_ID uint64   `json:"business_id,omitempty"`
	User_ID     uint64   `json:"user_id,omitempty"`
	Status_ID   uint64   `json:"status_id,omitempty"`
	Tags        []uint64 `json:"tags,omitempty"`
}

// retorna string com json do client ou err
func (c *Client) String() string {
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("error to convert Client to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de Client
type ClientList struct {
	List []*Client `json:"list"`
}

// Metodo de ClientList, retorna string com json da lista de client ou erro
func (cl *ClientList) String() string {
	data, err := json.Marshal(cl)

	if err != nil {
		log.Println("error to convert ClientList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// NewClient Construtor de Client - recebe dados no parâmetro e transforma em um client
func NewClient(client_id uint64, client_name string, client_email string, client_role uint64, customer_id uint64, release_id uint64) *ClientUpdate {
	return &ClientUpdate{
		ID:          client_id,
		Name:        client_name,
		Email:       client_email,
		Role:        client_role,
		Customer_ID: customer_id,
		Release_ID:  release_id,
	}
}
