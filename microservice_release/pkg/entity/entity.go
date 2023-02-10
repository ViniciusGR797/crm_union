package entity

import (
	"encoding/json"
	"log"
)

type ReleaseInterface interface {
	String() string
}

// Estrutura de dados de Release
type Release struct {
	ID                 uint64 `json:"release_id,omitempty"`
	Code               string `json:"release_code,omitempty"`
	Name               string `json:"release_name,omitempty"`
	Business_Name      string `json:"business_name,omitempty"`
	Status_Description string `json:"status_description,omitempty"`
	Tags               []Tag  `json:"tags,omitempty"`
}

type Tag struct {
	Tag_Name string `json:"tag_name,omitempty"`
}

// Retorna string com json do client ou err
func (c *Release) String() string {
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("error to convert Client to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de Client
type ReleaseList struct {
	List []*Release `json:"list"`
}

// Metodo de ClientList, retorna string com json da lista de client ou erro
func (cl *ReleaseList) String() string {
	data, err := json.Marshal(cl)

	if err != nil {
		log.Println("error to convert ClientList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}
