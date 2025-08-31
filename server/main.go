package main

import (
	"example.com/fullstack-template/controllers"
	"example.com/fullstack-template/database"
	"example.com/fullstack-template/middleware"
	"example.com/fullstack-template/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDatabase()

	database.DB.AutoMigrate(&models.User{})

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowCredentials: true,
	}))

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)
	router.GET("/validate", middleware.RequireAuth, controllers.ShowValidatedUser)

	router.Run("localhost:8080")
}
