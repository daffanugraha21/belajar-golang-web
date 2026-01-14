package services

import (
	"go-api/models"
)

var users = []models.User{
	{ID:1, Name: "Alice"},
	{ID:2, Name: "Bob"},
}

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id int) (*models.User, bool) {
	for _, user := range users {
		if user.ID == id {
			return &user, true
		}
	}
	return nil, false
}

func CreateUser(name string) models.User {
	newUser := models.User{
		ID: len(users) + 1,
		Name : name,
	}
	users = append(users, newUser)
	return newUser
}

func UpdateUser(id int, name string) (*models.User, bool) {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name
			return &users[i], true
		}
	}
	return nil, false
}

func DeleteUser(id int) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}