package routes

import (
	"github.com/gin-gonic/gin"
	"example/Golang-Projects/controllers"
)
// SetupAuthRoutes control routes for debugging
func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}

