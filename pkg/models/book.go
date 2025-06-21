package models

import (
	"github.com/ayushmaurya461/go-book-store.git/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `gorm:"" json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(id int) (Book, *gorm.DB) {
	var book Book
	result := db.First(&book, id)
	return book, result
}

func DeleteBook(id int) {
	var book Book
	db.Where("id = ?", id).Delete(&book)
}
