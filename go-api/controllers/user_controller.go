package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api/models"
	"go-api/services"
	"go-api/config"
	"go-api/helpers"
	"strconv"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to fetch users"})
		return
	}
	c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required,min=3"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := helpers.FormatValidationErrors(err)

		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors,})
		return
	}

	user := models.User{
		Name: req.Name,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	user.Name = input.Name
	config.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}