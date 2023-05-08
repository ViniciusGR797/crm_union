package entity

import (
	"encoding/json"
	"log"
)

// RemarkInterface que define um único método chamado String. A interface é usada para especificar um contrato que um tipo deve cumprir para ser considerado um Remark.
type RemarkInterface interface {
	String() string
}

// Remark Estrutura de dados de Remark
type Remark struct {
	ID                 uint64 `json:"id,omitempty"`
	Remark_Name        string `json:"remark_name,omitempty"`
	Text               string `json:"text,omitempty"`
	Date               string `json:"date,omitempty"`
	Date_Return        string `json:"date_return,omitempty"`
	Status_Description string `json:"status_description,omitempty"`
	Client_ID          uint64 `json:"client_id,omitempty"` //auterado
	Client_Name        string `json:"client_name,omitempty"`
	Client_Email       string `json:"client_email,omitempty"`
	User_Name          string `json:"user_name,omitempty"`
	User_ID            uint64 `json:"user_id,omitempty"`
	CreatedBy_id       uint64 `json:"createdBy_id,omitempty"`
	CreatedBy_name     string `json:"createdBy_name,omitempty"`
	Release_ID         uint64 `json:"release_id,omitempty"` //auterado
	Release_Name       string `json:"release_name,omitempty"`
	Subject_ID         uint64 `json:"subject_id,omitempty"` //auterado
	Subject_Name       string `json:"subject_name,omitempty"`
	Subject_Title      string `json:"subject_title,omitempty"`
	Business_ID        uint64 `json:"business_id,omitempty"` //auterado
	Business_Name      string `json:"business_name,omitempty"`
}

//RemarkUpdate Estrutura de dados de Remark utilizada para criar, atualizar e efetuar softdelete

type RemarkUpdate struct {
	Remark_Name string `json:"remark_name,omitempty"`
	Text        string `json:"text,omitempty"`
	Date        string `json:"date,omitempty"`
	Date_Return string `json:"date_return,omitempty"`
	Subject_ID  uint64 `json:"subject_id,omitempty"`
	Client_ID   uint64 `json:"client_id,omitempty"`
	Release_ID  uint64 `json:"release_id,omitempty"`
	User_ID     uint64 `json:"user_id,omitempty"`
	Status_ID   uint64 `json:"status_id,omitempty"`
}

// String para o tipo Remark, que implementa a interface RemarkInterface. Método de Remark - retorna string com json do Remark ou erro
func (p *Remark) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Remark to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

// RemarkList Estrutura de dados para lista de Remark
type RemarkList struct {
	List []*Remark `json:"list"`
}

// String Método de RemarkList - retorna string com json da lista de Remarks ou erro
func (pl *RemarkList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert RemarkList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}
