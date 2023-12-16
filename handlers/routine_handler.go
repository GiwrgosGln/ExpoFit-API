package handlers

import (
	"context"
	"log"
	"net/http"

	"FitnessAPI/models"

	"github.com/gin-gonic/gin"
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