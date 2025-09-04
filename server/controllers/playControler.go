package controllers

import (
	"fmt"
	"net/http"

	"example.com/fullstack-template/database"
	"example.com/fullstack-template/models"
	"github.com/gin-gonic/gin"
)

func CreatePlay(c *gin.Context) {
	fmt.Println("CreatePlay endpoint hit")

	var CreatePlayInput struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		CreatorID   uint   `json:"creatorId" binding:"required"`
	}

	// Bind request body to struct and validate
	if c.ShouldBindJSON(&CreatePlayInput) != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Create play
	play := models.Play{
		Name:        CreatePlayInput.Name,
		Description: CreatePlayInput.Description,
		CreatorID:   CreatePlayInput.CreatorID,
	}

	result := database.DB.Create(&play)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create play"})
		return
	}

	// Look up user and preload plays
	var user models.User
	if err := database.DB.Preload("Plays").First(&user, CreatePlayInput.CreatorID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Return updated user with their plays
	c.JSON(http.StatusOK, user)
}
