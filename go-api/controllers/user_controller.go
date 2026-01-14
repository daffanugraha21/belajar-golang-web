package controllers

import (
	"go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Get /users
func GetUsers(c *gin.Context){
	c.JSON(http.StatusOK, services.GetAllUsers())
}

// Get /users/:id
func GetUserByID(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	user, found := services.GetUserByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message":"user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	user := services.CreateUser(body.Name)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	user, updated := services.UpdateUser(id, body.Name)
	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"message" : "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// delete /users/:id
func DeleteUser(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	if !services.DeleteUser(id) {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}