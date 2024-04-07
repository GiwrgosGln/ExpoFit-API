package models

type User struct {
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Gender      string      `json:"gender,omitempty"`
	DateOfBirth string      `json:"dateofbirth,omitempty"`
}
