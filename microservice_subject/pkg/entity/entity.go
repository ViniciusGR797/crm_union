package entity

import (
	"encoding/json"
	"log"
)

type Subject struct {
	Subject_id    uint64 `json:"subject_id,omitempty"`
	Subject_title string `json:"subject_title,omitempty"`
	User          string `json:"user_name,omitempty"`
	Release       string `json:"release_name,omitempty"`
	Business      string `json:"business_name,omitempty"`
	Client        string `json:"client_name,omitempty"`
	Status        Status `json:"status,omitempty"`
	Created_at    string `json:"created_at,omitempty"`
	Domain        Domain `json:"domain,omitempty"`
}

func (p *Subject) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type SubjectID struct {
	Subject_id    uint64 `json:"subject_id,omitempty"`
	Subject_title string `json:"subject_title,omitempty"`
	Client        Client `json:"client,omitempty"`
	Business_name string `json:"business_name,omitempty"`
	Release_name  string `json:"release_name,omitempty"`
	Subject_text  string `json:"subject_text,omitempty"`
	Created_at    string `json:"created_at,omitempty"`
	Domain        Domain `json:"domain,omitempty"`
	Status        Status `json:"status,omitempty"`
}

func (p *SubjectID) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type Client struct {
	Client_id    uint64 `json:"client_id,omitempty"`
	Client_email string `json:"client_email,omitempty"`
	Client_name  string `json:"client_name,omitempty"`
}

func (p *Client) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type Status struct {
	Status_id          uint64 `json:"status_id,omitempty"`
	Status_description string `json:"status_description,omitempty"`
}

func (p *Status) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type Subject_list struct {
	List []*Subject `jason:"list,omitempty"`
}

func (p *Subject_list) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type CreateSubject struct {
	Subject_title string `json:"subject_title,omitempty"`
	Subject_text  string `json:"subject_text,omitempty"`
	Subject_type  uint64 `json:"subject_type,omitempty"`
	Client_id     uint64 `json:"client_id,omitempty"`
	Release_id    uint64 `json:"release_id,omitempty"`
}

func (p *CreateSubject) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type Domain struct {
	Domain_id    uint64 `json:"domain_id,omitempty"`
	Domain_value string `json:"domain_value,omitempty"`
}

func (p *Domain) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert User to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type UpdateSubject struct {
	Subject_title string `json:"subject_title,omitempty"`
	Subject_text  string `json:"subject_text,omitempty"`
}
