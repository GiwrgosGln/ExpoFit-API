package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"FitnessAPI/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SaveWorkoutHandler handles the POST request to save a workout.
func SaveWorkoutHandler(c *gin.Context, collection *mongo.Collection) {
	var workout models.Workout

	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the date to the current time if not provided
	if workout.Date.IsZero() {
		workout.Date = time.Now()
	}

	// Insert the workout into MongoDB
	_, err := collection.InsertOne(context.Background(), workout)
	if err != nil {
		log.Printf("Error inserting workout: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert workout"})
		return
	}

	// Return the _id as part of the response
	c.JSON(http.StatusCreated, gin.H{"id": workout.ID})
}