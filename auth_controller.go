package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"
	

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"example/Golang-Projects/database"
	"example/Golang-Projects/models"
	"example/Golang-Projects/utils"
	
)

// @Summary     Register New User
// @Description Create New Account For User
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       user body models.RegisterRequest true "Data User"
// @Success     201 {object} map[string]string "account has been created"
// @Failure     400 {object} map[string]string "Register is cannot possible"
// @Failure     500 {object} map[string]string "Register is failed"
// @Router      /auth/register [post]
func Register(c *gin.Context) {
	var newUser models.User

	fmt.Print(&newUser)
	//bind json
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Register is error"})
		return
	}

	if err := database.DB.Create(&newUser). Error; err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Failed to register"})
		return
	}
	
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password is incorrect"})
		return
	}
	newUser.Password = hashedPassword
	c.JSON(http.StatusCreated, gin.H{"message" : "Account has been created!"})
}

// @Summary      Authentication Users
// @Description  Login With The Username, Email and Password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.LoginRequest true "Data User"
// @Success      200  {object}  map[string]string "Login is completed"
// @Failure      400  {object}  map[string]string "Login is failed"
// @Failure      401  {object}  map[string]string "Errors"
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Input is not valid"})
		return
	}
	
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Email is incorrect"})
		return
	}
	// check password either
	if err := utils.CheckPassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password is incorrect"})
		return
	
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Token is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
//more safety
func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp" : time.Now().Add(time.Hour * 24). Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))
}