package handlers

import (
	"context"
	"log"
	"net/http"

	"FitnessAPI/models"

	"github.com/gin-gonic/gin"
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
