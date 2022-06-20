package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Mudassir-Munir/rest-api/pkg/modles"
	"github.com/gorilla/mux"
)

func (h handler) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id parameter
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	var updatedBook modles.Book
	json.Unmarshal(body, &updatedBook)

	//iterate over all books
	var book modles.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
