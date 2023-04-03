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
	Business_Id        string `json:"business_id,omitempty"`
	Status_Description string `json:"status_description,omitempty"`
	Tags               []Tag  `json:"tags,omitempty"`
}

type Release_Update struct {
	ID          uint64 `json:"release_id,omitempty"`
	Code        string `json:"release_code,omitempty"`
	Name        string `json:"release_name,omitempty"`
	Business_ID uint64 `json:"business_id,omitempty"`
	Status_ID   uint64 `json:"status_id,omitempty"`
	Tags        []Tag  `json:"tags,omitempty"`
}

type Tag struct {
	Tag_ID   uint64 `json:"tag_id,omitempty"`
	Tag_Name string `json:"tag_name,omitempty"`
}

// Retorna string com json do release ou err
func (c *Release) String() string {
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("error to convert release to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de release
type ReleaseList struct {
	List []*Release `json:"list"`
}

// Metodo de releaselist, retorna string com json da lista de release ou erro
func (cl *ReleaseList) String() string {
	data, err := json.Marshal(cl)

	if err != nil {
		log.Println("error to convert releaselist to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}
