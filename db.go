package database

import (
	"example/Golang-Projects/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
	var DB *gorm.DB

	// fatching env 
	func InitDB() {
		err := godotenv.Load("config.env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}	
		//fetching data to mysql
		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
	)

	//connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database:%v", err)
	}
		DB = db
		DB.AutoMigrate(&models.Book{})
	}
