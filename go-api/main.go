package main

import (
	"go-api/config"
	"go-api/models"
	"go-api/routes"
)

func main() {
	config.ConnectDB()

	// Auto migrate
	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRoutes()
	r.Run(":8080")
}