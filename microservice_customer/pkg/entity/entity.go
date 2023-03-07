package entity

import (
	"encoding/json"
	"log"
)

type CustomerInterface interface {
	String() string
}

// Customer Estrutura de dados de Costumer
type Customer struct {
	ID         uint64 `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Created_At string `json:"created_at,omitempty"`
	Status     string `json:"status,omitempty"`
	Tags       []Tag  `json:"tags,omitempty"`
}

// Tag Estrutura de dados de Costumer
type Tag struct {
	Tag_ID   uint64 `json:"tag_id,omitempty"`
	Tag_Name string `json:"tag_name,omitempty"`
}

// String Método de customer - retorna string com json do customer ou erro
func (p *Customer) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Customer to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// CustomerList Estrutura de dados para lista de costumer
type CustomerList struct {
	List []*Customer `json:"list"`
}

// String Método de CustomerList - retorna string com json da lista de Customers ou erro
func (pl *CustomerList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert CostumerList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// NewCustomer Construtor de Customer - recebe dados no parâmetro
func NewCustomer(name string) *Customer {
	return &Customer{
		Name: name,
	}
}
