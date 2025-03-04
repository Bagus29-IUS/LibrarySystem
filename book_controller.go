package controllers

import (
	"example/Golang-Projects/models"
	"example/Golang-Projects/database"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(database *gorm.DB) {
	DB = database
}

// GetBooks godoc
// @Summary Get All Books
// @Description Get all books data in database
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks (c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	if database.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
        return	
	}
	c.IndentedJSON(http.StatusOK, books)
}

// GetBookByID godoc
// @Summary Get Detail About Books With ID
// @Description Find ID to Get Specific Books
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "ID Book"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [get]
func GetBookByID (c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

// @Summary     Create New Books
// @Description Get New Books to Database
// @Tags        books
// @Accept      json
// @Produce     json
// @Param       request body models.CreateRequest true "Data Book"
// @Success     201 {object} models.Book
// @Failure     400 {object} map[string]string "Input not valid"
// @Router      /books/create [post]
func CreateBook(c *gin.Context) {
	var NewBook models.Book

	if err := c.ShouldBindJSON(&NewBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Format JSON is not valid"})
		return
	}

	if err := database.DB.Create(&NewBook).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Failed to load a book"})
		return
	}

	c.IndentedJSON(http.StatusCreated, NewBook)
}


// @Summary     Checkout Books
// @Description Checkout Books, Will be Decrease Quantity.
// @Tags        books
// @Accept      json
// @Produce     json
// @Param       id query string true "ID Book"
// @Success     200 {object} models.Book
// @Failure     400 {object} map[string]string "Id not found"
// @Failure     404 {object} map[string]string "book not found"
// @Router      /books/checkout [post]
func CheckoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Id not found"})
		return
	}

	fmt.Println("ID query params:", id)

	var book models.Book

	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found"})
		return
	}

	database.DB.Model(&book).Where("id = ?", id).Update("quantity", gorm.Expr("quantity - 1"))
	c.IndentedJSON(http.StatusOK, book)
}

// @Summary     Return Books
// @Description Return Books by User, Will be Increase Books Quantity
// @Tags        books
// @Accept      json
// @Produce     json
// @Param       id query string true "ID Book"
// @Success     200 {object} models.Book
// @Failure     400 {object} map[string]string "Id not found"
// @Failure     404 {object} map[string]string "Book not found"
// @Router      /books/return [PATCH]
func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Format JSON is not valid"})
		return
	}

	var book models.Book
	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"please fill the ID"})
		return
	}

	database.DB.Model(&book).Where("id = ?", id).Update("quantity", gorm.Expr("quantity + 1"))
	c.IndentedJSON(http.StatusOK, book)
}	 

