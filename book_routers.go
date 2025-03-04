package routes

import (
	"example/Golang-Projects/controllers"
	"github.com/gin-gonic/gin"
)
//router for access
func BookRoutes(router *gin.Engine) {
	books := router.Group("/books") 
	
{
	books.GET("/", controllers.GetBooks)
	books.GET("/:id", controllers.GetBookByID)
	books.POST("/create", controllers.CreateBook)
	books.POST("/checkout", controllers.CheckoutBook)
	books.PATCH("/return", controllers.ReturnBook)
}
}