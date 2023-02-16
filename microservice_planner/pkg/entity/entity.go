package entity

import (
	"encoding/json"
	"log"
	"strings"
)

type PlannerInterface interface {
	String() string
}

// Estrutura de dados de Planner para GET
type Planner struct {
	ID         uint64   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Date       string   `json:"date,omitempty"`
	Duration   uint64   `json:"duration,omitempty"`
	Subject    string   `json:"subject,omitempty"`
	Client     string   `json:"client,omitempty"`
	Business   string   `json:"business,omitempty"`
	Release    string   `json:"release,omitempty"`
	User       string   `json:"user,omitempty"`
	Status     string   `json:"status"`
	Created_At string   `json:"created_at,omitempty"`
	Guest      []Client `json:"guest,omitempty"`
}

// Estrutura de dados de Planner para CREATE e UPDATE
type PlannerUpdate struct {
	ID         uint64   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Date       string   `json:"date,omitempty"`
	Duration   int      `json:"duration,omitempty"`
	Subject    uint64   `json:"subject,omitempty"`
	Client     uint64   `json:"client,omitempty"`
	Release    uint64   `json:"release,omitempty"`
	User       uint64   `json:"user,omitempty"`
	Status     uint64   `json:"status"`
	Created_At string   `json:"created_at,omitempty"`
	Guest      []Client `json:"guest,omitempty"`
}

// Estrutura de dados de Client
type Client struct {
	ID   uint64 `json:"client_id,omitempty"`
	Name string `json:"client_name,omitempty"`
}

// Método de planner - retorna string com json do planner ou erro
func (p *Planner) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Planner to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de Planners
type PlannerList struct {
	List []*Planner `json:"list"`
}

// Estrutura de dados de groupID
type GroupID struct {
	ID uint64
}

// Estrutura de dados para lista de groupID
type GroupIDList struct {
	List []*GroupID
}

// Estrutura de dados de User
type User struct {
	ID uint64
}

// Estrutura de dados para lista de Users
type UserList struct {
	List []*User
}

// Método de PlannerList - retorna string com json da lista de planners ou erro
func (pl *PlannerList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert PlannerList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

func (p *Planner) Prepare() error {
	if erro := p.validate(); erro != nil {
		return erro
	}

	if erro := p.format(); erro != nil {
		return erro
	}

	return nil
}

func (p *Planner) format() error {
	p.Name = strings.TrimSpace(p.Name)

	return nil
}

func (p *Planner) validate() error {

	return nil
}
