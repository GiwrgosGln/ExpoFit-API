package handlers

import (
	"context"
	"log"
	"net/http"

	"FitnessAPI/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetRoutinesByUserIDHandler handles the fetching of all routines with a specific UserID.
func GetRoutinesByUserIDHandler(c *gin.Context, collection *mongo.Collection) {
	// Get the user ID from the request parameters
	userID := c.Param("id")

	// Define a filter to find all routines by UserID
	filter := bson.M{"userid": userID}

	// Find all routines in MongoDB
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error fetching routines: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routines"})
		return
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and collect routines
	var routines []models.Routine
	for cursor.Next(context.Background()) {
		var routine models.Routine
		if err := cursor.Decode(&routine); err != nil {
			log.Printf("Error decoding routine: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode routine"})
			return
		}
		routines = append(routines, routine)
	}

	// Check if any routines were found
	if len(routines) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No routines found"})
		return
	}

	// Return the routines as part of the response
	c.JSON(http.StatusOK, routines)
}

// DeleteRoutineHandler handles the deletion of a routine by ID.
func DeleteRoutineHandler(c *gin.Context, collection *mongo.Collection) {
    // Get the routine ID from the request parameters
    routineID := c.Param("id")

    // Convert the routine ID to an ObjectId
    objID, err := primitive.ObjectIDFromHex(routineID)
    if err != nil {
        // Return error response if the ID is invalid
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid routine ID"})
        return
    }

    // Define a filter to find the routine by ID
    filter := bson.M{"_id": objID}

    // Delete the routine from MongoDB
    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        log.Printf("Error deleting routine: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete routine"})
        return
    }

    // Check if any routines were deleted
    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No routine found with specified ID"})
        return
    }

    // Return success message
    c.JSON(http.StatusOK, gin.H{"message": "Routine deleted successfully"})
}



// EditRoutineHandler handles the editing of an existing routine.
func EditRoutineHandler(c *gin.Context, collection *mongo.Collection) {
    // Get the routine ID from the request parameters
    routineID := c.Param("id")

    // Convert the routine ID to an ObjectId
    objID, err := primitive.ObjectIDFromHex(routineID)
    if err != nil {
        // Return error response if the ID is invalid
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid routine ID"})
        return
    }

    // Define a filter to find the routine by ID
    filter := bson.M{"_id": objID}

    // Define a struct to hold the updated routine data
    var updatedRoutine models.Routine

    // Bind the request body to the updatedRoutine struct
    if err := c.ShouldBindJSON(&updatedRoutine); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the routine in MongoDB
    update := bson.M{"$set": updatedRoutine}
    _, err = collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Printf("Error updating routine: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update routine"})
        return
    }

    // Return success message
    c.JSON(http.StatusOK, gin.H{"message": "Routine updated successfully"})
}
