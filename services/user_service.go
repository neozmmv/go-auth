package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/neozmmv/go-auth/database"
	"github.com/neozmmv/go-auth/models"
	"github.com/neozmmv/go-auth/utils"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching users: %w", result.Error)
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("error creating user: %w", result.Error)
	}
	return nil
}

func FindUser(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("error finding user: %w", result.Error)
	}
	return &user, nil
}

func DeleteUser(id uuid.UUID) error {
	result := database.DB.Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("error deleting user: %w", result.Error)
	}
	return nil
}
