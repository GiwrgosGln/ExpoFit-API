// routine.go
package models

// Routine represents a routine document.
type Routine struct {
	ID        string              `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    string              `json:"userID"`
	Title     string              `json:"title"`
	Exercises []Exercise          `json:"exercises"`
}

// Exercise represents an exercise.
type Exercise struct {
	Name             string             `json:"name"`
	BodyPart         string             `json:"bodyPart"`
	Equipment        string             `json:"equipment"`
	GifURL           string             `json:"gifURL"`
	ID               string             `json:"id"`
	Instructions     []string           `json:"instructions"`
	SecondaryMuscles []string           `json:"secondaryMuscles"`
	Target           string             `json:"target"`
}
