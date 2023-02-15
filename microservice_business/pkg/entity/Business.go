package entity

import (
	"encoding/json"
	"log"
)

type BusinessInterface interface {
	String() string
}

type Business struct {
	Business_id     uint64 `json:"business_id,omitempty"`
	Business_name   string `json:"business_name,omitempty"`
	Business_code   string `json:"business_code,omitempty"`
	BusinessSegment BusinessSegment
	Status          Status
	Tags            []Tag `json:"tags,omitempty"`
}

type Tag struct {
	Tag_ID   uint64 `json:"tag_id,omitempty"`
	Tag_Name string `json:"tag_name,omitempty"`
}

type BusinessSegment struct {
	BusinessSegment_id          int    `json:"businessSegment_id,omitempty"`
	BusinessSegment_description string `json:"businessSegment_description,omitempty"`
}

type Status struct {
	Status_id          int    `json:"status_id,omitempty"`
	Status_description string `json:"status_description,omitempty"`
}

func (p *Business) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Business to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type BusinessList struct {
	List []*Business `json:"business_list"`
}

func (pl *BusinessList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert BusinessList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type Business_Update struct {
	ID         uint64 `json:"business_id,omitempty"`
	Code       string `json:"business_code,omitempty"`
	Name       string `json:"business_name,omitempty"`
	Segment_Id int64  `json:"segment_id,omitempty"`
	Status_id  int64  `json:"status_id,omitempty"`
	Tags       []Tag  `json:"tags,omitempty"`
}

func NewBusiness(name string) *Business {
	return &Business{
		Business_name: name,
	}
}
