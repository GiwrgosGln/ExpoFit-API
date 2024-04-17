package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"FitnessAPI/models"
)

// CreateMeasurementHandler handles the POST request to create a new measurement.
func CreateMeasurementHandler(c *gin.Context, collection *mongo.Collection) {
	// Parse the request body into a Measurement object
	var measurement models.Measurement
	if err := c.BindJSON(&measurement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Add the current date to the measurement
	measurement.Date = time.Now()

	// Insert the measurement into the database
	result, err := collection.InsertOne(context.Background(), measurement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create measurement"})
		return
	}

	// Return the ID of the newly created measurement
	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// GetMeasurementsByUserIDHandler handles the GET request to retrieve all measurements for a specific user ID.
func GetMeasurementsByUserIDHandler(c *gin.Context, collection *mongo.Collection) {
	// Extract the user ID from the request path
	userID := c.Param("userID")

	// Define a filter to query measurements by user ID
	filter := bson.M{"user_id": userID}

	// Find all measurements matching the filter
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch measurements"})
		return
	}
	defer cursor.Close(context.Background())

	// Initialize a slice to store measurements
	var measurements []models.Measurement

	// Iterate over the cursor and decode each measurement
	for cursor.Next(context.Background()) {
		var measurement models.Measurement
		if err := cursor.Decode(&measurement); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode measurements"})
			return
		}
		measurements = append(measurements, measurement)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error"})
		return
	}

	// Return the measurements as a response
	c.JSON(http.StatusOK, measurements)
}

// DeleteMeasurementHandler handles the DELETE request to delete a measurement by its ID.
func DeleteMeasurementHandler(c *gin.Context, collection *mongo.Collection) {
	// Extract the measurement ID from the request path
	measurementID := c.Param("id")
	log.Println("Measurement ID:", measurementID)

	// Convert the measurement ID string to BSON ObjectId
	objectID, err := primitive.ObjectIDFromHex(measurementID)
	if err != nil {
		log.Println("Invalid measurement ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid measurement ID"})
		return
	}

	// Define a filter to find the measurement by its ID
	filter := bson.M{"_id": objectID}

	// Delete the measurement from the database
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting measurement:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete measurement"})
		return
	}

	// Check if the measurement was found and deleted
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Measurement not found"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Measurement deleted successfully"})
}
