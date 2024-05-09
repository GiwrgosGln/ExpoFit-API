package models

import "time"

// Measurement represents a user's measurement data.
type Measurement struct {
	ID         string    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BodyWeight float64   `json:"body_weight,omitempty" bson:"body_weight,omitempty"`
	Date       time.Time `json:"date,omitempty" bson:"date,omitempty"`
}
