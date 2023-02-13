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
}

type SubjectID struct {
	Subject_id    uint64 `json:"subject_id,omitempty"`
	Subject_title string `json:"subject_title,omitempty"`
	Client        Client `json:"client,omitempty"`
	Business_name string `json:"business_name,omitempty"`
	Release_name  string `json:"release_name,omitempty"`
	Subject_text  string `json:"subject_text,omitempty"`
	Created_at    string `json:"created_at,omitempty"`
}

type Client struct {
	Client_id    uint64 `json:"client_id,omitempty"`
	Client_email string `json:"client_email,omitempty"`
	Client_name  string `json:"client_name,omitempty"`
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

type Status struct {
	Status_id          uint64 `json:"status_id,omitempty"`
	Status_description string `json:"status_description,omitempty"`
}
type Subject_list struct {
	List []*Subject `jason:"list,omitempty"`
}

func NewSubjecet(subject_id uint64, subject_title string, business string, client string, release string, user string, status_id uint64, status_description string) *Subject {
	return &Subject{
		Subject_id:    subject_id,
		Subject_title: subject_title,
		Business:      business,
		Client:        client,
		Release:       release,
		User:          user,
		Status:        Status{Status_id: status_id, Status_description: status_description},
	}
}
