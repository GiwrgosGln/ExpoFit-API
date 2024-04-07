package handlers

import (
	"context"
	"log"
	"net/http"

	"FitnessAPI/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterHandler(c *gin.Context, collection *mongo.Collection) {
	var user models.User

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

func GetUserHandler(c *gin.Context, collection *mongo.Collection) {
	// Get user ID from path parameter
	userID := c.Param("id")

	// Find user in MongoDB by ID
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Printf("Error finding user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	// Return user data
	c.JSON(http.StatusOK, user)
}

// EditUserHandler is responsible for updating user information in the database.
func EditUserHandler(c *gin.Context, collection *mongo.Collection) {
	// Get user ID from path parameter
	userID := c.Param("id")

	// Create a struct to hold the updated user data
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user in MongoDB
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"username":     updatedUser.Username,
			"email":        updatedUser.Email,
			"gender":       updatedUser.Gender,
			"dateofbirth": updatedUser.DateOfBirth,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
