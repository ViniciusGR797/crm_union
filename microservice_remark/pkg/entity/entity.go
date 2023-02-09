package entity

import (
	"encoding/json"
	"log"
	"time"
)

type RemarkInterface interface {
	String() string
}

// Estrutura de dados de Remark
type Remark struct {
	ID                 uint64     `json:"id,omitempty"`
	Subject            string     `json:"subject_title,omitempty"`
	Text               string     `json:"text,omitempty"`
	Date               *time.Time `json:"date,omitempty"`
	Date_Return        *time.Time `json:"date_return,omitempty"`
	Status_Description string     `json:"status_description,omitempty"`
	Client_Name        string     `json:"client_name,omitempty"`
	Client_Email       string     `json:"client_email,omitempty"`
	User_Name          string     `json:"user_name,omitempty"`
	Release_Name       string     `json:"release_name,omitempty"`
	Subject_ID         uint64     `json:"subject_id,omitempty"`
	Business_Name      string     `json:"business_name,omitempty"`
}

// Estrutura de dados de Remark utilizada para criar, atualizar e efetuar softdelete

// Método de Remark - retorna string com json do Remark ou erro
func (p *Remark) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Remark to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Estrutura de dados para lista de Remark
type RemarkList struct {
	List []*Remark `json:"list"`
}

// Método de RemarkList - retorna string com json da lista de Remarks ou erro
func (pl *RemarkList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert RemarkList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// Construtor de Remark - recebe dados no parâmetro e transforma em um Remark
// func NewRemark(ID uint64, subject, text string, date, date_return time.Time, release_id, user_id, client_id, date, text, subject, created_at, status_id uint64, id uint64) *Remark {
// 	return &Remark{
// 		ID:            ID,
// 		Subject:       subject,
// 		Created_At:    created_at,
// 		Text:          text,
// 		Date:          date,
// 		Date_Return: date_return,
// 		Status_Description:     status_id,
// 		Client_Id:     client_id,
// 		User_Id:       user_id,
// 		Release_Id:    release_id,
// 		Subject_ID:    subject_id,
// 	}
// }
