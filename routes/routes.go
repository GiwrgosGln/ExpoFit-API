// routes.go
package routes

import (
	"FitnessAPI/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine, usersCollection, exercisesCollection, workoutsCollection, routinesCollection *mongo.Collection) {
	// Define endpoint for registering users
	r.POST("/register", func(c *gin.Context) {
		handlers.RegisterHandler(c, usersCollection)
	})

	// Define endpoint for getting all exercises
	r.GET("/exercises", func(c *gin.Context) {
		handlers.GetAllExercisesHandler(c, exercisesCollection)
	})

	// Register the endpoint for saving workouts
	r.POST("/workouts", func(c *gin.Context) {
		handlers.SaveWorkoutHandler(c, workoutsCollection)
	})

	// Create a new routine
	r.POST("/create-routine", func(c *gin.Context) {
		handlers.CreateRoutineHandler(c, routinesCollection)
	})

	// Set up additional endpoints as needed
}
