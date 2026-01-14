package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HelloRequest struct {
	Name string `json:"name"`
}

func Hello (c *gin.Context){
	var req HelloRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message" : "invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "hello" + " " + req.Name,
	})
}