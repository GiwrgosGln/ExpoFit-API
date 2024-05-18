package handlers

import (
	"context"
	"math"
	"net/http"
	"time"

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

// GetExerciseSetsHandler retrieves workout sets for a specific exercise, user ID, and date
func GetExerciseSetsHandler(c *gin.Context, collection *mongo.Collection) {
	// Parse exercise name, user ID, and date from request parameters
	exerciseName := c.Query("exercise")
	userID := c.Query("userid")
	date := c.Query("date")

	// Query to find workouts matching exercise name, user ID, and date
	filter := bson.M{"user_id": userID, "exercises.name": exerciseName}
	if date != "" {
		filter["date"] = date
	}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve workouts"})
		return
	}
	defer cursor.Close(context.Background())

	// Initialize map to store sets grouped by workout date
	setsByDate := make(map[string][]models.Set)

	var workouts []models.Workout
	if err := cursor.All(context.Background(), &workouts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode workouts"})
		return
	}

	// Extract sets with non-null values for weight and reps
	for _, workout := range workouts {
		for _, exercise := range workout.Exercises {
			if exercise.Name == exerciseName {
				for _, set := range exercise.Sets {
					if set.Weight != nil && set.Reps != nil {
						// Add set to map, grouped by workout date
						dateString := workout.Date.Format(time.RFC3339)
						setsByDate[dateString] = append(setsByDate[dateString], set)
					}
				}
			}
		}
	}

	// Prepare response
	var response []struct {
		Date string         `json:"date"`
		Sets []models.Set `json:"sets"`
	}
	for date, sets := range setsByDate {
		response = append(response, struct {
			Date string         `json:"date"`
			Sets []models.Set `json:"sets"`
		}{
			Date: date,
			Sets: sets,
		})
	}

	c.JSON(http.StatusOK, response)
}

// CalculateOneRepMaxHandler calculates the one-rep max for each set of a specific exercise
func CalculateOneRepMaxHandler(c *gin.Context, collection *mongo.Collection) {
    // Parse exercise name, user ID, and date from request parameters
    exerciseName := c.Query("exercise")
    userID := c.Query("userid")
    date := c.Query("date")

    // Query to find workouts matching exercise name, user ID, and date
    filter := bson.M{"user_id": userID, "exercises.name": exerciseName}
    if date != "" {
        filter["date"] = date
    }
    cursor, err := collection.Find(context.Background(), filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve workouts"})
        return
    }
    defer cursor.Close(context.Background())

    var workouts []models.Workout
    if err := cursor.All(context.Background(), &workouts); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode workouts"})
        return
    }

    // Initialize a map to store the one-rep max value for each workout date
    oneRepMaxByDate := make(map[string]float64)

    // Calculate one-rep max for each set
    for _, workout := range workouts {
        dateString := workout.Date.Format("2006-01-02") // Format date as "YYYY-MM-DD"
        for _, exercise := range workout.Exercises {
            if exercise.Name == exerciseName {
                for _, set := range exercise.Sets {
                    if set.Weight != nil && set.Reps != nil {
                        // Convert Weight to float64 and calculate one-rep max using the formula
                        weight := float64(*set.Weight)
                        oneRepMax := weight / (1.0278 - 0.0278*float64(*set.Reps))
                        // Take the absolute value of the calculated one-rep max
                        oneRepMax = math.Abs(oneRepMax)
                        // Update the one-rep max value for the workout date
                        oneRepMaxByDate[dateString] = oneRepMax
                    }
                }
            }
        }
    }

    // Prepare response with dates and corresponding one-rep max values
    var response []map[string]interface{}
    for date, max := range oneRepMaxByDate {
        entry := map[string]interface{}{
            "date": date,
            "max":  max,
        }
        response = append(response, entry)
    }

    c.JSON(http.StatusOK, response)
}
