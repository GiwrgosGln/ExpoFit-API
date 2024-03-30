package models

import "time"

// ExerciseInWorkout represents an exercise within a workout, excluding unnecessary fields.
type ExerciseInWorkout struct {
	ID     string `json:"_id" bson:"_id,omitempty"`
	ExerciseID string `json:"exercise_id" bson:"exercise_id"`
	Name   string `json:"name" bson:"name"`
	Sets   []Set  `json:"sets" bson:"sets"`
}

// Workout represents a workout.
type Workout struct {
	ID        string               `json:"id" bson:"_id,omitempty"`
	UserID    string               `json:"user_id" bson:"user_id"`
	Title     string               `json:"title"`
	Date      time.Time            `json:"date"`
	Exercises []ExerciseInWorkout `json:"exercises"`
}

// Set represents a set within an exercise.
type Set struct {
	Type   string `json:"type"`
	Reps   int    `json:"reps"`
	Weight int    `json:"weight"`
	RPE    string    `json:"rpe"`
}
