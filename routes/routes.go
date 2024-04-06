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

	// Define endpoint for getting a user by their ID
	r.GET("/user/:id", func(c *gin.Context) {
		handlers.GetUserHandler(c, usersCollection)
	})

	// Define endpoint for getting all exercises
	r.GET("/exercises", func(c *gin.Context) {
		handlers.GetAllExercisesHandler(c, exercisesCollection)
	})

	// Register the endpoint for saving workouts
	r.POST("/workouts", func(c *gin.Context) {
		handlers.SaveWorkoutHandler(c, workoutsCollection)
	})

	r.GET("/workouts/:userID", func(c *gin.Context) {
		handlers.GetWorkoutsByUserIDHandler(c, workoutsCollection)
	})

	// Delete a workout by ID
	r.DELETE("/delete-workout/:id", func(c *gin.Context) {
		handlers.DeleteWorkoutHandler(c, workoutsCollection)
	})

	// Create a new routine
	r.POST("/create-routine", func(c *gin.Context) {
		handlers.CreateRoutineHandler(c, routinesCollection)
	})

	// Fetch a routines by UserId
	r.GET("/routines/:id", func(c *gin.Context) {
		handlers.GetRoutinesByUserIDHandler(c, routinesCollection)
	})	

	// Delete a routine by ID
	r.DELETE("/delete-routine/:id", func(c *gin.Context) {
		handlers.DeleteRoutineHandler(c, routinesCollection)
	})


	
	
}
