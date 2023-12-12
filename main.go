package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents the user structure
type User struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	Username string      `json:"username"`
	Sex      string      `json:"sex"`
	Weight   int         `json:"weight"`
}

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load MongoDB URI and Database Name from the environment variables
	mongoURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DATABASE_NAME")

	// Set up the MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Set the collection variable for further use
	collection = client.Database(dbName).Collection("Users")
}

func main() {
	r := gin.Default()

	r.POST("/register", registerHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func registerHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the user into MongoDB with the specified _id
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error inserting user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	// Return the _id as part of the response
	c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}
