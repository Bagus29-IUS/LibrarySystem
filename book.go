package models

//struct for more precise creating books
type CreateRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

//struct book for data foundation
type Book struct {
	ID     uint   `json:"id"       gorm:"primaryKey"`
	Title  string `json:"title"    binding:"required"`
	Author string `json:"author"   binding:"required"`
	Genre  string `json:"genre"    binding:"required"`
	Quantity int  `json:"quantity" binding:"required"`
}
