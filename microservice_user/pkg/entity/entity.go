package entity

import "time"

type User struct {
	User_ID    uint64    `json:"user_id,omitempty"`
	Name       string    `json:"user_name,omitempty"`
	Email      string    `json:"user_email,omitempty"`
	Level      int       `json:"user_level,omitempty"`
	Created_At time.Time `json:"created_at,omitempty"`
	Status     string    `json:"status_id,omitempty"`
}
