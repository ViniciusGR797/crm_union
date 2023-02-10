package entity

import (
	"encoding/json"
	"log"
)

type UserGroup struct {
	User_id  uint64 `json:"user_id"` //chave estrangeira
	Group_id Group
}

func NewUserGroup(user_id uint64, group_id Group) *UserGroup {
	return &UserGroup{
		User_id:  user_id,
		Group_id: group_id,
	}
}

func (p *UserGroup) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert UserGroup to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

type UserGroupList struct {
	List []*UserGroup `json:"user_group_list"`
}

func (pl *UserGroupList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert UserGroupList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}
