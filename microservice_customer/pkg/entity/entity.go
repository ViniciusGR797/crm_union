package entity

import (
	"encoding/json"
	"log"
)

type CustomerInterface interface {
	String() string
}

// Estrutura de dados de Costumer
type Customer struct {
	ID         uint64 `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Created_At string `json:"created_at,omitempty"`
	Status     string `json:"status,omitempty"`
}

// Método de customer - retorna string com json do customer ou erro
func (p *Customer) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Customer to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de costumer
type CustomerList struct {
	List []*Customer `json:"list"`
}

// Método de CustomerList - retorna string com json da lista de Customers ou erro
func (pl *CustomerList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert CostumerList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Construtor de Customer - recebe dados no parâmetro e transforma em um user
func NewCostumer(name, created_at, status string, id uint64) *Customer {
	return &Customer{
		ID:         id,
		Name:       name,
		Created_At: created_at,
		Status:     status,
	}
}

// testetetetete
