package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mudassir-Munir/rest-api/pkg/modles"
	"github.com/gorilla/mux"
)

func (h handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//w.Header().Add("Content-Type", "application/json")

	//Read dynamic Id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//iterate through all books
	// for _, book := range mocks.Books {
	// 	if book.Id == id {
	//

	// 	}
	// }

	var book modles.Book
	//if ids are same send book as response
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

}
