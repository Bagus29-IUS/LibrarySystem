package models

import(
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterRequest struct for registration request
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email     string  `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//LoginRequest struct for login request
type LoginRequest struct {
	    Username string `json:"username" binding:"required"`
		Email string `gorm:"unique" json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

// User struct for table users in database
type User struct{
	ID        uint    `gorm:"primaryKey" json:"id"`
	Username  string  `gorm:"unique" json:"username" binding:"required"` 
	Email     string  `gorm:"unique" json:"email" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}