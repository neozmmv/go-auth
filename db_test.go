package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/neozmmv/go-auth/database"
	"github.com/neozmmv/go-auth/models"
	"github.com/neozmmv/go-auth/services"
	"github.com/neozmmv/go-auth/utils"
)

func TestConnectToDb(t *testing.T) {
	godotenv.Load()
	t.Log("Connecting to database...")
	database.ConnectDatabase()
	if database.DB == nil {
		t.Fatal("Failed to connect to database")
	}
}

func TestGetAllUsers(t *testing.T) {
	godotenv.Load()
	database.ConnectDatabase()
	users, err := services.GetAllUsers()
	if err != nil {
		t.Fatalf("Error fetching users: %v", err)
	}
	for _, user := range users {
		t.Logf("User: %s, Email: %s", user.Name, user.Email)
	}
}

func TestCreateUser(t *testing.T) {
	godotenv.Load()
	database.ConnectDatabase()
	newUser := models.User{
		Name:     "Test User",
		Email:    "test@mail.com",
		Password: "testpassword",
	}
	err := services.CreateUser(&newUser)
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
}

func TestComparePassword(t *testing.T) {
	password := "testpassword"
	hashedPassword, _ := utils.HashPassword(password)
	if !utils.ComparePassword(hashedPassword, password) {
		t.Fatal("Password comparison failed")
	}
}

func TestDeleteUser(t *testing.T) {
	godotenv.Load()
	database.ConnectDatabase()
	// put a uuid here to delete
	uuid, _ := uuid.Parse("7d4cd97d-82f1-4f7a-b7d6-540dcbd29cf4")
	err := services.DeleteUser(uuid)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}
