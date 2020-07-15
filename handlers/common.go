package handlers

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Check if user is Authenticated
func IsAuthenticated(c *gin.Context) (bool, *Claims) {

	// Get Cookie
	cookie, _ := c.Cookie("auth_token")

	// Check if Empty Cookie
	if cookie == "" {
		log.Println("Empty Token !!")
		return false, nil
	}

	claims := &Claims{}

	// Parse Cookie
	token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {

		log.Println("Error reading token :", err)
		return false, nil
	}
	if !token.Valid {

		log.Println("Unauthorized User!!")
		return false, nil
	}

	return true, claims
}
