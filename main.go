// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"FitnessAPI/database"
	"FitnessAPI/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection
var exercisesCollection *mongo.Collection

func main() {
	r := gin.Default()

	var err error

	// Connect to "Users" collection
	usersCollection, err = database.ConnectMongoDB("Users")
	if err != nil {
		log.Fatal("Error connecting to Users collection:", err)
	}

	// Connect to "Exercises" collection
	exercisesCollection, err = database.ConnectMongoDB("Exercises")
	if err != nil {
		log.Fatal("Error connecting to Exercises collection:", err)
	}

	// Connect to "Workouts" collection
	workoutsCollection, err := database.ConnectMongoDB("Workouts")
	if err != nil {
		log.Fatal("Error connecting to Workouts collection:", err)
	}

	// Connect to "Routines" collection
	routinesCollection, err := database.ConnectMongoDB("Routines")
	if err != nil {
		log.Fatal("Error connecting to Routines collection:", err)
	}

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
