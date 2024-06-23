package models

import "time"

type ExerciseInWorkout struct {
	ID     string `json:"_id" bson:"_id,omitempty"`
	ExerciseID string `json:"exercise_id" bson:"exercise_id"`
	Name   string `json:"name" bson:"name"`
	Sets   []Set  `json:"sets" bson:"sets"`
}

type Workout struct {
	ID        string               `json:"id" bson:"_id,omitempty"`
	UserID    string               `json:"user_id" bson:"user_id"`
	Date      time.Time            `json:"date"`
	RoutineName string               `json:"routine_name" bson:"routine_name"`
	Exercises []ExerciseInWorkout `json:"exercises"`
}

type Set struct {
	Type   string `json:"type"`
	Reps   *int    `json:"reps"`
	Weight *int    `json:"weight"`
	RPE    *string    `json:"rpe"`
}