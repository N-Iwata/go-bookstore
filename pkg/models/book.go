package models

import (
	"github.com/N-Iwata/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	db = config.DbConnect()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(ID int64) (*Book, error) {
	var book Book
	if err := db.Where("ID=?", ID).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func DeleteBook(ID int64) error {
	var book Book
	if err := db.Where("ID=?", ID).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

func (b *Book) UpdateBook() (*Book, error) {
	if err := db.Model(&b).Updates(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}
