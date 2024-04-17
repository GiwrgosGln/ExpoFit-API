// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"FitnessAPI/database"
	"FitnessAPI/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection
var exercisesCollection *mongo.Collection
var workoutsCollection *mongo.Collection
var routinesCollection *mongo.Collection
var measurementsCollection *mongo.Collection

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
	workoutsCollection, err = database.ConnectMongoDB("Workouts")
	if err != nil {
		log.Fatal("Error connecting to Workouts collection:", err)
	}

	// Connect to "Routines" collection
	routinesCollection, err = database.ConnectMongoDB("Routines")
	if err != nil {
		log.Fatal("Error connecting to Routines collection:", err)
	}

	// Connect to "Measurements" collection
	measurementsCollection, err = database.ConnectMongoDB("Measurements")
	if err != nil {
		log.Fatal("Error connecting to Measurements collection:", err)
	}

	// Initialize routes
	routes.SetupRoutes(r, usersCollection, exercisesCollection, workoutsCollection, routinesCollection, measurementsCollection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
