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


// GetWorkoutsPerWeekHandler handles the GET request to fetch workouts per week.
func GetWorkoutsPerWeekHandler(c *gin.Context, collection *mongo.Collection) {
    // Extract the user ID from the request path
    userID := c.Param("userID")

    // Get the current time in the server's time zone
    now := time.Now()

    // Calculate the start of the current week (Monday) in the server's time zone
    currentWeekStart := now.AddDate(0, 0, -int(now.Weekday())+1).Truncate(24 * time.Hour)
    if now.Weekday() == time.Sunday {
        currentWeekStart = now.AddDate(0, 0, -6).Truncate(24 * time.Hour)
    }

    // Define the start date for the last 5 weeks to include the current week
    fiveWeeksAgo := currentWeekStart.AddDate(0, 0, -28)

    // Initialize a map to store workout counts per week
    workoutsPerWeek := make(map[string]int)

    // Iterate over the last 5 weeks to include the current week
    for i := 0; i < 5; i++ {
        // Define the start and end date for the week
        weekStart := fiveWeeksAgo.AddDate(0, 0, 7*i).Truncate(24 * time.Hour)
        weekEnd := weekStart.AddDate(0, 0, 6)

        // Define a filter to query workouts by user ID and within the week
        filter := bson.M{
            "user_id": userID,
            "date": bson.M{
                "$gte": weekStart,
                "$lte": weekEnd.Add(24*time.Hour).Add(-time.Second), // Set the end of the day by subtracting one second
            },
        }

        // Execute the count operation to get the number of workouts matching the filter
        count, err := collection.CountDocuments(context.Background(), filter)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workouts"})
            return
        }

        // Store the workout count for the week
        workoutsPerWeek[weekStart.Format("2006-01-02")] = int(count)
    }

    // Return the workout counts per week as a response
    c.JSON(http.StatusOK, workoutsPerWeek)
}

