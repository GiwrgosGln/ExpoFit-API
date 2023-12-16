package models

// SetType represents the type of a set.
type SetType string

const (
	WarmUp   SetType = "warm-up"
	Working  SetType = "working"
	Failure  SetType = "failure"
	Drop     SetType = "drop"
)

// Set represents a set in a routine exercise.
type Set struct {
	SetType SetType `json:"setType"`
	Sets    int     `json:"sets"`
}

// RoutineExercise represents an exercise in a routine.
type RoutineExercise struct {
	ExerciseName string `json:"exerciseName"`
	Sets         []Set  `json:"sets"`
}

// Routine represents a routine document.
type Routine struct {
	ID           string            `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID       string            `json:"userID"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Exercises    []RoutineExercise `json:"exercises"`
}