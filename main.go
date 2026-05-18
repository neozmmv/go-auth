package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/neozmmv/go-auth/controllers"
	"github.com/neozmmv/go-auth/middleware"
)

func main() {
	// http://localhost:8000
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file! Make sure you have a .env file with JWT_SECRET variable. (32 chars)")
	}
	r := gin.Default()
	r.GET("/users", middleware.Auth, controllers.GetUsers)
	r.POST("/users", middleware.Auth, controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.POST("/verify", controllers.VerifyToken)
	r.GET("/me", middleware.Auth, controllers.WhoAmI)
	r.Run(":8000")
}
