package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aron-helu/Book_Management-Go/pkg/models"
	"github.com/aron-helu/Book_Management-Go/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

var NewBook models.Book

func CreatBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	NewBook = models.Book{}
	utils.ParseBody(r, &NewBook)
	NewBook.CreateBook()
	p, _ := json.Marshal(NewBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", p)
}

func GetBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Books := models.GetAllBooks()
	p, _ := json.Marshal(Books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", p)
}

func GetBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookID, err := strconv.ParseInt(ps.ByName("bookId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookDetails, _ := models.GetBookById(bookID)
	if bookDetails == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	responseJSON, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")		
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)

	if err !=  nil {
		fmt.Println("error while parasing")
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author	
	}
	if updateBook.Publisher != ""{
		bookDetails.Publisher = updateBook.Publisher
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content Type", "pkgication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
