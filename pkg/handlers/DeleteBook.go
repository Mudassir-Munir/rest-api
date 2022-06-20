package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mudassir-Munir/rest-api/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//iterate over all books
	for index, book := range mocks.Books {
		if book.Id == id {
			//Delete book and send respons if id matches with dynamic id
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted Book")
			break

		}
	}

}
