package handlers

import (
	"context"
	"log"
	"net/http"

	"FitnessAPI/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateRoutineHandler handles the creation of a new routine.
func CreateRoutineHandler(c *gin.Context, collection *mongo.Collection) {
	var routine models.Routine

	if err := c.ShouldBindJSON(&routine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the routine into MongoDB with the specified _id
	_, err := collection.InsertOne(context.Background(), routine)
	if err != nil {
		log.Printf("Error inserting routine: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert routine"})
		return
	}

	// Return the _id as part of the response
	c.JSON(http.StatusCreated, gin.H{"id": routine.ID})
}

// GetRoutineHandler handles the fetching of a routine by UserID.
func GetRoutineHandler(c *gin.Context, collection *mongo.Collection) {
	// Get the user ID from the request parameters
	userID := c.Param("id")

	// Define a filter to find the routine by UserID
	filter := bson.M{"userid": userID}

	// Find the routine in MongoDB
	var routine models.Routine
	err := collection.FindOne(context.Background(), filter).Decode(&routine)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Routine not found"})
			return
		}
		log.Printf("Error fetching routine: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routine"})
		return
	}

	// Return the routine as part of the response
	c.JSON(http.StatusOK, routine)
}


