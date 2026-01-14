package controllers

import (
	"github.com/gin-gonic/gin"

	"go-api/services"
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
	var body struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" {
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	user, err := services.CreateUser(body.Name)
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to create user"})
		return
	}

	c.JSON(201, user)
}