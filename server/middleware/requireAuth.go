package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"example.com/fullstack-template/database"
	"example.com/fullstack-template/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("RequireAuth middleware hit")

	// Get cookie from request
	tokenString, err := c.Cookie("authorization_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Decode/validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//Check the expiration date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		// Find user with token subject
		var user models.User
		result := database.DB.First(&user, claims["sub"])
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Attach user to context
		// var currentUser = gin.H{
		// 	"id":        user.ID,
		// 	"username":  user.Username,
		// 	"firstName": user.FirstName,
		// 	"lastName":  user.LastName,
		// 	"email":     user.Email,
		// 	"createdAt": user.CreatedAt,
		// 	"updatedAt": user.UpdatedAt,
		// 	"deletedAt": user.DeletedAt,
		// 	"plays":     user.Plays,
		// }

		// Continue to next handler
		c.Set("currentUserId", user.ID)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

}
