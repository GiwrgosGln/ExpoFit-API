package models

type Exercise struct {
	ID               interface{} `json:"_id" bson:"_id,omitempty"`
	ExerciseID       string      `json:"id" bson:"id"`
	Name             string      `json:"name" bson:"name"`
	Target           string      `json:"target" bson:"target"`
	BodyPart         string      `json:"bodypart" bson:"bodypart"`
	Equipment        string      `json:"equipment" bson:"equipment"`
	GifURL           string      `json:"gifurl" bson:"gifurl"`
	SecondaryMuscles []string    `json:"secondarymuscles" bson:"secondarymuscles"`
	Instructions     []string    `json:"instructions" bson:"instructions"`
}