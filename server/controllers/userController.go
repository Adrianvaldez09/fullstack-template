package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"example.com/fullstack-template/database"
	"example.com/fullstack-template/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	fmt.Println("Signup endpoint hit")

	var SignupInput struct {
		Username  string `json:"username" binding:"required"`
		FirstName string `json:"firstName" binding:"required"`
		LastName  string `json:"lastName" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
	}

	// Bind request body to struct and validate
	if c.Bind(&SignupInput) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	fmt.Println("Parsed input:", SignupInput)

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(SignupInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user

	user := models.User{
		Username:  SignupInput.Username,
		FirstName: SignupInput.FirstName,
		LastName:  SignupInput.LastName,
		Email:     SignupInput.Email,
		Password:  string(hash),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

func Login(c *gin.Context) {
	fmt.Println("Login endpoint hit")

	//Extract login details from request
	var LoginInput struct {
		Identity string `json:"identity" binding:"required"` // Can be username or email
		Password string `json:"password" binding:"required"`
	}

	// Bind input to struct
	if c.Bind(&LoginInput) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	fmt.Print("Login input struct initialized: ", LoginInput)

	//Look up user by username or email
	var user models.User

	result := database.DB.Where("username = ? OR email = ?", LoginInput.Identity, LoginInput.Identity).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/email or password"})
		return
	}

	//Compare provided password with stored hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginInput.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/email or password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	//Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	//Send the token in a cookie
	c.SetCookie("authorization_token", tokenString, 3600*72, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func ShowValidatedUser(c *gin.Context) {
	fmt.Println("Validate endpoint hit")

	currentUser, _ := c.Get("currentUser")

	c.JSON(http.StatusOK, gin.H{"currentUser": currentUser})
}

func Logout(c *gin.Context) {
	fmt.Println("Logout endpoint hit")
	// Clear the authorization_token cookie by setting its MaxAge to -1
	c.SetCookie("authorization_token", "", -1, "/", "", false, true)

}
