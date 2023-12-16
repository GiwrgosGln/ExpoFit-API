package models

type User struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	Username string      `json:"username"`
	Sex      string      `json:"sex"`
	Weight   int         `json:"weight"`
}