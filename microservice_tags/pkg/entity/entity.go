package entity

import (
	"encoding/json"
	"log"
)

// TagInterface  interface para padronização do CRUD para transformar em JSON.
type TagsInterface interface {
	String() string
}

// Tag estrutura de dados de Tag
type Tags struct {
	Tag_ID   uint64 `json:"tag_id,omitempty"`
	Tag_Name string `json:"tag_name,omitempty"`
}

// String converte em Json a estrutra passada
func (p *Tags) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Tag to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// TagList  lista para Tag
type TagsList struct {
	List []*Tags `json:"tag_list"`
}

func (pl *TagsList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert TagList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Tag_Update estrutra de dados para Tag_Update
type Tag_Update struct {
	Tag_ID   uint64 `json:"tag_id,omitempty"`
	Tag_Name string `json:"tag_name,omitempty"`
}

func NewTag(name string) *Tags {
	return &Tags{
		Tag_Name: name,
	}
}
