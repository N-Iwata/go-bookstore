package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/N-Iwata/go-bookstore/pkg/models"
	"github.com/N-Iwata/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

type Response struct {
	Success bool `json:"success"`
}

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	book, err := models.GetBookByID(ID)
	if err != nil {
		fmt.Println(err)
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	if err := models.DeleteBook(ID); err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(&Response{Success: true})
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	book, err := models.GetBookByID(ID)
	if err != nil {
		fmt.Println(err)
	}
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	b, err := book.UpdateBook()
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(b)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
