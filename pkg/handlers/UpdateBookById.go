package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Mudassir-Munir/rest-api/pkg/mocks"
	"github.com/Mudassir-Munir/rest-api/pkg/modles"
	"github.com/gorilla/mux"
)

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
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
	for index, book := range mocks.Books {
		if book.Id == id {
			//update book and send response when id mathces dynamic id
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
