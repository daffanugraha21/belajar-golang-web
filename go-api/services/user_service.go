package services

import (
	"go-api/config"
	"go-api/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return &user, result.Error
}

func CreateUser(name string) (*models.User, error) {
	user := models.User{Name: name}
	result := config.DB.Create(&user)
	return &user, result.Error
}

func UpdateUser(id uint, name string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	user.Name = name
	config.DB.Save(&user)
	return &user, nil
}

func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}