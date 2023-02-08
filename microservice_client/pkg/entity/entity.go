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
	ID          uint64    `json:"client_id,omitempty"`
	Name        string    `json:"client_name,omitempty"`
	Email       string    `json:"client_email,omitempty"`
	Role        string    `json:"client_role,omitempty"`
	Costumer_ID string    `json:"costumer_id,omitempty"`
	Relase_ID   string    `json:"relase_id,omitempty"`
	Business_ID string    `json:"business_id,omitempty"`
	User_ID     string    `json:"user_id,omitempty"`
	Created_At  time.Time `json:"created_at,omitempty"`
	Status_ID   string    `json:"status_id,omitempty"`
}

type ClientUpdate struct {
	ID          uint64 `json:"client_id,omitempty"`
	Name        string `json:"client_name,omitempty"`
	Email       string `json:"client_email,omitempty"`
	Role        uint64 `json:"client_role,omitempty"`
	Costumer_ID uint64 `json:"costumer_id,omitempty"`
	Relase_ID   uint64 `json:"relase_id,omitempty"`
	Business_ID uint64 `json:"business_id,omitempty"`
	User_ID     uint64 `json:"user_id,omitempty"`
	Status_ID   uint64 `json:"status_id,omitempty"`
}

// retorna string com json do client ou err
func (c *Client) String() string {
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("error to convert Client to JASON")
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
func NewClient(client_id uint64, client_name string, client_email string, client_role uint64, costumer_id uint64, relase_id uint64) *ClientUpdate {
	return &ClientUpdate{
		ID:          client_id,
		Name:        client_name,
		Email:       client_email,
		Role:        client_role,
		Costumer_ID: costumer_id,
		Relase_ID:   relase_id,
	}
}
