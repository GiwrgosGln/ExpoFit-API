// routine.go
package models

// Routine represents a routine document.
type Routine struct {
	ID        string     `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    string     `json:"userID"`
	Title     string     `json:"title"`
	Exercises []Exercise `json:"exercises"`
}


// SetType represents the type of a set.
type SetType string