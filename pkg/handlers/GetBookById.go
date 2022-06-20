package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mudassir-Munir/rest-api/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//w.Header().Add("Content-Type", "application/json")

	//Read dynamic Id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//iterate through all books
	for _, book := range mocks.Books {
		if book.Id == id {
			//if ids are same send book as response
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}

}
