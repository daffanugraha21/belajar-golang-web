package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()

	r .GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/hello", func(c *gin.Context) {
		var body struct {
			Name string `jsone:"name"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
            "message": "hello " + body.Name,
  	})
	})
	r.Run(":8080")
}