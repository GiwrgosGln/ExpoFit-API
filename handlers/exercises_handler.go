package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"FitnessAPI/models"
)

// GetAllExercisesHandler retrieves all exercises from the database
func GetAllExercisesHandler(c *gin.Context, collection *mongo.Collection) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve exercises"})
		return
	}
	defer cursor.Close(context.Background())

	var exercises []models.Exercise
	if err := cursor.All(context.Background(), &exercises); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode exercises"})
		return
	}

	c.JSON(http.StatusOK, exercises)
}
