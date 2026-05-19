package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/neozmmv/go-auth/models"
	"github.com/neozmmv/go-auth/services"
	"github.com/neozmmv/go-auth/utils"
)

func Login(c *gin.Context) {
	errors := []string{}
	var user models.User
	// error handling
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request! (missing body)"})
		return
	}
	if user.Email == "" {
		errors = append(errors, "Email is required")
	}
	if user.Password == "" {
		errors = append(errors, "Password is required")
	}
	if len(errors) > 0 {
		c.JSON(400, gin.H{"errors": errors})
		return
	}
	// logic to find user and validate
	foundUser, err := services.FindUser(user.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.ComparePassword(foundUser.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(foundUser.Name, foundUser.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func SignUp(c *gin.Context) {
	errors := []string{}
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if newUser.Name == "" {
		errors = append(errors, "Name is required")
	}
	if newUser.Email == "" {
		errors = append(errors, "Email is required")
	}
	if newUser.Password == "" {
		errors = append(errors, "Password is required")
	}
	if len(errors) > 0 {
		c.JSON(400, gin.H{"errors": errors})
		return
	}
	err := services.CreateUser(&newUser)
	if err != nil {
		c.JSON(409, gin.H{"error": "User already exists with this email"})
		return
	}
	c.JSON(201, newUser)
}

func VerifyToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	if tokenString == "" {
		c.JSON(400, gin.H{"error": "Authorization header is required"})
		return
	}
	_, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(200, gin.H{"message": "Token is valid"})
}

func WhoAmI(c *gin.Context) {
	tokenString := c.GetString("token")
	token, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(500, gin.H{"error": "Failed to extract claims"})
		return
	}
	c.JSON(200, gin.H{
		"name":  claims["name"],
		"email": claims["email"],
	})
}
