package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

//Get /users
func GetUsers(c *gin.Context){
	c.JSON(http.StatusOK, users)
}

// Get /users/:id
func GetUserByID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "invalid id",
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message" : "user not found",
	})
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

	newUser := User{
		ID:   len(users) + 1,
		Name: body.Name,
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
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

	for i, user := range users {
		if user.ID == id {
			users[i].Name = body.Name
			c.JSON(http.StatusOK, users[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message" : "user not found",
	})
}

// delete /users/:id
func DeleteUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message" : "user deleted",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message" : "user not found",
	})
}