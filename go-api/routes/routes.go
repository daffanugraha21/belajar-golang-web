package routes

import (
	"go-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.POST("/hello", controllers.Hello)
}