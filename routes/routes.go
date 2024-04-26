package routes

import (
	"FitnessAPI/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine, usersCollection, exercisesCollection, workoutsCollection, routinesCollection, measurementsCollection *mongo.Collection) {
	// Define endpoint for registering users
	r.POST("/register", func(c *gin.Context) {
		handlers.RegisterHandler(c, usersCollection)
	})

	// Define endpoint for getting a user by their ID
	r.GET("/user/:id", func(c *gin.Context) {
		handlers.GetUserHandler(c, usersCollection)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		handlers.EditUserHandler(c, usersCollection)
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

	r.GET("/workouts-last5weeks/:userID/", func(c *gin.Context) {
		handlers.GetWorkoutsPerWeekHandler(c, workoutsCollection)
	})

	// Create a new routine
	r.POST("/create-routine", func(c *gin.Context) {
		handlers.CreateRoutineHandler(c, routinesCollection)
	})

	// Fetch a routine by UserId
	r.GET("/routines/:id", func(c *gin.Context) {
		handlers.GetRoutinesByUserIDHandler(c, routinesCollection)
	})

	// Edit a routine by RoutineID
	r.PATCH("/edit-routine/:id", func(c *gin.Context) {
		handlers.EditRoutineHandler(c, routinesCollection)
	})	

	// Delete a routine by ID
	r.DELETE("/delete-routine/:id", func(c *gin.Context) {
		handlers.DeleteRoutineHandler(c, routinesCollection)
	})

	// Create a new measurement
	r.POST("/create-measurement", func(c *gin.Context) {
		handlers.CreateMeasurementHandler(c, measurementsCollection)
	})

	// Define endpoint for getting all measurements for a user
	r.GET("/measurements/:userID", func(c *gin.Context) {
		handlers.GetMeasurementsByUserIDHandler(c, measurementsCollection)
	})

	// Delete a measurement by ID
	r.DELETE("/delete-measurement/:id", func(c *gin.Context) {
		handlers.DeleteMeasurementHandler(c, measurementsCollection)
	})

	// Define endpoint for getting workouts by exercise name and user ID
	r.GET("/exercise-sets", func(c *gin.Context) {
		handlers.GetExerciseSetsHandler(c, workoutsCollection)
	})

	// Add the route for calculating one-rep max
	r.GET("/one-rep-max", func(c *gin.Context) {
		handlers.CalculateOneRepMaxHandler(c, workoutsCollection)
	})


}
