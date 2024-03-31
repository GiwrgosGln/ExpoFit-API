package models

import (
	"time"
)

// ExerciseSet represents a set of exercises in a workout.
type ExerciseSet struct {
	Reps *int     `json:"reps,omitempty" bson:"reps,omitempty"` // Use pointer to int to allow null values
	KG   *float64 `json:"kg,omitempty" bson:"kg,omitempty"`     // Use pointer to float64 to allow null values
	RPE  *float64 `json:"rpe,omitempty" bson:"rpe,omitempty"`   // Use pointer to float64 to allow null values
}

// WorkoutExercise represents an exercise in a workout.
type WorkoutExercise struct {
	Sets []ExerciseSet `json:"sets"`
}

// Workout represents a workout document.
type Workout struct {
	ID        string                     `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    string                     `json:"userID"`
	Date      time.Time                  `json:"date"`
	Exercises map[string]WorkoutExercise `json:"exercises"`
	Duration  time.Duration              `json:"duration"`
}
