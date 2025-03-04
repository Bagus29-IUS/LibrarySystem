package main

import (
	"example/Golang-Projects/database"
	"example/Golang-Projects/routes"
   "example/Golang-Projects/docs"
   "fmt"
	
	"github.com/swaggo/gin-swagger"	
	"github.com/swaggo/files"
	"github.com/gin-gonic/gin"
)
// @title Documentation API using Swagger
// @version 1.0
// @description API for register, login, checkout books, etc.
// @host localhost:8080
// @BasePath /
	func main() {
	
		database.InitDB()
		
		router := gin.Default()
		
		// Swagger setup
		docs.SwaggerInfo.Title = "My API"
		docs.SwaggerInfo.Description = "Dummy Library System"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.BasePath = "/"

		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		//setup routes
		routes.AuthRoutes(router)
		routes.BookRoutes(router)

		//jalankan server
		port :=":8080"
		fmt.Println("Server berjalan di http://localhost" + port)
		router.Run(port)
	}