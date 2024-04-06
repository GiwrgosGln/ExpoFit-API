package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"FitnessAPI/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetWorkoutsByUserIDHandler handles the GET request to fetch workouts by user ID.
func GetWorkoutsByUserIDHandler(c *gin.Context, collection *mongo.Collection) {
	userID := c.Param("userID") // Extract the user ID from the request path

	// Define a filter to query workouts by user ID
	filter := bson.M{"user_id": userID}

	// Execute the find operation to get workouts matching the filter
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workouts"})
		return
	}
	defer cursor.Close(context.Background())

	// Initialize a slice to store fetched workouts
	var workouts []models.Workout

	// Iterate through the cursor and decode each document into a Workout struct
	for cursor.Next(context.Background()) {
		var workout models.Workout
		if err := cursor.Decode(&workout); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode workout"})
			return
		}
		workouts = append(workouts, workout)
	}

	// Check if any error occurred during cursor iteration
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor iteration error"})
		return
	}

	// Return the fetched workouts as a response
	c.JSON(http.StatusOK, workouts)
}

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

// DeleteWorkoutHandler handles the deletion of a workout by ID.
func DeleteWorkoutHandler(c *gin.Context, collection *mongo.Collection) {
    // Get the workout ID from the request parameters
    workoutID := c.Param("id")

    // Convert the workout ID to an ObjectId
    objID, err := primitive.ObjectIDFromHex(workoutID)
    if err != nil {
        // Return error response if the ID is invalid
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workout ID"})
        return
    }

    // Define a filter to find the workout by ID
    filter := bson.M{"_id": objID}

    // Delete the workout from MongoDB
    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        log.Printf("Error deleting workout: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete workout"})
        return
    }

    // Check if any workouts were deleted
    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No workout found with specified ID"})
        return
    }

    // Return success message
    c.JSON(http.StatusOK, gin.H{"message": "Routine deleted successfully"})
}